package structure

const Simple_HTML = `<html>
<head><title>{{ .title }}</title></head>
<goat 
    $key_status = true
/>
<body>
    <h1{{ .title }}</h1>
    {{ .content }}
</body>
</html>`

const Simple_MD = `---
$key_file : index.html
title : Welcome to GOat
---
content : Try to make a simple web with simple way`
