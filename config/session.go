package config

import (
	"github.com/gorilla/sessions"
)

const SESSION_ID = "user_logged_token"

var Store = sessions.NewCookieStore([]byte("ajksgk0934712qwfqqr"))
