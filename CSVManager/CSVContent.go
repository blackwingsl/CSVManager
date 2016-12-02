package CSVManager

import (
	"bytes"
	"strings"
)

// CSVContent ...
type CSVContent struct {
	Name  string
	Key   []string
	Lines [][]string
}

// New create new CSVContent
func (c *CSVContent) New(name string, file []byte) CSVContent {
	content := CSVContent{}
	content.Name = name
	content.Key = []string{}
	content.Lines = [][]string{}
	lines := bytes.Split(file, []byte{'\r', '\n'})
	for _, v := range lines {
		if v[0] == '#' { // 注释行
			if v[1] == '!' { // 标题行
				content.Key = strings.Split(string(v[2:]), ",")
			}
			continue
		}

		line := strings.Split(string(v), ",")
		content.Lines = append(content.Lines, line)
	}
	return content
}

// GetLineByKV ..
func (c *CSVContent) GetLineByKV(key string, value string) string {

	return ""
}
