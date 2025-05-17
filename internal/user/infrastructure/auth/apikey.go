package auth

type Apikey struct {
	UserId uint32
}

type TokenPair struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}
