package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type pageData struct {
	IsDev   bool
	Modules template.HTML
}

type manifest map[string]chunk

type chunk struct {
	File   string   `json:"file"`
	Name   string   `json:"name"`
	Source string   `json:"src"`
	CSS    []string `json:"css"`
}

var index = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<link rel="icon" type="image/svg+xml" href="/vite.svg" />
{{- if .IsDev }}
		<script type="module" src="http://localhost:5173/@vite/client"></script>
		<script type="module" src="http://localhost:5173/src/index.tsx"></script>
{{- else }}
		{{ .Modules }}
{{- end }}
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta name="color-scheme" content="light dark" />
		<link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
		<title>InvoiceGen</title>
	</head>
	<body>
		<div id="app"></div>
	</body>
</html>
`

var indexTempl = template.Must(template.New("index").Parse(index))
var prodPageData *pageData

func getManifest() manifest {
	file, err := os.ReadFile("./internal/frontend/dist/.vite/manifest.json")
	if err != nil {
		panic(err)
	}
	m := new(manifest)
	err = json.Unmarshal(file, m)
	if err != nil {
		panic(err)
	}
	return *m
}

func executeDevIndex(w http.ResponseWriter) {
	_ = indexTempl.Execute(w, pageData{IsDev: true})
}

func executeProdIndex(w http.ResponseWriter) {
	if prodPageData == nil {
		prodPageData = &pageData{
			IsDev: false,
		}

		ep := getManifest()["src/index.tsx"]
		s := strings.Builder{}
		_, _ = s.WriteString("<script type=\"module\" src=\"/")
		_, _ = s.WriteString(ep.File)
		_, _ = s.WriteString("\"></script>\n")
		for _, css := range ep.CSS {
			_, _ = s.WriteString("<link rel=\"stylesheet\" href=\"/")
			_, _ = s.WriteString(css)
			_, _ = s.WriteString("\" />\n")
		}
		prodPageData.Modules = template.HTML(s.String())
	}
	_ = indexTempl.Execute(w, prodPageData)
}
