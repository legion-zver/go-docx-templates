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

/* КОДИРОВАНИЕ */

// Encode - кодирование
func (h *Header) Encode(writer io.Writer) error {
    encoder := xml.NewEncoder(writer)
    if encoder != nil {
        // Начало документа
        var attrs = make([]xml.Attr, 0)
        for key, val := range h.Scheme {
            attrs = append(attrs, xml.Attr{Name:xml.Name{Local:key}, Value: val})
        } 
        if len(h.SkipScheme) > 0 {
            attrs = append(attrs, xml.Attr{Name:xml.Name{Local:"Ignorable"}, Value: h.SkipScheme})
        }
        hStart := xml.StartElement{Name: xml.Name{Local:"hdr"}, Attr: attrs}
        err := encoder.EncodeToken(hStart)
        if err != nil {
            return err
        }
        // Отдаем кодирование глубже - элементам        
        for _, item := range h.Items {            
            if err := item.encode(encoder); err != nil {
                return err
            }
        }
        // Конец документа
        err = encoder.EncodeToken(hStart.End())
        if err != nil {
            return err
        }        
        return encoder.Flush()
    }
    return errors.New("Error create encoder")
}