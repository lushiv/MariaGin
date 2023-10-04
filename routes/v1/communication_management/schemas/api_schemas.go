package communication_management

type ErrorResponse struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	Message string `json:"message"`
}

type SendEmailTestRequest struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	SendTo  string `json:"sendTo"`
}

type PublishRequest struct {
	Message string `json:"message"`
}
