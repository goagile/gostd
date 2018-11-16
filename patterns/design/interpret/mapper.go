package interpret

//
// Mapper
//
func Mapper(setters ...Setter) *mapper {
	return &mapper{setters}
}

type mapper struct {
	setters []Setter
}

func (m *mapper) Map(data map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, set := range m.setters {
		set.To(result, data)
	}
	return result
}
