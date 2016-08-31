package docx

import (    
    "archive/zip"
    "encoding/xml"
)

// SimpleDocxFile - файл docx
type SimpleDocxFile struct {
    zipFile  *zip.ReadCloser
    document *Document
}

// OpenFile - Открытие файла DOCX
func OpenFile(fileName string) (*SimpleDocxFile,error) {
    z, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}    
    d := new(SimpleDocxFile)
    d.zipFile = z
    // Перебор файлов в Zip архиве
    for _, f := range z.File {
        if f != nil {
            // Загрузка документа
            if f.Name == "word/document" {
                d.document = new(Document)
                err := readXMLFromZipFile(f, d.document)
                if err != nil {
                    return nil, err
                }
            }
        }
    }    
    return d, nil 
}

// readXMLFromZipFile 
func readXMLFromZipFile(zipFile *zip.File, out interface{}) error {
    rc, err := zipFile.Open()
    if err != nil {
        return err
    }
    decoder := xml.NewDecoder(rc)
    err = decoder.Decode(out)
    if err != nil {
        return err
    }
    return nil
}