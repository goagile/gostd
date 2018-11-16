package interpret

//
// Get
//
func Get(key string, getters ...KeyGetter) *get {
	if len(getters) == 0 {
		return &get{key, NothingToGet()}
	}
	if len(getters) > 1 {
		return &get{key, GetAll(getters...)}
	}
	return &get{key, getters[0]}
}

type get struct {
	key    string
	getter Getter
}

func (g *get) Key() string {
	return g.key
}

func (g *get) From(data map[string]interface{}) interface{} {
	switch v := data[g.key].(type) {

	case map[string]interface{}:
		return g.getter.From(v)

	default:
		return data[g.key]
	}
}
