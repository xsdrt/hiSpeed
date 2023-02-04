package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {

	h := &Session{ //Values used for a receiver (the session init uses a receiver with value(s)...)
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "hiSpeed",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager

	ses := h.InitSession() //Make sure a return value from this call to InitSession so cerate some variables to start with...

	var sessKind reflect.Kind //From the reflect package
	var sessType reflect.Type //From the reflect package

	rv := reflect.ValueOf(ses) //Give it the ses vale getting back from InitSession and use for initial testing...

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For Loop:", rv.Kind(), rv.Type(), rv) //Print from the for loop to see return values...
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}
	//Then check/test to make sure the values are valid...
	if !rv.IsValid() {
		t.Error("Invalid Kind or Type; Kind:", rv.Kind(), "Type:", rv.Type()) //If not valid print out info to see...
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned testing cookie session. Expected", reflect.ValueOf(sm).Kind(), "and got", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned testing cookie session. Expected", reflect.ValueOf(sm).Type(), "and got", sessType)
	}

	//Remember to test: cd into session first :)....
}
