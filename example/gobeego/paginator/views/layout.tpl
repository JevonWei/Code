<!DOCTYPE html>
<html>
  <head>
    <title>welcome</title>
    <link rel="stylesheet" type="text/css" href="static/css/bootstrap.css">
  </head>
  <body>
	<table border="1" width=400>
  		<tr>
    			<th>姓名</th>
    			<th>年龄</th>
    			<th>性别</th>
  		</tr>
		{{range .Users}}
  		<tr>
    			<td>{{.Name}}</td>
    			<td>{{.Age}}</td>
    			<td>{{.Sex}}</td>
  		</tr>
		{{end}}
	</table>
	{{.LayoutContent}}
  </body>
</html>
