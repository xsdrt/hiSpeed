package session

import (
	"strconv"
	"strings"

	"github.com/alexedwards/scs/v2"
)

type Session struct { //This will be exportable..
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	//How long should sessions last?
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	//Should cookies persist?
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	//Must cookies be secure?
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}
}
