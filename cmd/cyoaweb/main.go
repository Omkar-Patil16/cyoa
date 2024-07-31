package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/Omkar-Patil16/cyoa.git"
)

func main() {
	// creating a flag to pass the filename to be parsed
	port := flag.Int("port", 3000, "application running port no")
	fileName := flag.String("file", "gopher.json", "file with adventure story")
	flag.Parse()

	fmt.Printf("Name Of the file selected %s\n", *fileName)

	// open the file

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonDecode(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)
	fmt.Printf("Server is running on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func pathFn(r *http.Request) string {
	path := r.URL.Path
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
