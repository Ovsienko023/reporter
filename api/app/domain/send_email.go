package domain

type SendEmailRequest struct {
	Token      string   `json:"token,omitempty" swaggerignore:"true"`
	Email      string   `json:"email,omitempty"`
	Password   string   `json:"password,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Subject    string   `json:"subject,omitempty"`
	Body       string   `json:"body,omitempty"`
}
