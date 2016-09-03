package docx

import (
    "errors"
    "encoding/xml"
)

// TableItem - элемент таблици
type TableItem struct {
    Params       TableParams   `xml:"tblPr"`
    Grid         TableGrid     `xml:"tblGrid"`
    Rows         []*TableRow   `xml:"tr,omitempty"`
}

// TableGrid - Grid table 
type TableGrid struct {
    Cols []*WidthValue `xml:"gridCol,omitempty"`
}

// TableParamsEx - Other params table 
type TableParamsEx struct {    
    Shadow   ShadowValue   `xml:"shd"`    
}

// Tag - имя тега элемента
func (item *TableItem) Tag() string {
    return "tbl"
}

// PlainText - текст
func (item *TableItem) PlainText() string {
    return ""
}

// Type - тип элемента
func (item *TableItem) Type() DocItemType {
    return Table
}

// Clone - клонирование
func (item *TableItem) Clone() DocItem {
    result := new(TableItem)
    result.Grid.Cols = make([]*WidthValue, 0)
    for _, col := range item.Grid.Cols {
        if col != nil {
            w := new(WidthValue)
            w.Type = col.Type
            w.Value = col.Value
            result.Grid.Cols = append(result.Grid.Cols, w)
        }
    }
    if item.Params.DocGrid != nil {
        result.Params.DocGrid = new(IntValue)
        result.Params.DocGrid.Value = item.Params.DocGrid.Value
    }
    if item.Params.Ind != nil {
        result.Params.Ind = new(WidthValue)
        result.Params.Ind.Type = item.Params.Ind.Type
        result.Params.Ind.Value = item.Params.Ind.Value
    }    
    if item.Params.Jc != nil {
        result.Params.Jc = new(StringValue)        
        result.Params.Jc.Value = item.Params.Jc.Value
    }
    if item.Params.Layout != nil {
        result.Params.Layout = new(TableLayout)
        result.Params.Layout.Type = item.Params.Layout.Type
    }
    if item.Params.Shadow != nil {
        result.Params.Shadow = new(ShadowValue)
        result.Params.Shadow.From(item.Params.Shadow)       
    }
    if item.Params.Width != nil {
        result.Params.Width = new(WidthValue)
        result.Params.Width.From(item.Params.Width)
    }
    if item.Params.Borders != nil {
        result.Params.Borders = new(TableBorders)
        result.Params.Borders.From(item.Params.Borders)    
    }
    // Клонирование строк
    result.Rows = make([]*TableRow, 0)
    for _, row := range item.Rows {
        if row != nil {
            result.Rows = append(result.Rows, row.Clone())
        }
    }
    return result
}

// TableParams - Params table 
type TableParams struct {
    Width    *WidthValue    `xml:"tblW,omitempty"`
    Jc       *StringValue   `xml:"jc,omitempty"`
    Ind      *WidthValue    `xml:"tblInd,omitempty"`
    Borders  *TableBorders  `xml:"tblBorders,omitempty"`
    Shadow   *ShadowValue   `xml:"shd,omitempty"`
    Layout   *TableLayout   `xml:"tblLayout,omitempty"`
    DocGrid  *IntValue      `xml:"docGrid,omitempty"`
}

// TableLayout - layout params
type TableLayout struct {
    Type string `xml:"type,attr"`
}

// TableBorders in table
type TableBorders struct {
    Top      TableBorder     `xml:"top"`    
    Left     TableBorder     `xml:"left"`
    Bottom   TableBorder     `xml:"bottom"`
    Right    TableBorder     `xml:"right"`
    InsideH  *TableBorder    `xml:"insideH,omitempty"`
    InsideV  *TableBorder    `xml:"insideV,omitempty"`
}

// From (TableBorders)
func (b *TableBorders) From(b1 *TableBorders) {
    if b1 != nil {
        b.Top.From(&b1.Top)
        b.Left.From(&b1.Left)
        b.Bottom.From(&b1.Bottom)
        b.Right.From(&b1.Right)
        if b1.InsideH != nil {
            b.InsideH = new(TableBorder)
            b.InsideH.From(b1.InsideH)
        }
        if b1.InsideV != nil {
            b.InsideV = new(TableBorder)
            b.InsideV.From(b1.InsideV)
        }
    }
}

// TableBorder in borders
type TableBorder struct {
    Value   string  `xml:"val,attr"`
    Color   string  `xml:"color,attr"`
    Size    int64   `xml:"sz,attr"`
    Space   int64   `xml:"space,attr"`
    Shadow  int64   `xml:"shadow,attr"`
    Frame   int64   `xml:"frame,attr"`
}

// From (TableBorder)
func (b *TableBorder) From(b1 *TableBorder) {
    if b1 != nil {
        b.Value = b1.Value
        b.Color = b1.Color
        b.Frame = b1.Frame
        b.Shadow = b1.Shadow
        b.Size = b1.Size
        b.Space = b1.Space
    }
}

