package CSVManager

import (
	"bytes"
	"strings"
)

// CSVContent ...
type CSVContent struct {
	Name  string
	Key   *CSVKeys // 标题行
	Lines []CSVLine
}

// New create new CSVContent
func (content *CSVContent) New(name string, file []byte) CSVContent {
	// init content
	contentRead := CSVContent{}
	content.Lines = []CSVLine{}

	newKey := new(CSVKeys)
	// parse csv
	lines := bytes.Split(file, []byte{'\r', '\n'})
	// fmt.Println(lines)
	for _, v := range lines {
		if len(v) <= 0 { // 空行过滤
			continue
		}

		if v[0] == '#' { // 注释行
			if v[1] == '!' && len(newKey.Value) <= 0 { // 标题行
				newKey.Value = strings.Split(string(v[2:]), ",")
			}
			continue
		}

		contentRead.Name = name
		contentRead.Key = newKey
		line := CSVLine{}
		line.New(v)
		line.Key = newKey
		line.Values = strings.Split(string(v), ",")
		contentRead.Lines = append(contentRead.Lines, line)
	}

	return contentRead
}

// Get ...
func (content *CSVContent) Get(key, value string) (CSVLine, bool) {
	n, result := content.Key.GetN(key)
	if !result {
		return CSVLine{}, false
	}

	for _, v := range content.Lines {
		if v.Values[n] == value {
			return v, true
		}
	}

	return CSVLine{}, false
}
