package interpret

import "log"

//
// GetEach
//
func GetEach(key string, getters ...KeyGetter) *geteach {
	if len(getters) == 0 {
		return &geteach{key, []KeyGetter{NothingToGet()}}
	}
	return &geteach{key, getters}
}

type geteach struct {
	key     string
	getters []KeyGetter
}

func (e *geteach) From(inputresult map[string]interface{}, data map[string]interface{}) interface{} {
	result := []interface{}{}

	switch items := data[e.key].(type) {

	case []interface{}:
		for _, i := range items {
			itemresult := map[string]interface{}{}
			elem, ok := i.(map[string]interface{})
			if !ok {
				log.Println("ELEM IS NO MAP")
				continue
			}
			for _, getter := range e.getters {
				itemresult[getter.Key()] = getter.From(inputresult, elem)
			}
			result = append(result, itemresult)
		}

	default:
		log.Println("DATA NOT ITERABLE")
	}
	return result
}
