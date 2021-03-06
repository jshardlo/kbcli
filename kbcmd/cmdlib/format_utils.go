package cmdlib

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// defaultListIndent - default indent level for lists
const defaultIndent = "  "

// Converts to generic interface that jsonpath can work with
func toGenericObject(v interface{}) (interface{}, error) {
	var jsonData interface{}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal. %v", err)
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal. %v", err)
	}
	return jsonData, nil
}

// getFieldValue returns value of field in an object through reflection
func getFieldValue(v interface{}, fieldPath string) (interface{}, error) {
	for _, fieldName := range strings.Split(fieldPath, ".") {
		val := reflect.ValueOf(v)
		if v == nil || val.IsNil() {
			return nil, nil
		}

		if val.Type().Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Type().Kind() != reflect.Struct {
			return nil, fmt.Errorf("invalid type %s for path %s", val.Type().String(), fieldName)
		}

		var found bool
		for i := 0; i < val.Type().NumField(); i++ {
			if val.Type().Field(i).Name == fieldName {
				val = val.Field(i)
				v = val.Interface()
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("unable to find field %s", fieldName)
		}
	}

	return v, nil
}

// computeMaxColumnSize computes the max column size
func computeMaxColumnSize(o Output, includeHeader bool) []int {

	colMaxSize := make([]int, len(o.Columns))
	for _, r := range o.Rows {
		for i, c := range r.Values {
			if len(c) > colMaxSize[i] {
				colMaxSize[i] = len(c)
			}
			if includeHeader && len(o.Columns[i]) > colMaxSize[i] {
				colMaxSize[i] = len(o.Columns[i])
			}
		}
	}
	return colMaxSize
}

// printSingleRow prints single row from given column values.
func printSingleRow(values []string, maxColSize []int, fo FormatOptions) string {
	var result string
	// write column headers
	for i, c := range values {
		result += c
		if i < len(values)-1 {
			result += strings.Repeat(" ", maxColSize[i]-len(c)+1)
		}
	}

	return result
}

// printList converts structured input into simple list of key/value pairs
func printList(out Output, fo FormatOptions, indent string) ([]string, error) {
	if len(out.Rows) == 0 {
		return nil, nil
	}
	var result []string
	if !fo.NoHeader {
		result = append(result, fmt.Sprintf("%s%s:", indent, out.Title))
	}

	indent += defaultIndent

	// Compute the header value that has the longest size.
	var maxHeaderLength int
	for _, c := range out.Columns {
		if len(c) > maxHeaderLength {
			maxHeaderLength = len(c)
		}
	}

	for _, r := range out.Rows {
		for i, v := range r.Values {
			reqIndent := strings.Repeat(" ", maxHeaderLength-len(out.Columns[i])+1)
			result = append(result, fmt.Sprintf("%s%s:%s%s", indent, out.Columns[i], reqIndent, v))
		}
		for _, child := range r.Children {
			childResult, err := printList(child, fo, indent)
			if err != nil {
				return nil, err
			}
			result = append(result, childResult...)
		}
		result = append(result, "")
	}
	return result, nil
}

// printColumns converts structured input into set of simple lines that is ready for printing.
func printColumns(out Output, fo FormatOptions, indent string) ([]string, error) {
	if len(out.Rows) == 0 {
		return nil, nil
	}

	var colMaxSize []int
	colMaxSize = computeMaxColumnSize(out, !fo.NoHeader)

	var rawRows []string

	if !fo.NoHeader {
		rawRows = append(rawRows, indent+printSingleRow(out.Columns, colMaxSize, fo))
	}

	for rowIndex, r := range out.Rows {
		row := printSingleRow(r.Values, colMaxSize, fo)

		rawRows = append(rawRows, indent+row)

		// Process sub items
		if len(r.Children) > 0 && fo.Type == FormatTypeTabular {
			for _, so := range r.Children {
				subItemIndent := indent + "  "

				// Print sub item title
				rawRows = append(rawRows, "")
				rawRows = append(rawRows, subItemIndent+"@"+so.Title+":")

				// Generate the rows
				subRows, err := printColumns(so, fo, subItemIndent)
				if err != nil {
					return nil, err
				}
				rawRows = append(rawRows, subRows...)
			}
			// Reprint the header if there are sub items.
			if !fo.NoHeader && rowIndex < len(out.Rows)-1 {
				rawRows = append(rawRows, "")
				rawRows = append(rawRows, indent+printSingleRow(out.Columns, colMaxSize, fo))
			}
		}
	}

	return rawRows, nil
}

// getFormatter returns formatter for the given type.
// If formatter is not registered, it will return the default formatter.
func getFormatter(log Logger, v interface{}) Formatter {
	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Slice {
		tp = tp.Elem()
	}
	f, ok := formatRegistry[tp]
	if ok {
		return f
	}

	log.Warningf("formatter for type %s not found, returning default", reflect.TypeOf(v))
	return Formatter{
		CustomFn: func(v interface{}, fo FormatOptions) Output {
			contents := MarshalJSON(v)
			var rows []OutputRow
			for _, r := range strings.Split(contents, "\n") {
				rows = append(rows, OutputRow{
					Values: []string{r},
				})
			}
			return Output{
				Columns: []string{"RAW"},
				Rows:    rows,
			}
		},
		Columns: []Column{
			{
				Name: "RAW",
			},
		},
	}
}
