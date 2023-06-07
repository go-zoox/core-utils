package strings

import (
	"bytes"
	"fmt"
	"text/template"
)

// Template formats a template string with go template syntax.
func Template(rawTemplate string, data map[string]any) (builtText string, err error) {
	// 定义模板
	t := template.New("go-zoox.core-utils.strings.template")

	// 解析模板
	t, err = t.Parse(rawTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	// 使用 buffer 来格式化字符串，再把 buffer 写入 result
	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buffer.String(), nil
}
