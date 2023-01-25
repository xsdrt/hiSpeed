package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder() //Need a response writer , but not a real response writer; using NewRecorder...

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page.", err)
	}
	//Above test passed...

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
	//With this last test the page rendering func has been completely covered in the render.go file and passes the testing...
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

//Some test cmd for the render file : run tests in render (go test .), run test verbose (go test -v .)
// Some test cmd if we go up one to hiSpeed : cd .. , then (make cover) opens browser to show what has
// been covered in the test(green), and what has not (red).
