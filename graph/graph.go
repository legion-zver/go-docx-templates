package graph

import (
    "fmt"
    "strconv"
    "reflect"    
)

// Node - элeмент графа
type Node struct {
    name     string
    values   map[string]interface{}    
    nodes []*Node
}

func (n *Node) String() string {    
    return fmt.Sprintln(n.ListMap()) 
}

// ListMap - перевод в лист
func (n *Node) ListMap() []map[string]interface{} {
    var m = make([]map[string]interface{}, 0)
    n.toListMap("", make(map[string]interface{}), &m)
    return m
} 

func (n *Node) toListMap(name string, current map[string]interface{}, out *[]map[string]interface{}) {
    for key, value := range n.values {
        if len(n.name) > 0 {
            key = n.name + "_" + key
        } 
        if len(name) > 0 {
            key = name + "_" + key
        }
        current[key] = value
    }
    if len(n.nodes) > 0 {
        path := name
        if len(n.name) > 0 {
            if len(path) > 0 {
                path += "_"
            } 
            path += n.name
        }
        // Проходимся и формируем счетчик
        var counter = make(map[string]int)        
        for _, node := range n.nodes {
            nname := path
            if len(nname) > 0 {
                if len(node.name) > 0 {
                    nname += "_"
                }                 
            }
            nname += node.name            
            if cv, ok := counter[nname+"_length"]; ok {
                counter[nname+"_length"] = cv + 1
            } else { 
                counter[nname+"_length"] = 1
            }
        }
        for key, cv := range counter {
            current[key] = cv
        }
        // Формируем строки
        for _, node := range n.nodes {
            var nw = make(map[string]interface{})
            for k,v := range current {
                nw[k] = v
            }       
            node.toListMap(path, nw, out)
        }
    } else {
        *out = append(*out, current)
    }
}

// FromObject (Node) - разложение объекта на граф
func (n* Node) FromObject(obj interface{}) {
    if n.values == nil {
        n.values = make(map[string]interface{}, 0)
    }
    if n.nodes == nil {
        n.nodes = make([]*Node, 0)
    }
    val := reflect.ValueOf(obj)
    kind := val.Type().Kind()
    if kind == reflect.Ptr || kind == reflect.Interface {
        val  = val.Elem()
        kind = val.Type().Kind()
    }
    if kind == reflect.Struct {
        t := val.Type()
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            fv := val.FieldByIndex(field.Index)
            kind = fv.Type().Kind()
            if kind == reflect.Ptr || kind == reflect.Interface {
                fv  = fv.Elem()
                kind = fv.Type().Kind()
            }
            if kind == reflect.Map {
                node := new(Node)
                node.name = field.Name
                node.FromObject(fv.Interface())
                n.nodes = append(n.nodes, node)
            } else if kind == reflect.Array || kind == reflect.Slice {                
                for j := 0; j < fv.Len(); j++ {
                    node := new(Node)
                    node.name = field.Name
                    node.FromObject(fv.Index(j).Interface())
                    n.nodes = append(n.nodes, node)
                }
            } else {
                n.values[field.Name] = val.FieldByIndex(field.Index).Interface()
            }
        }
    } else if kind == reflect.Map {
        for _, key := range val.MapKeys() {
            mapItem := val.MapIndex(key)
            kind = mapItem.Type().Kind()
            if kind == reflect.Ptr || kind == reflect.Interface {
                mapItem  = mapItem.Elem()
                kind = mapItem.Type().Kind()
            }
            if kind == reflect.Map {
                node := new(Node)
                node.name = key.String()
                node.FromObject(mapItem.Interface())
                n.nodes = append(n.nodes, node)
            } else if kind == reflect.Array || kind == reflect.Slice {                
                for j := 0; j < mapItem.Len(); j++ {
                    node := new(Node)
                    node.name = key.String()
                    node.FromObject(mapItem.Index(j).Interface())
                    n.nodes = append(n.nodes, node)
                }
            } else {
                n.values[key.String()] = mapItem.Interface()
            }
        }
    } else if kind == reflect.Array || kind == reflect.Slice {
        for i := 0; i < val.Len(); i++ {
            node := new(Node)
            node.name = n.name
            if len(node.name) > 0 {
                node.name += "_"+strconv.FormatInt(int64(i),10)
            }
            node.FromObject(val.Index(i).Interface())
            n.nodes = append(n.nodes, node)
        }
    }
}