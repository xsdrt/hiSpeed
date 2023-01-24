package Render

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var testRenderer = Render{ //Need variable of render; then populate some fields...
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

//special function, anytime  test are run in this directory; if Go sees a file name setup_test.go with a func named TestMain that
//takes a param pointing to testing.M it will run this func; this func will run any test found in that dir...

func TestMain(m *testing.M) {
	os.Exit(m.Run()) //This will exit but run all the test prior to exiting...
}
