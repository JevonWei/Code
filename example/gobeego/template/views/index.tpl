<!DOCTYPE html>
<html>
  <head>
    <title>welcome</title>
  </head>
  <body>

	{{template "block"}}
	{{template "header"}}
	{{template "blocks/block.tpl"}}

	<h2>{{.person.Welcome "hello" }}</h2>
	<h2>{{call .person.Introduce }}</h2>
	<h2>{{call .person.Introduce }}</h2>
	<p> This is struct: {{ .person }}</p>
  </body>
</html>
