package handlers

// ResponseMsg is a struct used to wrap responses indicating success
type ResponseMsg struct {
	Message string `json:"message"`
	Result  interface{}
	Status  int
}
