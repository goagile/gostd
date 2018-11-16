package interpret

//
// Getter
//
type Getter interface {
	From(result map[string]interface{}, data map[string]interface{}) interface{}
}
