package handler

type API struct {
	User *UserHandler
}

func NewAPI(user *UserHandler) API {
	return API{User: user}
}
