package docx

import (
    "io"    
    "errors"
    "encoding/xml"        
)

// Header - разметка заголовка DOCX
type Header struct {
    Scheme       map[string]string    
    SkipScheme   string
    Items      []DocItem    
}

/* ДЕКОДИРОВАНИЕ */

// Decode (Document) - декодирование документа 
func (h *Header) Decode(reader io.Reader) error {
    decoder := xml.NewDecoder(reader)
    if decoder != nil {
        h.Scheme = make(map[string]string)
        h.Items  = make([]DocItem, 0)
        for {
            token, _ := decoder.Token()		
            if token == nil {
                break
            }
            switch element := token.(type) {            
                case xml.StartElement: {
                    if element.Name.Local == "hdr" {
                        for _, attr := range element.Attr {
                            if attr.Name.Local == "Ignorable" {
                                h.SkipScheme = attr.Value
                            } else {
                                h.Scheme[attr.Name.Local] = attr.Value
                            }
                        }
                    } else {                        
                        item := decodeItem(&element, decoder)
                        if item != nil {
                            h.Items = append(h.Items, item)
                        }
                    }
                }
            }
        }
        return nil
    }
    return errors.New("Error create decoder")
}
