package response

type Response struct {
	Data           interface{}
	Error          string
	HttpStatusCode int
}

type httpResponse struct {
	Header headerResponse `json:"header"`
	Data   interface{}    `json:"data"`
}

type headerResponse struct {
	Error string `json:"error"`
}
