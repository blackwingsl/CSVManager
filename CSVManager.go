package CSVManager

// CSVManager ...
type CSVManager struct {
	List CSVList
}

// Load target csv file
func (m *CSVManager) Load(pathAndFilename string) error {
	e := m.List.Update(pathAndFilename)
	return e
}
