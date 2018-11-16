package interpret

//
// GetAll
//
func GetAll(getters ...KeyGetter) *getall {
	return &getall{getters}
}

type getall struct {
	getters []KeyGetter
}

func (g *getall) Key() string {
	return ""
}

func (g *getall) From(inputresult map[string]interface{}, data map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	for _, getter := range g.getters {
		result[getter.Key()] = getter.From(inputresult, data)
	}
	return result
}
