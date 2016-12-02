package CSVManager

import "strings"

// CSVLine ...
type CSVLine struct {
	Values []string
}

// New ...
func (l *CSVLine) New(b []byte) CSVLine {
	return CSVLine{Values: strings.Split(string(b), ",")}
}

// GetValue ...
func (l *CSVLine) GetValue(byN int) (string, bool) {
	if len(l.Values) >= byN {
		return l.Values[byN], true
	}
	return "", false
}
