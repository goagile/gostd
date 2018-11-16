package interpret

import (
	"errors"
	"log"
)

//
// Get
//
func Get(key string, values ...interface{}) *get {
	if len(values) == 0 {
		return &get{key, NothingToGet()}
	}
	if len(values) > 1 {
		getters := []KeyGetter{}
		for _, v := range values {
			g := v.(KeyGetter)
			getters = append(getters, g)
		}
		return &get{key, GetAll(getters...)}
	}
	return &get{key, values[0]}
}

type get struct {
	key   string
	value interface{}
}

func (g *get) Key() string {
	return g.key
}

func (g *get) From(inputresult map[string]interface{}, data map[string]interface{}) interface{} {
	switch v := data[g.key].(type) {

	case map[string]interface{}:
		log.Printf("%T\n", g.value)
		switch x := g.value.(type) {
		case KeyGetter:
			return x.From(inputresult, v)
		case KeySetter:
			inner := map[string]interface{}{}
			x.To(inner, v)
			inputresult[x.Key()] = inner
		default:
			panic(errors.New("тип внутри Getter не поддерживается"))
			return nil
		}

	default:
		return data[g.key]
	}
	return nil
}
