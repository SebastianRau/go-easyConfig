package templating

import (
	"bytes"
	"fmt"
	"strings"

	"text/template"

	"gopkg.in/yaml.v3"
)

func hasNext(array []interface{}, idx int) bool {
	return idx < len(array)-1
}
func isLast(array []interface{}, idx int) bool {
	return idx == len(array)-1
}
func arrayJoin(array []interface{}, separator string, addLast bool) string {
	var buf strings.Builder
	for i, v := range array {
		if i == len(array)-1 && !addLast {
			buf.WriteString(fmt.Sprintf("%v", v))
		} else {
			buf.WriteString(fmt.Sprintf("%v%s", v, separator))
		}

	}
	return buf.String()
}

func ParseTemplate(templ []byte, data []byte) ([]byte, error) {
	return ParseTemplateAddFunc(templ, data, template.FuncMap{})
}

func ParseTemplateAddFunc(templ []byte, data []byte, additinalFunctions template.FuncMap) ([]byte, error) {
	m := map[string]interface{}{}

	if len(templ) == 0 {
		return []byte{}, fmt.Errorf("no template ybtes given")
	}

	if len(data) == 0 {
		return []byte{}, fmt.Errorf("no data bytes given")
	}

	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return []byte{}, err
	}

	funcMap := template.FuncMap{
		"hasNext":   hasNext,
		"isLast":    isLast,
		"arrayjoin": arrayJoin,
		"join":      strings.Join,
	}

	for k, v := range additinalFunctions {
		funcMap[k] = v
	}

	t, err := template.New("").Funcs(funcMap).Parse(string(templ))
	if err != nil {
		return []byte{}, err
	}

	var outBuf bytes.Buffer

	err = t.Execute(&outBuf, m)
	if err != nil {
		return []byte{}, err
	}

	return outBuf.Bytes(), nil
}
