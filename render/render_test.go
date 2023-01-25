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
	errorExpected bool   // Is there a error expected, when test run should eor be nil or !nil ...
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

	/* //Above test passed...		removed/added/modified some to above; Commented all the rest out; original testing, see if above works...

	err = testRenderer.Page(w, r, "no-file", nil, nil) //Test method calling for a non-existent file and indicating it should be == nil;
	if err == nil {                                    // this should fail, which indicates a pass for testing for an err for non existent file to be parsed
		t.Error("Error rendering non-existent template.", err)
	}

	testRenderer.Renderer = "jet" //  Tested before adding the page  to testdata/views to make sure it fails; it did....Great...
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page.", err)
	}
	//Test after add the page passed...OK   //* Just an FYI: use go test -cover . to check how much a folder has been tested;
	//but this will not tell you what you have tested in the statements themselves.  So create test coverage Makefile in the root of the project to do this...

	err = testRenderer.Page(w, r, "no-file", nil, nil) //Test method calling for a non-existent jet file and indicating it should be == nil;
	if err == nil {                                    // this should work, which indicates a test pass for testing for an err for non-existent jet file to be parsed...
		t.Error("Error rendering non-existent jet template.", err)
	}

	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "home", nil, nil) //Test method calling for a invalid render indicating it should be == nil;
	if err == nil {                                 // this should fail, which indicates a test pass for testing for an invalid render
		t.Error("No error returned while rendering with invalid renderer specified.", err) //(in this case non specified so falls to default in render.go that returns an error)...
	}
	//With this last test the page rendering func has been completely covered in the render.go file and passes the testing... */
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
