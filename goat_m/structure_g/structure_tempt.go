package pkg_structure

const simple_basic_HTML = `<html>
<head><title>{{ .title }}</title></head>
<goat 
    $key_status = true
/>
<body>
    <h1>{{ .title1 }}</h1>
    {{ .body1 }}
</body>
</html>`

const simple_basic_MD = `---
$key_file: index.html
title: Welcome to GOat
---
t-ct_stitle1: Introduction
b-ct_body1: GOat is a modular web framework created as a competitor to Hugo. It supports various rendering types and can be used for both simple and complex projects.

t-ct_stitle2: Modular and Flexible
b-ct_body2: Each GOat project is tailored to its complexity level. For simple projects, you can use the basic structure. For more advanced projects, additional modules can be installed as needed.
`