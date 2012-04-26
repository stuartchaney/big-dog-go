package main

import (
	"http"
	"io"
	"log"
	"os"
)

var port = func() string {
	tmpport := os.Getenv("PORT")
	if tmpport == "" {
		tmpport = "5000"
	}

	return tmpport
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":" + port(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}

const html = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

	<title>Stuart Go Test</title>
	<link href='http://fonts.googleapis.com/css?family=Lato:regular,bold' rel='stylesheet' type='text/css'>
	<style type="text/css">
		body {
  		margin-top: 1.0em;
		  font-family: 'Lato', arial, serif;
  		color: #fff;
    }
    #container {
      margin: 0 auto;
      width: 700px;
    }
		h1 { font-size: 3.8em; margin-bottom: 3px; }
		h1 .small { font-size: 0.4em; }
		h1 a { text-decoration: none }
		h2 { font-size: 1.5em;}
    h3 { text-align: center;}
    a { color:#C83535; }
    .description { font-size: 1.2em; margin-bottom: 30px; margin-top: 30px; font-style: italic;}
    .download { float: right; }
		pre { background: #000; color: #fff; padding: 15px;}
    hr { border: 0; width: 80%; border-bottom: 1px solid #aaa}
    .footer { text-align:center; padding-top:30px; font-style: italic; }
	</style>
	
</head>

<body>
  <div id="container">

    <div class="description">
      The Big Dog on GO..... shit is about to get messy.....
    </div>
    <br/>
  
    <img src="http://f.cl.ly/items/2T0G2C0T2D3W3K033R04/Screen%20Shot%202012-03-12%20at%201.24.03%20PM.png" />
	
  </div>

  
</body>
</html>
`

