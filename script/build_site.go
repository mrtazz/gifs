package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	gifs         = make(map[string][]string)
	htmlTemplate = `

<!DOCTYPE html>
<html lang="en-us">

<head>
  <meta charset="UTF-8">
  <title>gifs.unwiredcouch.com</title>
  <link href='http://fonts.googleapis.com/css?family=Open+Sans:400,300,800' rel='stylesheet' type='text/css'>
  <link rel="stylesheet" href="css/main.css">
  <meta name="viewport" content="initial-scale=1">
</head>


<body>

  <header class="page-header">
    <div class="img-preview">
      <img src="excited/aha.gif" alt="" class="banner-img">
      <img src="" id="preview" class="gif-preview">
    </div>


    <h1>My Gif Collection</h1>

  </header>

  <div class="gif-listing">


  <ul>
  {{range $k, $v := .}}

  <li>
  {{$k}}
  <ul>
  {{range $v}}
  <li>
  <a href="{{ $k }}/{{ . }}" title="{{ $k }}/{{ . }}" class="gif">{{ . }}</a>
  </li>
  <li>
  {{end}}
  </ul>
  </li>
  {{end}}
  </ul>

  </div>

  <script type="text/javascript">
    var preview = document.getElementById("preview");

    function hoverListener(evt) {
      preview.src = evt.target.href;
    };

    var links = document.getElementsByClassName("gif");
    for(var i=0; i<links.length; i++) {
      links[i].addEventListener("mouseover", hoverListener);
    }
  </script>

</body>
</html>
`
)

func main() {
	err := filepath.Walk(".",
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(fpath, ".gif") {
				gifs[path.Dir(fpath)] = append(gifs[path.Dir(fpath)], path.Base(fpath))
			}
			return nil
		})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	tmpl, err := template.New("index.html").Parse(htmlTemplate)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	f, err := os.Create("index.html")
	if err != nil {
		log.Println("create file: ", err)
		os.Exit(1)
	}
	defer f.Close()
	err = tmpl.Execute(f, gifs)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
