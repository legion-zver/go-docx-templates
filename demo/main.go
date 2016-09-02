package main

import (
    "fmt"    
    "github.com/legion-zver/go-docx-templates"
)

func main() {
    template, err := docxt.OpenTemplate("./example.docx")
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := template.Save("result.docx"); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Success")
}