package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const tmpl = `{{define "index"}}<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{.Title}} - {{.Hostname}}</title>
	<link href="https://fonts.googleapis.com/css?family=Baloo+Paaji" rel="stylesheet">
	<style>
		* {
			margin	: 0;
			padding : 0;
			border	: none;
			outline : none;
		}

		body {
			font-family: Arial, Helvetica, sans-serif;
			font-size: 20px;
			line-height: 1.5em;
			background-color: #eaf7f5;
			color: #888;
			-webkit-font-smoothing: antialiased;
		}

		h1 {
			font-family: 'Baloo Paaji';
			font-size: 2.5em;
			line-height: 1.2em;
			color: #333;
		}

		a, a:active, a:visited {
			color: rgba(0, 0, 255, 0.5);
			padding: 1px 2px;
			text-decoration: none;
			border-bottom: 1px solid rgba(0, 0, 255, 0.5);
		}
		a:hover {
			color: rgba(0, 0, 255, 0.9);
			border-color: blue;
		}

		main {
			margin: 32px auto;
			border-radius: 10px;
			width: 400px;
			min-height: 300px;
			text-align: center;
			background: #fff url('https://golang.org/doc/gopher/appenginegopher.jpg') bottom center no-repeat;
			background-size: 50%;
			padding: 50px;
			border: 1px solid rgba(0,0,0,0.2);
			padding-bottom: 140px;
		}

		.program {
			opacity: 0.5;
		}

		hr {
			margin: 24px auto;
		}

		footer {
			text-align: center;
			opacity: 0.7;
		}
		blockquote{
			display:block;
			background: #fff;
			padding: 15px 20px 15px 50px;
			margin: 0 0 20px;
			position: relative;

			/*Font*/
			font-size: 18px;
			line-height: 1.3;
			color: #666;
			text-align: left;


			/*Box Shadow - (Optional)*/
			-moz-box-shadow: 2px 2px 15px #efefef;
			-webkit-box-shadow: 2px 2px 15px #efefef;
			box-shadow: 2px 2px 15px #efefef;
		}

		blockquote::before{
			content: "\201C"; /*Unicode for Left Double Quote*/

			/*Font*/
			font-size: 60px;
			font-weight: bold;
			color: #ccc;

			/*Positioning*/
			position: absolute;
			left: 10px;
			top:0px;
		}

		blockquote::after{
			/*Reset to make sure*/
			content: "";
		}

		blockquote a{
			text-decoration: none;
			background: #eee;
			cursor: pointer;
			padding: 0 3px;
			color: #c76c0c;
		}

		blockquote a:hover{
		 color: #666;
		}

		blockquote em{
			font-style: italic;
		}
	</style>
</head>
<body>
	<main>
		<h1>Gopher Wisdom</h1>
		<br>
		<br>
		<blockquote>
			{{ .Proverb }}
		</blockquote>
		<br>
	</main>
	<footer>
		<p>
			<small>
				A demo page written in Go. Check out source code <a href="https://github.com/elizar/amigo">here</a>
			</small>
		</p>
		<p>
			<strong>
				<small>{{.Hostname}}</small>
			</strong>
		</p>
	</footer>
</body>
</html>{{end}}`

func main() {
	host, _ := os.Hostname()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		proverb := "To code or not to code."
		resp, err := http.Get("http://proverbs-app.antjan.us/random")
		if err == nil {
			b, _ := ioutil.ReadAll(resp.Body)
			proverb = strings.Replace(string(b), "\"", "", -1)
		}

		t, _ := template.New("amigo").Parse(tmpl)
		t.ExecuteTemplate(w, "index", struct {
			Title    string
			Hostname string
			Year     int
			Proverb  string
		}{"Amigo", host, time.Now().Year(), proverb})

		elapse := time.Since(now) / time.Millisecond
		log.Printf("%s %s %dms\n", r.Method, r.RequestURI, elapse)
	})

	log.Println(fmt.Sprintf("Server running on %s:%s", host, port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
