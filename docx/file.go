package docx


import ( 
    "io" 
    "os"  
    "bufio"      
    "bytes"
    "errors"    
    "regexp"
    "strings"
    "io/ioutil"    
    "archive/zip"    
)

var (
    rxStartTags = regexp.MustCompile("<(\\w+)")
    rxEndTags = regexp.MustCompile("<\\/(\\w+)")
    rxVals = regexp.MustCompile("\\s(\\w+)=")
    rxXMLNsVals = regexp.MustCompile("\\s(\\w+)=\"http://schemas.")
    rxURnVals = regexp.MustCompile("\\s(\\w+)=\"urn:")    

    emptyTags = []string{"top","left","bottom","right","insideV","insideH",
                         "shd","jc","vAlign","vMerge", "noWrap", "docGrid",
                         "b","bCs","i","u", "sz", "szCs", "color", "hideMark",
                         "tblLayout","tblHeader","tblInd","tblW","gridCol", "gridSpan",
                         "pStyle","rFonts","rtl","tcW","bidi","trHeight","lang", 
                         "pgSz", "pgMar", "headerReference", "footerReference", "br", "tab"}
)

// SimpleDocxFile - файл docx
type SimpleDocxFile struct {
    zipFile  *zip.ReadCloser
    headers   map[string]*Header
    document *Document    
}

// OpenFile - Открытие файла DOCX
func OpenFile(fileName string) (*SimpleDocxFile,error) {
    z, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}    
    d := new(SimpleDocxFile)
    d.headers = make(map[string]*Header)    
    d.zipFile = z
    // Перебор файлов в Zip архиве
    for _, f := range z.File {
        if f != nil {
            // Загрузка документа            
            if f.Name == "word/document.xml" {
                reader, err := f.Open()
                if err != nil {
                    return nil, err
                }
                d.document = new(Document)
                d.document.Decode(reader)
                if err := reader.Close(); err != nil {
                    return nil, err
                }                
            } else if strings.Index(f.Name, "word/header") >= 0 {
                reader, err := f.Open()
                if err != nil {
                    return nil, err
                }
                header := new(Header)
                header.Decode(reader)
                if err := reader.Close(); err != nil {
                    return nil, err
                }
                d.headers[f.Name] = header
            } 
        }
    }    
    return d, nil 
}

// Render (SimpleDocxFile) - рендер шаблона
func (f *SimpleDocxFile) Render(v interface{}) error {
    return renderTemplateDocument(f.document, v)
}

// RenderHeader (SimpleDocxFile) - рендер заголовка шаблона
func (f *SimpleDocxFile) RenderHeader(index int, v interface{}) error {
    pos := 0
    for _, header := range f.headers {
        if header != nil {
            if pos == index {
                return renderTemplateHeader(header, v)
            }
            pos++
        }
    }
    return nil
}

// Write (SimpleDocxFile)
func (f *SimpleDocxFile) Write(writer io.Writer) error {
    if f.zipFile != nil {
        if f.document != nil {
            w := zip.NewWriter(writer)
            defer w.Close()
            // Перебор файлов в Zip архиве        
            for _, zf := range f.zipFile.File {
                if zf != nil {                    
                    // Загрузка документа
                    if zf.Name == "word/document.xml" {
                        wzf, _ := w.Create(zf.Name)
                        if wzf != nil {
                            if b, err := wordDocumentToXML(f.document); b != nil && err == nil {                                
                                wzf.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>"))                                
                                wzf.Write(b)
                            }
                        }
                    } else if header, ok := f.headers[zf.Name]; ok && strings.Index(zf.Name, "word/header") >= 0 {
                        wzf, _ := w.Create(zf.Name)
                        if wzf != nil {
                            if b, err := wordHeaderToXML(header); b != nil && err == nil {                                
                                wzf.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>"))                                
                                wzf.Write(b)
                            }
                        }
                    } else {        
                        r, _ := zf.Open()                
                        if r != nil {
                            wzf, _ := w.Create(zf.Name)                        
                            if wzf != nil {
                                b, err := ioutil.ReadAll(r)
                                if err == nil {                                    
                                    wzf.Write(b)
                                }
                            }
                            r.Close()
                        }
                    }
                }
            } 
            err := w.Flush()
            if err != nil {
                return err
            }
            return nil
        }
        return errors.New("Not valid document")
    }
    return errors.New("Not loaded file")
}

