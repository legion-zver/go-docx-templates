package docx

import (
    "fmt"
    "regexp"
    "errors"
    "strings"  
    "reflect"  
    "github.com/aymerick/raymond"
    "github.com/legion-zver/go-docx-templates/graph"
)

var (
    rxTemplateItem  = regexp.MustCompile(`\{\{\s*([\w|\.]+)\s*\}\}`)
)

// Функционал шаблонизатора
func renderTemplateDocument(document *Document, v interface{}) error {
    if document != nil {
        // Проходимся по структуре документа
        for _, item := range document.Body.Items {
            if err := renderDocItem(item, v); err != nil {
                return err
            }
        }
        return nil
    }
    return errors.New("Not valid template document")
}

// Рендер элемента документа
func renderDocItem(item DocItem, v interface{}) error {
    switch elem := item.(type) {
        // Параграф
        case *ParagraphItem: {
            for _, i := range elem.Items {
                if err := renderDocItem(i, v); err != nil {
                    return err
                }
            }
        }
        // Запись
        case *RecordItem: {
            if len(elem.Text) > 0 {
                fmt.Println(elem.Text)
                out, err := raymond.Render(modeTemplateText(elem.Text), v)
                if err != nil {
                    return err
                }
                elem.Text = out
            }
        }
        // Таблица
        case *TableItem: {            
            for rowIndex, row := range elem.Rows {
                if row != nil {
                    // Если массив
                    if obj, ok := haveArrayInRow(row, v); ok {
                        lines       := objToLines(obj)
                        template    := row.Clone()
                        currentRow  := row
                        index       := rowIndex
                        for _, line := range lines {                            
                            if currentRow == nil {
                                currentRow = template.Clone()
                                // Insert Row
                                elem.Rows = append(elem.Rows[:index], append([]*TableRow{currentRow}, elem.Rows[index:]...)...)
                            }
                            if err := renderRow(currentRow, &line); err != nil {
                                return err
                            }
                            currentRow = nil
                            index++
                        }
                        template = nil
                        continue
                    }
                    // Если нет
                    if err := renderRow(row, v); err != nil {
                        return err
                    }
                }
            }
            // После обхода таблицы проходимся по ячейкам и проверяем merge флаги
        }
    }
    return nil
}

func objToLines(v interface{}) []map[string]interface{} {
    node := new(graph.Node)
    node.FromObject(v)
    return node.ListMap()
}

// renderRow - вывод строки таблицы
func renderRow(row *TableRow, v interface{}) error {
    if row != nil {
        for _, cell := range row.Cells {
            if cell != nil {
                for _, item := range cell.Items {
                    if err := renderDocItem(item, v); err != nil {
                        return err
                    }
                }
            }
        }
    }
    return nil
}

// Модифицируем текст шаблона
func modeTemplateText(tpl string) string {    
    tpl = strings.Replace(tpl, "{{", "{{{", -1)
	tpl = strings.Replace(tpl, "}}", "}}}", -1)
    tpl = strings.Replace(tpl,".","_",-1)    
    return strings.Replace(tpl,":length","_length",-1) 
}

// haveArrayInRow - содержится ли массив в строке
func haveArrayInRow(row *TableRow, v interface{}) (interface{}, bool) {
    if row != nil {
        for _, cell := range row.Cells {
            if match := rxTemplateItem.FindStringSubmatch(plainTextFromTableCell(cell)); match != nil && len(match) > 1 {                
                names := strings.Split(match[1], ".")
                if len(names) > 0 {
                    t   := reflect.TypeOf(v)
                    val := reflect.ValueOf(v)
                    var lastVal reflect.Value 
                    for _, name := range names {
                        t      := findType(t, name)
                        val, _ := findValue(val, name)
                        if t != nil {
                            if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
                                if lastVal.IsValid() {                                    
                                    return lastVal.Interface(), true                                                       
                                }
                                return v, true                                
                            }
                        } else {
                            break
                        }                        
                        lastVal = val                        
                    }
                }
            }
        } 
    }
    return nil, false
}

// Простой текс у ячейки
func plainTextFromTableCell(cell *TableCell) string {
    var result string
    if cell != nil {
        for _, item := range cell.Items {
            result += item.PlainText()            
        }
    }
    return result
}

// findType - получаем тип по имени
func findType(t reflect.Type, name string) reflect.Type {
    kind := t.Kind()
    // Если это ссылка, то получаем истенный тип
    if kind == reflect.Ptr || kind == reflect.Interface {
        t = t.Elem()
    }
    kind = t.Kind()
    if kind == reflect.Struct {
        if field, ok := t.FieldByName(name); ok {
            return field.Type
        }
    } 
    return nil
}

// findValue - получаем значение по имени
func findValue(v reflect.Value, name string) (reflect.Value, bool) {
    kind := v.Type().Kind()
    // Если это ссылка, то получаем истенный тип
    if kind == reflect.Ptr || kind == reflect.Interface {
        v = v.Elem()
    }
    kind = v.Type().Kind()
    if kind == reflect.Struct {
        v := v.FieldByName(name)
        if v.IsValid() {
            return v, true
        }        
    } 
    return v, false
}