// TableRow - row in table
type TableRow struct {
    OtherParams  *TableParamsEx     `xml:"tblPrEx,omitempty"`
    Params       TableRowParams     `xml:"trPr"`    
    Cells        []*TableCell       `xml:"tc,omitempty"`
}

// TableRowParams - row params
type TableRowParams struct {
    Height       HeightValue  `xml:"trHeight"`
    IsHeader     bool
}
 
// TableCell - table cell
type TableCell struct {
    Params     TableCellParams `xml:"tcPr"`
    Items      []DocItem
}

// TableCellParams - cell params
type TableCellParams struct {
    Width           *WidthValue     `xml:"tcW,omitempty"`
    Borders         *TableBorders   `xml:"tcBorders,omitempty"`
    Shadow          *ShadowValue    `xml:"shd,omitempty"`
    Margins         *Margins        `xml:"tcMar,omitempty"`
    VerticalAlign   *StringValue    `xml:"vAlign,omitempty"`
    VerticalMerge   *StringValue    `xml:"vMerge,omitempty"`
    GridSpan        *IntValue       `xml:"gridSpan,omitempty"`
    HideMark        *EmptyValue     `xml:"hideMark,omitempty"`
    NoWrap          *EmptyValue     `xml:"noWrap,omitempty"`
}

// Clone (TableCell) - клонирование ячейки
func (cell *TableCell) Clone() *TableCell {
    result := new(TableCell)
    if cell.Params.GridSpan != nil {
        result.Params.GridSpan = new(IntValue)
        result.Params.GridSpan.Value = cell.Params.GridSpan.Value 
    }
    if cell.Params.HideMark != nil {
        result.Params.HideMark = new(EmptyValue)         
    }
    if cell.Params.NoWrap != nil {
        result.Params.NoWrap = new(EmptyValue)         
    }
    if cell.Params.Shadow != nil {
        result.Params.Shadow = new(ShadowValue)
        result.Params.Shadow.From(cell.Params.Shadow)                 
    }
    if cell.Params.VerticalAlign != nil {
        result.Params.VerticalAlign = new(StringValue)
        result.Params.VerticalAlign.Value = cell.Params.VerticalAlign.Value 
    }
    if cell.Params.VerticalMerge != nil {
        result.Params.VerticalMerge = new(StringValue)
        result.Params.VerticalMerge.Value = cell.Params.VerticalMerge.Value 
    }
    if cell.Params.Margins != nil {
        result.Params.Margins = new(Margins)
        result.Params.Margins.From(cell.Params.Margins)
    }
    if cell.Params.Width != nil {
        result.Params.Width = new(WidthValue)
        result.Params.Width.From(cell.Params.Width)
    }
    if cell.Params.Borders != nil {
        result.Params.Borders = new(TableBorders)
        result.Params.Borders.From(cell.Params.Borders)
    }
    result.Items = make([]DocItem, 0)
    for _, item := range cell.Items {
        if item != nil {
            result.Items = append(result.Items, item.Clone())
        }
    }
    return result
}

// Clone (TableRow) - клонирование строки таблицы
func (row *TableRow) Clone() *TableRow {
    result := new(TableRow)
    result.Params = row.Params
    result.OtherParams = new(TableParamsEx)
    result.OtherParams.Shadow = row.OtherParams.Shadow
    // Клонируем ячейки
    result.Cells = make([]*TableCell, 0)
    for _, cell := range row.Cells {
        if cell != nil {
            result.Cells = append(result.Cells, cell.Clone())
        }
    }
    return result
}

/* ДЕКОДИРОВАНИЕ */

// Декодирование таблицы
func (item *TableItem) decode(decoder *xml.Decoder) error {
    if decoder != nil {
        item.Rows = make([]*TableRow, 0)
        var end bool 
        for !end {
            token, _ := decoder.Token()
            if token == nil {
                break
            }
            switch element := token.(type) {
                case xml.StartElement: {
                    if element.Name.Local == "tblPr" {
                        decoder.DecodeElement(&item.Params, &element)
                    } else if element.Name.Local == "tblGrid" {
                        decoder.DecodeElement(&item.Grid, &element)
                    } else if element.Name.Local == "tr" {                        
                        row := new(TableRow)
                        if row.decode(decoder) == nil {
                            item.Rows = append(item.Rows, row)
                        }
                    }
                }
                case xml.EndElement: {
                    if element.Name.Local == "tbl" {
                        end = true
                    }
                }
            }
        }        
        return nil
    }
    return errors.New("Not have decoder")
}

