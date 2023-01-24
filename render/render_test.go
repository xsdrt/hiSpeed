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

	testRenderer.Renderer = "jet" //  Tested before adding the page  to testdata/views to make sure it fails; it did....Great...
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page.", err)
	}
	//Test after add the page passed...OK   //* Just an FYI: use go test -cover . to check how much a folder has been tested;
	//but this will not tell you what statements themselves.  So create a make file in the root of the project to do this...
}
