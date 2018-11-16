package interpret

//
// Getter
//
type Getter interface {
	From(data map[string]interface{}) interface{}
}
