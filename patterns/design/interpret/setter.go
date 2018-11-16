package interpret

//
// Setter
//
type Setter interface {
	To(result, data map[string]interface{})
}
