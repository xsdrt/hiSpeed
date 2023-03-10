package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

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
	}

	//Must cookies be secure?
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	//Create the session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	//Which session store to use...
	switch strings.ToLower(c.SessionType) {
	case "redis":

	case "mysql", "mariadb":

	case "postgres", "postgresql":

	case "mssql":

	default:
		// cookie

	}

	return session
}
