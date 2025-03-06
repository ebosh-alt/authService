package server

// GetUserTokenResponse represents a response for getting user token
type GetUserTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type errorResponse struct {
	Message string `json:"message"`
}
