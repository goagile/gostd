package interpret

//
// SetAll
//
func SetAll(setters ...Setter) *setall {
	return &setall{setters}
}

type setall struct {
	setters []Setter
}

func (e *setall) To(result map[string]interface{}, data map[string]interface{}) {
	for _, s := range e.setters {
		s.To(result, data)
	}
}
