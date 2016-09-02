package docxt

import (
    "github.com/legion-zver/go-docx-templates/docx"
)

// DocxTemplateFile - файл шаблонизатора
type DocxTemplateFile struct {
    file *docx.SimpleDocxFile
}

// OpenTemplate - open template
func OpenTemplate(fileName string) (*DocxTemplateFile, error) {
    f, err := docx.OpenFile(fileName)
    if err != nil {
        return nil, err
    }
    return &DocxTemplateFile{file:f}, nil
}

// Save (DocxTemplateFile)
func (t *DocxTemplateFile) Save(fileName string) error {
    return t.file.Save(fileName)
}