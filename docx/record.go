package docx

import (    
    "errors"
    "encoding/xml"
)

// RecordItem - record item
type RecordItem struct {
    Params  RecordParams  `xml:"rPr,omitempty"`
    Text    string        `xml:"t"`
    Break   bool          `xml:"br,omitempty"`
}

// RecordParams - params record
type RecordParams struct {
    Fonts  *RecordFonts      `xml:"rFonts,omitempty"`
    Rtl    *IntValue         `xml:"rtl,omitempty"`
    Lang   *StringValue      `xml:"lang,omitempty"`
}

// RecordFonts - fonts in record
type RecordFonts struct {
    ASCII         string `xml:"ascii,attr"`
    CS            string `xml:"cs,attr"`
    HandleANSI    string `xml:"hAnsi,attr"`
    EastAsia      string `xml:"eastAsia,attr"`
    HandleInt     string `xml:"hint,attr,omitempty"`
}

// Tag - имя тега элемента
func (item *RecordItem) Tag() string {
    return "r"
}

// Type - тип элемента
func (item *RecordItem) Type() DocItemType {
    return Record
}

// Декодирование записи
func (item *RecordItem) decode(decoder *xml.Decoder) error {
    if decoder != nil {
        var end bool 
        for !end {
            token, _ := decoder.Token()
            if token == nil {
                break
            }
            switch element := token.(type) {
                case xml.StartElement: {
                    if element.Name.Local == "rPr" {
                        decoder.DecodeElement(&item.Params, &element)
                    } else if element.Name.Local == "t" {
                        decoder.DecodeElement(&item.Text, &element)
                    } else if element.Name.Local == "br" {
                        item.Break = true
                    }
                }
                case xml.EndElement: {
                    if element.Name.Local == "r" {
                        end = true
                    }
                }
            }
        }                   
        return nil
    }
    return errors.New("Not have decoder")
}

/* КОДИРОВАНИЕ */

// Кодирование записи
func (item *RecordItem) encode(encoder *xml.Encoder) error {
    if encoder != nil {
        // Начало записи        
        start := xml.StartElement{Name:xml.Name{Local:item.Tag()}}
        if err := encoder.EncodeToken(start); err != nil {
            return err
        }
        // Параметры записи
        if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name:xml.Name{Local:"rPr"}}); err != nil {
            return err
        }
        // Текст
        if err := encoder.EncodeElement(&item.Text, xml.StartElement{Name:xml.Name{Local:"t"}}); err != nil {
            return err
        }
        // <br />
        if item.Break {                   
            startBr := xml.StartElement{Name:xml.Name{Local:"br"}} 
            if err := encoder.EncodeToken(startBr); err != nil {
                return err
            }
            if err := encoder.EncodeToken(startBr.End()); err != nil {
                return err
            }
            if err := encoder.Flush(); err != nil {
                return err
            }
        }
        // Конец записи        
        if err := encoder.EncodeToken(start.End()); err != nil {
            return err
        }        
        return encoder.Flush()
    }
    return errors.New("Not have encoder")
}