// Save (SimpleDocxFile) - сохранить
func (f *SimpleDocxFile) Save(fileName string) error {
    if f.zipFile != nil {
        if f.document != nil {
            file, err := os.Create(fileName)
            if err != nil {
                return nil
            }  
            defer file.Close()
            w := zip.NewWriter(file)
            defer w.Close()

            // Перебор файлов в Zip архиве        
            for _, zf := range f.zipFile.File {
                if zf != nil {                    
                    // Загрузка документа
                    if zf.Name == "word/document.xml" {
                        wzf, _ := w.Create(zf.Name)
                        if wzf != nil {
                            if b, err := wordDocumentToXML(f.document); b != nil && err == nil {                                
                                wzf.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>"))                                
                                wzf.Write(b)
                            }
                        }
                    } else {        
                        r, _ := zf.Open()                
                        if r != nil {
                            wzf, _ := w.Create(zf.Name)                        
                            if wzf != nil {
                                b, err := ioutil.ReadAll(r)
                                if err == nil {                                    
                                    wzf.Write(b)
                                }
                            }
                            r.Close()
                        }
                    }
                }
            } 
            err = w.Flush()
            if err != nil {
                return err
            }
            return nil
        }
        return errors.New("Not valid document")
    }
    return errors.New("Not loaded file")
}

func wordHeaderToXML(h *Header) (data []byte, err error) {
    if h != nil {
        var buffer bytes.Buffer
        writer := bufio.NewWriter(&buffer)
        err = h.Encode(writer)
        if err == nil && buffer.Len() > 0 {
            data = buffer.Bytes()
            buffer.Reset()
            // Замены            
            data = bytes.Replace(data, []byte(" Ignorable="), []byte(" mc:Ignorable="), 1)
            data = bytes.Replace(data, []byte(" id="), []byte(" r:id="), -1) 
            // Замены empty tags
            for _, emptyTag := range emptyTags {
                data = bytes.Replace(data, []byte("></"+emptyTag+">"), []byte(" />"), -1)
            }
            data = rxStartTags.ReplaceAll(data, []byte("<w:$1"))
            data = rxEndTags.ReplaceAll(data, []byte("</w:$1"))
            data = rxXMLNsVals.ReplaceAll(data, []byte(" xmlns:$1=\"http://schemas."))
            data = rxURnVals.ReplaceAll(data, []byte(" xmlns:$1=\"urn:"))
            data = rxVals.ReplaceAll(data, []byte(" w:$1="))            
        }
    }
    return
}

func wordDocumentToXML(d *Document) (data []byte, err error) {
    if d != nil {
        var buffer bytes.Buffer
        writer := bufio.NewWriter(&buffer)
        err = d.Encode(writer)
        if err == nil && buffer.Len() > 0 {
            data = buffer.Bytes()
            buffer.Reset()
            // Замены            
            data = bytes.Replace(data, []byte(" Ignorable="), []byte(" mc:Ignorable="), 1)
            data = bytes.Replace(data, []byte(" id="), []byte(" r:id="), -1) 
            // Замены empty tags
            for _, emptyTag := range emptyTags {
                data = bytes.Replace(data, []byte("></"+emptyTag+">"), []byte(" />"), -1)
            }
            data = rxStartTags.ReplaceAll(data, []byte("<w:$1"))
            data = rxEndTags.ReplaceAll(data, []byte("</w:$1"))
            data = rxXMLNsVals.ReplaceAll(data, []byte(" xmlns:$1=\"http://schemas."))
            data = rxURnVals.ReplaceAll(data, []byte(" xmlns:$1=\"urn:"))
            data = rxVals.ReplaceAll(data, []byte(" w:$1="))            
        }
    }
    return
}