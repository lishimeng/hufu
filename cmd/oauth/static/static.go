package static

import (
	"embed"
	"net/http"
)

//go:embed *
var Static embed.FS

var FS = http.FS(Static)
