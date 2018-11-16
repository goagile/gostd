package interpret

//
// Set
//
func Set(key string, values ...interface{}) *set {
	if len(values) == 0 {
		return &set{key, nil}
	}
	if len(values) > 1 {
		setters := []Setter{}
		for _, v := range values {
			s := v.(Setter)
			setters = append(setters, s)
		}
		return &set{key, SetAll(setters...)}
	}
	return &set{key, values[0]}
}

type set struct {
	key   string
	value interface{}
}

func (s *set) Key() string {
	return s.key
}

func (s *set) To(result map[string]interface{}, data map[string]interface{}) {
	switch v := s.value.(type) {

	case Setter:
		inner := map[string]interface{}{}
		v.To(inner, data)
		result[s.key] = inner

	case Getter:
		result[s.key] = v.From(result, data)

	default:
		result[s.key] = v
	}
}
