package CSVManager

import (
	"bytes"
	"strings"
)

// CSVContent ...
type CSVContent struct {
	Name  string
	Key   []string
	Lines []CSVLine
}

// New create new CSVContent
func (c *CSVContent) New(name string, file []byte) CSVContent {
	content := CSVContent{}
	content.Name = name
	content.Key = []string{}
	content.Lines = []CSVLine{}
	lines := bytes.Split(file, []byte{'\r', '\n'})
	for _, v := range lines {
		if v[0] == '#' { // 注释行
			if v[1] == '!' { // 标题行
				content.Key = strings.Split(string(v[2:]), ",")
			}
			continue
		}
		// 正文截取
		line := new(CSVLine)
		content.Lines = append(content.Lines, line.New(v))
	}
	return content
}

// GetLineByKV ..
func (c *CSVContent) GetLineByKV(key, value, needKey string) (string, bool) {
	n, result := c.GetN(key)
	if !result {
		return "", false
	}

	for _, v := range c.Lines {
		if v.Values[n] == value {
			need, needResult := c.GetN(needKey)
			if needResult {
				return v.Values[need], true
			}
			break
		}
	}

	return "", false
}

// GetContent ...
func (c *CSVContent) GetContent(key, value string) (CSVLine, bool) {
	n, result := c.GetN(key)
	if !result {
		return CSVLine{}, false
	}

	for _, v := range c.Lines {
		if v.Values[n] == value {
			return v, true
		}
	}

	return CSVLine{}, false
}

// GetN ...
func (c *CSVContent) GetN(key string) (int, bool) {
	for k, v := range c.Key {
		if v == key {
			return k, true
		}
	}
	return 0, false
}
