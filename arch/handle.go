package arch

// Response is a final response to a request
type Response struct {
	body    []byte
	status  int
	headers map[string]string
}
