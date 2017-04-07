package help

import (
	"net/http"
	"github.com/gobly/ui"
)

var singleLayout = ui.LoadSingle("html/main.html")

func AboutHandler(out http.ResponseWriter, in *http.Request) {
	singleLayout.Execute(out, nil)
}
