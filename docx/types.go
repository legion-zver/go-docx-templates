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
    Tables     []*Table     `xml:"tbl,omitempty"`
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

// Table in Body
type Table struct {
    Params       *TablePr       `xml:"tblPr,omitempty"`
    Grid         *TableGrid     `xml:"tblGrid,omitempty"`
    Rows        []*TableRow     `xml:"tr,omitempty"`
}

// TableRow - row in table
type TableRow struct {
    Params       *TableRowPr  `xml:"trPr,omitempty"`
    OtherParams  *TablePrEx   `xml:"tblPrEx,omitempty"`
    Cells       []*TableCell  `xml:"tc,omitempty"`
}

// TableRowPr - row params
type TableRowPr struct {
    Height       XMLHeightValue  `xml:"trHeight"`
    IsHeader    *XMLEmptyValue   `xml:"tblHeader,omitempty"` // if != nil -> isHeader
}
 
// TableCell - table cell
type TableCell struct {
    Params     *TableCellPr `xml:"tcPr,omitempty"`
    Paragraph  *Paragraph   `xml:"p,omitempty"`
}

// TableCellPr - cell params
type TableCellPr struct {
    Width    XMLWidthValue  `xml:"tcW"`
    Borders  TableBorders   `xml:"tcBorders"`
    Shadow   Shadow         `xml:"shd"`
    Margin   XMLMargin      `xml:"tcMar"`
    VAlign   XMLStringValue `xml:"vAlign"`
}

// TableGrid - Grid table 
type TableGrid struct {
    Cols []*XMLWidthValue `xml:"gridCol,omitempty"`
}

// TablePr - Params table 
type TablePr struct {
    Width    XMLWidthValue    `xml:"tblW"`
    Js       XMLStringValue   `xml:"js"`
    Ind      XMLWidthValue    `xml:"tblInd"`
    Shadow   Shadow           `xml:"shd"`
    Borders  TableBorders     `xml:"tblBorders"`
    Layout   TableLayout      `xml:"tblLayout"`
}

// TablePrEx - Ex Params table in rows
type TablePrEx struct {
    Shadow   Shadow           `xml:"shd"`
}

// TableLayout - layout params
type TableLayout struct {
    Type string `xml:"type,attr"`
}

// TableBorders in table
type TableBorders struct {
    Top      TableBorder     `xml:"top"`
    Bottom   TableBorder     `xml:"bottom"`
    Left     TableBorder     `xml:"left"`
    Right    TableBorder     `xml:"right"`
    InsideH *TableBorder     `xml:"insideH,omitempty"`
    InsideV *TableBorder     `xml:"insideV,omitempty"`
}

// TableBorder in borders
type TableBorder struct {
    Value   string  `xml:"val,attr"`
    Color   string  `xml:"color,attr"`
    Size    int64   `xml:"sv,attr"`
    Space   int64   `xml:"space,attr"`
    Shadow  int64   `xml:"shadow,attr"`
    Frame   int64   `xml:"frame,attr"`
}

// Shadow - shadow
type Shadow struct {
    Value   string `xml:"val,attr"`
    Color   string `xml:"color,attr"`
    Fill    string `xml:"fill,attr"`
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
    Fonts  *RecordFonts         `xml:"rFonts,omitempty"`
}

// RecordFonts - fonts in record
type RecordFonts struct {
    ASCII    string `xml:"ascii,attr"`
    CS       string `xml:"cs,attr"`
    HANSI    string `xml:"hAnsi,attr"`
    EastAsia string `xml:"eastAsia,attr"`
    HINT     string `xml:"hint,attr"`
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
    Header    *int64    `xml:"header,attr,omitempty"`
    Footer    *int64    `xml:"footer,attr,omitempty"`
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

// XMLEmptyValue - empty value
type XMLEmptyValue struct {
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

// XMLWidthValue - width value
type XMLWidthValue struct {    
    Value    int64  `xml:"w,attr"`
    Type     string `xml:"type,attr,omitempty"`
}

// XMLHeightValue - height value
type XMLHeightValue struct {    
    Value    int64  `xml:"val,attr"`
    HRule    string `xml:"hRule,attr,omitempty"`
}