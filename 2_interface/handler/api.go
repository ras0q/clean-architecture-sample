package handler

type API struct {
	Ping *PingHandler
	User *UserHandler
}

func NewAPI(user *UserHandler) API {
	return API{User: user}
}