// Декодирование строк таблицы
func (row *TableRow) decode(decoder *xml.Decoder) error {
    if decoder != nil {
        row.Cells = make([]*TableCell, 0)
        var end bool 
        for !end {
            token, _ := decoder.Token()
            if token == nil {
                break
            }
            switch element := token.(type) {
                case xml.StartElement: {
                    if element.Name.Local == "trHeight" {
                        decoder.DecodeElement(&row.Params.Height, &element)
                    } else if element.Name.Local == "tblHeader" {
                        row.Params.IsHeader = true
                    } else if element.Name.Local == "tblPrEx" {
                        row.OtherParams = new(TableParamsEx)
                        decoder.DecodeElement(row.OtherParams, &element)
                    } else if element.Name.Local == "tc" {
                        cell := new(TableCell)
                        if cell.decode(decoder) == nil {
                            row.Cells = append(row.Cells, cell)
                        }
                    }
                }
                case xml.EndElement: {
                    if element.Name.Local == "tr" {
                        end = true
                    }
                }
            }
        }        
        return nil
    }
    return errors.New("Not have decoder")
}

// Декодирование ячеек таблицы
func (row *TableCell) decode(decoder *xml.Decoder) error {
    if decoder != nil {        
        var end bool 
        for !end {
            token, _ := decoder.Token()
            if token == nil {
                break
            }
            switch element := token.(type) {
                case xml.StartElement: {
                    if element.Name.Local == "tcPr" {
                        decoder.DecodeElement(&row.Params, &element)
                    } else {
                        i := decodeItem(&element, decoder)
                        if i != nil {
                            row.Items = append(row.Items, i)
                        }
                    }
                }
                case xml.EndElement: {
                    if element.Name.Local == "tc" {
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

// Кодирование таблицы
func (item *TableItem) encode(encoder *xml.Encoder) error {
    if encoder != nil {
        // Начало таблицы
        start := xml.StartElement{Name:xml.Name{Local:item.Tag()}}        
        if err := encoder.EncodeToken(start); err != nil {
            return err
        }
        // Параметры таблицы
        if err := encoder.EncodeElement(&item.Params, xml.StartElement{Name:xml.Name{Local:"tblPr"}}); err != nil {
            return err
        }
        // Сетка таблицы
        if err := encoder.EncodeElement(&item.Grid, xml.StartElement{Name:xml.Name{Local:"tblGrid"}}); err != nil {
            return err
        }
        // Строки таблицы
        for _, row := range item.Rows {
            if row != nil {
                if err := row.encode(encoder); err != nil {
                    return err
                }
            }
        }
        // Конец таблицы        
        if err := encoder.EncodeToken(start.End()); err != nil {
            return err
        }
        return encoder.Flush()
    }
    return errors.New("Not have encoder")
}

// Кодирование ячейки таблицы
func (cell *TableCell) encode(encoder *xml.Encoder) error {
    if encoder != nil {
        // Начало ячейки таблицы
        start := xml.StartElement{Name:xml.Name{Local:"tc"}}
        if err := encoder.EncodeToken(start); err != nil {
            return err
        }
        // Параметры ячейки таблицы
        if err := encoder.EncodeElement(&cell.Params, xml.StartElement{Name:xml.Name{Local:"tcPr"}}); err != nil {
            return err
        }        
        // Кодируем составные элементы
        for _, i := range cell.Items {
            if err := i.encode(encoder); err != nil {
                return err
            }
        }
        // Конец ячейки таблицы        
        if err := encoder.EncodeToken(start.End()); err != nil {
            return err
        }
        return encoder.Flush()
    }
    return errors.New("Not have encoder")
}

// Кодирование строки таблицы
func (row *TableRow) encode(encoder *xml.Encoder) error {
    if encoder != nil {
        // Начало строки таблицы
        start := xml.StartElement{Name:xml.Name{Local:"tr"}}        
        if err := encoder.EncodeToken(start); err != nil {
            return err
        }
        // Параметры строки таблицы
        if row.OtherParams != nil {
            if err := encoder.EncodeElement(row.OtherParams, xml.StartElement{Name:xml.Name{Local:"tblPrEx"}}); err != nil {
                return err
            }
        }
        // Кодируем Параметры
        startPr := xml.StartElement{Name:xml.Name{Local:"trPr"}}
        if err := encoder.EncodeToken(startPr); err != nil {
            return err
        }
        if err := encoder.EncodeElement(&row.Params.Height,xml.StartElement{Name:xml.Name{Local:"trHeight"}}); err != nil {
            return err
        }
        if row.Params.IsHeader {
            startHeader := xml.StartElement{Name:xml.Name{Local:"tblHeader"}}
            if err := encoder.EncodeToken(startHeader); err != nil {
                return err
            }
            if err := encoder.EncodeToken(startHeader.End()); err != nil {
                return err
            }
            if err := encoder.Flush(); err != nil {
                return err
            }
        }
        if err := encoder.EncodeToken(startPr.End()); err != nil {
            return err
        }
        if err := encoder.Flush(); err != nil {
            return err
        }
        // Кодируем ячейки
        for _, cell := range row.Cells {
            if cell != nil {
                if err := cell.encode(encoder); err != nil {
                    return err
                }
            }
        }
        // Конец строки таблицы        
        if err := encoder.EncodeToken(start.End()); err != nil {
            return err
        }
        return encoder.Flush()
    }
    return errors.New("Not have encoder")
}