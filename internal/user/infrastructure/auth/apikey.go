package auth

type Apikey struct {
	UserId uint32
}

type TokenPair struct {
	Token   string
	Refresh string
}
