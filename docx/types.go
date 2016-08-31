package docx

import (
    "encoding/xml"
)

// Document - Документ DOCX
type Document struct {
    // Заголовок
    XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main document"`
    XMLNsWP     string   `xml:"wp,attr,omitempty"`
    XMLNsW10    string   `xml:"w10,attr,omitempty"`
    XMLNsW14    string   `xml:"w14,attr,omitempty"`
    XMLNsWPS    string   `xml:"wps,attr,omitempty"`
    XMLNsWPG    string   `xml:"wpg,attr,omitempty"`
    XMLNsMC     string   `xml:"mc,attr,omitempty"`
    XMLNsV      string   `xml:"v,attr,omitempty"`
    XMLNsO      string   `xml:"o,attr,omitempty"`
    Ignorable   string   `xml:"Ignorable,attr,omitempty"`
    // Тело
    Body       *Body     `xml:"body"`
}

// Body in Document
type Body struct {
    Paragraphs []*Paragraph `xml:"p,omitempty"`
    Params       *SectionPr `xml:"sectPr,omitempty"`
}

// SectionPr - params body
type SectionPr struct {
    PageSize        *XMLSize        `xml:"pgSz,omitempty"`
    PageMargin      *XMLMargin      `xml:"pgMar,omitempty"`
    HeaderReference *XMLReference   `xml:"headerReference,omitempty"`
    FooterReference *XMLReference   `xml:"footerReference,omitempty"`
    Bidi            *XMLIntValue    `xml:"bidi,omitempty"`
}

// Paragraph in Body
type Paragraph struct {
    Params  *ParagraphPr `xml:"pPr,omitempty"`
    Records []*Record    `xml:"r,omitempty"`
}

// ParagraphPr - parameters paragraph
type ParagraphPr struct {
    Style   *XMLValue      `xml:"pStyle,omitempty"`
    Bidi    *XMLIntValue   `xml:"bidi,omitempty"`
}

// Record - запись
type Record struct {
    Params  *RecordPr `xml:"rPr,omitempty"`
    Text    string    `xml:"t,omitempty"`
    Break   *Break    `xml:"br,omitempty"`
}

// Break - <br/>
type Break struct {
    Type string `xml:"type,omitempty"`
}

// RecordPr - parameters record
type RecordPr struct {
    Rtl    *XMLIntValue         `xml:"rtl,omitempty"`
    Lang   *XMLStringValue      `xml:"lang,omitempty"`
}

// XMLReference - reference value
type XMLReference struct {
    Type    string `xml:"type,attr,omitempty"`
    ID      string `xml:"id,attr,omitempty"`
}

// XMLMargin - margin значение
type XMLMargin struct {    
    Top       int64     `xml:"top,attr"`
    Bottom    int64     `xml:"bottom,attr"`
    Left      int64     `xml:"left,attr"`
    Right     int64     `xml:"right,attr"`
    Header    int64     `xml:"header,attr"`
    Footer    int64     `xml:"footer,attr"`
}

// XMLSize - size значение
type XMLSize struct {    
    Width     int64     `xml:"w,attr"`
    Height    int64     `xml:"h,attr"`
    Orient    string    `xml:"orient,attr,omitempty"`
}

// XMLValueInfo - информация о значении
type XMLValueInfo struct {
    Name     string `xml:"name,attr,omitempty"`
    URI      string `xml:"uri,attr,omitempty"`
    Language string `xml:"lang,attr,omitempty"`    
}

// XMLValue - одиночное значение
type XMLValue struct {
    XMLValueInfo
    Value    string `xml:"val,attr"`
}

// XMLStringValue - одиночное string значение
type XMLStringValue struct {    
    Value    string `xml:"val,attr"`
}

// XMLBoolValue - одиночное int значение
type XMLBoolValue struct {    
    Value    bool `xml:"val,attr"`
}

// XMLIntValue - одиночное int значение
type XMLIntValue struct {    
    Value    int64 `xml:"val,attr"`
}

// XMLFloatValue - одиночное float значение
type XMLFloatValue struct {    
    Value    float64 `xml:"val,attr"`
}