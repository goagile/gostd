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

func (g *getall) From(data map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	for _, getter := range g.getters {
		result[getter.Key()] = getter.From(data)
	}
	return result
}
