package interpret

//
// Setter
//
type Setter interface {
	To(result, data map[string]interface{})
}

//
// Getter
//
type Getter interface {
	From(data map[string]interface{}) interface{}
}

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
		return &set{key, Fields(setters...)}
	}
	return &set{key, values[0]}
}

type set struct {
	key   string
	value interface{}
}

func (s *set) To(result map[string]interface{}, data map[string]interface{}) {
	switch v := s.value.(type) {

	case Setter:
		inner := map[string]interface{}{}
		v.To(inner, data)
		result[s.key] = inner

	case Getter:
		result[s.key] = v.From(data)

	default:
		result[s.key] = v
	}
}

//
// Get
//
func Get(key string, getters ...Getter) *get {
	if len(getters) == 0 {
		return &get{key, NothingToGet()}
	}
	return &get{key, getters[0]}
}

type get struct {
	key    string
	getter Getter
}

func (g *get) From(data map[string]interface{}) interface{} {
	switch v := data[g.key].(type) {
	case map[string]interface{}:
		return g.getter.From(v)
	default:
		return data[g.key]
	}
}

//
// NothingToGet
//
func NothingToGet() *nothingtoget {
	return &nothingtoget{}
}

type nothingtoget struct{}

func (n *nothingtoget) From(data map[string]interface{}) interface{} {
	return nil
}

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

//
// Fields
//
func Fields(setters ...Setter) *fields {
	return &fields{setters}
}

type fields struct {
	setters []Setter
}

func (f *fields) To(result map[string]interface{}, data map[string]interface{}) {
	for _, s := range f.setters {
		s.To(result, data)
	}
}
