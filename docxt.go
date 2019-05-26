package docxt

import (
    "io"
    "errors"
    "github.com/qida/go-docx-templates/docx"
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

// Write (DocxTemplateFile)
func (t *DocxTemplateFile) Write(w io.Writer) error {
    return t.file.Write(w)
}

// RenderTemplate (SimpleDocxFile) - рендер шаблона
func (t *DocxTemplateFile) RenderTemplate(v interface{}) error {
    if t.file != nil {
        return t.file.Render(v)
    }
    return errors.New("Not loading template file")
}

// RenderHeaderTemplate (SimpleDocxFile) - рендер шаблона
func (t *DocxTemplateFile) RenderHeaderTemplate(indexHeader int, v interface{}) error {
    if t.file != nil {
        return t.file.RenderHeader(indexHeader, v)
    }
    return errors.New("Not loading template file")
}