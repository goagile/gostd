package interpret

//
// NothingToGet
//
func NothingToGet() *nothingtoget {
	return &nothingtoget{}
}

type nothingtoget struct{}

func (n *nothingtoget) Key() string {
	return ""
}

func (n *nothingtoget) From(data map[string]interface{}) interface{} {
	return nil
}
