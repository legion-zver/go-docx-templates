package docx

import (
    "errors"
    "encoding/xml"
)

// ParagraphItem - параграф
type ParagraphItem struct {
    Params  ParagraphParams `xml:"pPr"`
    Items   []DocItem
}

// ParagraphParams - параметры параграфа
type ParagraphParams struct {
    Style   *StringValue    `xml:"pStyle,omitempty"`
    Bidi    *IntValue       `xml:"bidi,omitempty"`
}

// Tag - имя тега элемента
func (item *ParagraphItem) Tag() string {
    return "p"
}

// Type - тип элемента
func (item *ParagraphItem) Type() DocItemType {
    return Paragraph
}

// Декодирование параграфа
func (item *ParagraphItem) decode(decoder *xml.Decoder) error {
    if decoder != nil {
        var end bool
        for !end {
            token, _ := decoder.Token()
            if token == nil {
                break
            }
            switch element := token.(type) {
                case xml.StartElement: {
                    if element.Name.Local == "pPr" {
                        decoder.DecodeElement(&item.Params, &element)
                    } else {
                        i := decodeItem(&element, decoder)
                        if i != nil {
                            item.Items = append(item.Items, i)
                        }
                    }
                }
                case xml.EndElement: {
                    if element.Name.Local == "p" {
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

// Кодирование параграфа
func (item *ParagraphItem) encode(encoder *xml.Encoder) error {
    if encoder != nil {
        // Начало параграфа
        start := xml.StartElement{Name:xml.Name{Local:item.Tag()}}        
        if err := encoder.EncodeToken(start); err != nil {
            return err
        }
        // Параметры параграфа
        if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name:xml.Name{Local:"pPr"}}); err != nil {
            return err
        }
        // Кодируем составные элементы
        for _, i := range item.Items {
            if err := i.encode(encoder); err != nil {
                return err
            }
        }
        // Конец параграфа        
        if err := encoder.EncodeToken(start.End()); err != nil {
            return err
        }        
        return encoder.Flush()
    }
    return errors.New("Not have encoder")
}