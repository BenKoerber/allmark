// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

import (
	"fmt"
)

var masterTemplate = fmt.Sprintf(`<!DOCTYPE HTML>
<html lang="{{.LanguageTag}}">
<meta charset="utf-8">
<head>
	<title>{{.Title}}</title>

	<link rel="schema.DC" href="http://purl.org/dc/terms/">
	<meta name="DC.date" content="{{.Date}}">

	<link rel="stylesheet" type="text/css" href="/theme/screen.css">
</head>
<body>

<article>
%s
</article>

<script src="//ajax.googleapis.com/ajax/libs/jquery/2.0.0/jquery.min.js"></script>
<script type="text/Javascript">
	$(function() { 
		var socket;

		if (window["WebSocket"]) {
			routeParameter = "route=" + document.location.pathname;
			host = document.location.host;
			webSocketHandler = "/ws";
			websocketUrl = "ws://" + host + webSocketHandler + "?" + routeParameter;

			socket = new WebSocket(websocketUrl);

			socket.onclose = function(evt) {
				console.log("Connection closed.");
			};

			socket.onmessage = function(evt) {
				if (typeof(evt) !== 'object' || typeof(evt.data) !== 'string') {
					console.log("Invalid data from server.");
					return;
				}

				message = JSON.parse(evt.data);
				console.log(message);
			};

		} else {
			console.log("Your browser does not support WebSockets.");
		}
	});
</script>

</body>
</html>`, ChildTemplatePlaceholder)

const repositoryTemplate = `
<header>
<h1>
{{.Title}}
</h1>
</header>

<section>
{{.Description}}
</section>

<section>
{{.Content}}
</section>

<section>
<ul>
{{range .Childs}}
<li>
	<a href="{{.Route}}">{{.Title}}</a>
	<p>{{.Description}}</p>
</li>
{{end}}
</ul>
</section>
`

const collectionTemplate = `
<header>
<h1>
{{.Title}}
</h1>
</header>

<section>
{{.Description}}
</section>

<section>
{{.Content}}
</section>

<section class="collection">
<h1>Documents</h2>
<ol>
{{range .Childs}}
<li>
	<h2><a href="{{.Route}}" title="{{.Description}}">{{.Title}}</a></h2>
	<p>{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
`

const documentTemplate = `
<header>
<h1>
{{.Title}}
</h1>
</header>

<section>
{{.Description}}
</section>

<section>
{{.Content}}
</section>
`

const messageTemplate = `
<section>
{{.Content}}
</section>

<section>
{{.Description}}
</section>
`

const errorTemplate = `
<header>
<h1>
{{.Title}}
</h1>
</header>

<section>
{{.Description}}
</section>

<section>
{{.Content}}
</section>
`