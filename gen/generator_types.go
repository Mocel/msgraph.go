package main

import (
	"fmt"
	"strings"
)

const (
	colPrefix = "Collection("
)

// TODO: instead of current method, store
//       each namespace under its own structure
var nsPrefixes = map[string]struct{}{
	"microsoft.graph.": {},
	"graph.":           {},
}

var reservedTypeTable = map[string]string{
	"Edm.Boolean":          "bool",
	"Edm.Byte":             "byte",
	"Edm.Int16":            "int",
	"Edm.Int32":            "int",
	"Edm.Int64":            "int",
	"Edm.Decimal":          "int",
	"Edm.Single":           "float64",
	"Edm.Double":           "float64",
	"Edm.Binary":           "Binary",
	"Edm.Stream":           "Stream",
	"Edm.Guid":             "UUID",
	"Edm.String":           "string",
	"Edm.DateTimeOffset":   "time.Time",
	"Edm.Duration":         "Duration",
	"Edm.TimeOfDay":        "TimeOfDay",
	"Edm.Date":             "Date",
	"microsoft.graph.Json": "json.RawMessage",
}

func ptrType(t string) string {
	switch t {
	case "json.RawMessage":
		return t
	}
	return "*" + t
}

func stripNS(t string) (string, bool) {
	for ns := range nsPrefixes {
		if strings.HasPrefix(t, ns) {
			return t[len(ns):], true
		}
	}
	return t, false
}

func formatNS(t string) (string, bool) {
	s, ok := stripNS(t)
	return strings.Replace(s, ".", "", -1), ok
}

func formatNSUpper(t string) (string, bool) {
	s, ok := stripNS(t)
	arr := strings.Split(s, ".")
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Title(arr[i])
	}
	return strings.Join(arr, ""), ok
}

func exported(n string) string {
	return lintName(strings.Title(n))
}

func isCollectionType(t string) bool {
	return strings.HasPrefix(t, colPrefix)
}

func (g *Generator) AddSymTypeMapping(t string, sym string) {
	t, _ = stripNS(t)
	g.SymTypeMap[t] = sym
}

func (g *Generator) GetSymTypeMapping(t string) (string, bool) {
	t, _ = stripNS(t)
	sym, ok := g.SymTypeMap[t]
	return sym, ok
}

func (g *Generator) SymBaseType(t string) string {
	if x, ok := g.GetSymTypeMapping(t); ok {
		return x
	}
	if x, ok := formatNSUpper(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix) : len(t)-1])
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) SymFromType(t string) string {
	if x, ok := g.GetSymTypeMapping(t); ok {
		return x
	}
	if x, ok := formatNSUpper(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix):len(t)-1]) + "Collection"
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) TypeFromType(t string) string {
	if x, ok := reservedTypeTable[t]; ok {
		return ptrType(x)
	}
	if x, ok := g.GetSymTypeMapping(t); ok {
		return ptrType(x)
	}
	if x, ok := formatNSUpper(t); ok {
		return ptrType(exported(x))
	}
	if strings.HasPrefix(t, colPrefix) {
		return "[]" + g.TypeFromType(t[len(colPrefix) : len(t)-1])[1:]
	}
	panic(fmt.Errorf("Unknown type %s", t))
}
