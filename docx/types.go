package docx

// HeightValue - значение высоты
type HeightValue struct {
    Value         int64  `xml:"val,attr"`
    HeightRule    string `xml:"hRule,attr,omitempty"`
}

// WidthValue - значение длины
type WidthValue struct {
    Value    int64  `xml:"w,attr"`
    Type     string `xml:"type,attr,omitempty"`
} 

// SizeValue - значение размера
type SizeValue struct {    
    Width          int64     `xml:"w,attr"`
    Height         int64     `xml:"h,attr"`
    Orientation    string    `xml:"orient,attr,omitempty"`
}

// EmptyValue - пустое значение
type EmptyValue struct {        
}

// StringValue - одиночное string значение
type StringValue struct {    
    Value    string `xml:"val,attr,omitempty"`
}

// BoolValue - одиночное bool значение
type BoolValue struct {    
    Value    bool `xml:"val,attr"`
}

// IntValue - одиночное int значение
type IntValue struct {    
    Value    int64 `xml:"val,attr"`
}

// FloatValue - одиночное float значение
type FloatValue struct {    
    Value    float64 `xml:"val,attr"`
}

// ReferenceValue - reference value
type ReferenceValue struct {
    Type    string `xml:"type,attr"`
    ID      string `xml:"id,attr"`
}

// SpacingValue - spacing value
type SpacingValue struct {
    After     int64     `xml:"after,attr"`
    Before    int64     `xml:"before,attr"`
    Line      int64     `xml:"line,attr"` 
    LineRule  string    `xml:"lineRule,attr"`
}

// MarginValue - margin значение
type MarginValue struct {
    Top       int64     `xml:"top,attr"`    
    Left      int64     `xml:"left,attr"`
    Bottom    int64     `xml:"bottom,attr"`
    Right     int64     `xml:"right,attr"`
    Header    int64     `xml:"header,attr,omitempty"`
    Footer    int64     `xml:"footer,attr,omitempty"`
}

// Margins - margins значение
type Margins struct {
    Top     WidthValue `xml:"top"`    
    Left    WidthValue `xml:"left"`
    Bottom  WidthValue `xml:"bottom"`
    Right   WidthValue `xml:"right"`
}

// ShadowValue - значение тени
type ShadowValue struct {
    Value   string `xml:"val,attr"`
    Color   string `xml:"color,attr"`
    Fill    string `xml:"fill,attr"`
}