package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Implement a table test below
var pageData = []struct {
	name          string //Name of the test (meaningful so know what test failed...
	renderer      string // Will be go or jet or empty string or whatever engine used...
	template      string // What template used when test run...
	errorExpected bool   // Is there a error expected, when test run should error be nil or !nil ...
	errorMessage  string // What error message to display...
}{
	{"go_page", "go", "home", false, "error rendering go template"}, //go page tests populated...
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"}, //jet page tests populated...
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existent jet template, when one is expected"},
	{"invalid_render_engine", "foo", "home", true, "no error rendering non-existent template engine"}, //Default fall through if no valid engine...
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder() //Need a response writer , but not a real response writer; using NewRecorder...

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}

}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page.", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page.", err)
	}

}

//Some test cmd for the render file (cd render) : run tests in render (go test .), run test verbose (go test -v .)
// Some test cmd if we go up one to hiSpeed : cd .. , then (make cover) opens browser to show what has
// been covered in the test(green), and what has not (red).
