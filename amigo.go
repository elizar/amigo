package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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
			margin  : 0;
			padding : 0;
			border  : none;
			outline : none;
		}

		body {
			font-family: Arial, Helvetica, sans-serif;
			font-size: 18px;
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
			height: 300px;
			text-align: center;
			background: #fff url('https://golang.org/doc/gopher/appenginegopher.jpg') bottom center no-repeat;
			background-size: 50%;
			padding: 50px;
			border: 1px solid rgba(0,0,0,0.2);
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
	</style>
</head>
<body>
	<main>
		<h1>Gopher Wisdom</h1>
		<br>
		<br>
		<blockquote>
			A deployed MVP is worth two prototypes.
		</blockquote>
		<br>
	</main>
	<footer>
		<p>
			<small>&copy; {{.Year}} Elizar Pepino. All rights reserved.</small>
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

		t, _ := template.New("amigo").Parse(tmpl)
		t.ExecuteTemplate(w, "index", struct {
			Title    string
			Hostname string
			Year     int
		}{"Amigo", host, time.Now().Year()})

		elapse := time.Since(now) / time.Millisecond
		log.Printf("%s %s %dms\n", r.Method, r.RequestURI, elapse)
	})

	log.Println(fmt.Sprintf("Server running on %s:%s", host, port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
