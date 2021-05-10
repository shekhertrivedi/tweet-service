package model

// Tweet entity
type Tweet struct {
	ID       string   `json:"id"`
	Message  string   `json:"message"`
	HashTags []string `json:"hashTags"`
}

// ErrorResponse error reponse
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
