package CSVManager

import (
	"bytes"
	"fmt"
	"strings"
)

// CSVContent ...
type CSVContent struct {
	Name  string
	Key   *Keys
	Lines []CSVLine
}

// New create new CSVContent
func (c *CSVContent) New(name string, file []byte) CSVContent {
	content := CSVContent{}
	content.Name = name
	key := new(Keys)
	content.Lines = []CSVLine{}
	lines := bytes.Split(file, []byte{'\r', '\n'})
	keySet := false
	for _, v := range lines {
		if len(v) <= 0 {
			fmt.Println("noLine:", name)
			continue
		}
		if v[0] == '#' { // 注释行
			if v[1] == '!' && !keySet { // 标题行
				key.Value = strings.Split(string(v[2:]), ",")
				keySet = true
			}
			continue
		}
		// 正文截取
		line := new(CSVLine)
		content.Key = key
		content.Lines = append(content.Lines, line.New(v))
	}
	return content
}

// GetLineByKV ..
func (c *CSVContent) GetLineByKV(key, value, needKey string) (string, bool) {
	n, result := c.Key.GetN(key)
	if !result {
		return "", false
	}

	for _, v := range c.Lines {
		if v.Values[n] == value {
			need, needResult := c.Key.GetN(needKey)
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
	n, result := c.Key.GetN(key)
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
