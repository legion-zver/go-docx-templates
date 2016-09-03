package main

import (
    "fmt"    
    "github.com/legion-zver/go-docx-templates"
)

type TestStruct struct {
    FileName string
    Items []TestItemStruct
}

type TestItemStruct struct {
    Column1 string
    Column2 string
    Column3 string
    Column4 string
}

func main() {
    template, err := docxt.OpenTemplate("./example.docx")
    if err != nil {
        fmt.Println(err)
        return
    }  
    test := new(TestStruct)
    test.FileName = "example.docx"
    test.Items = []TestItemStruct{
        TestItemStruct{"1","2","3","4"},
        TestItemStruct{"2","3","4","1"},
        TestItemStruct{"3","4","1","2"},
        TestItemStruct{"4","1","2","3"},
    }    
    if err := template.Render(test); err != nil {
        fmt.Println(err)
        return
    }
    if err := template.Save("result.docx"); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Success")
}