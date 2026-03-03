package pkg_structure

const simple_basic_HTML = `<html>
<head><title>{{ .title }}</title></head>
<goat 
    $key_status = true
/>
<body>
    <h1>{{ .stitle1 }}</h1>
    {{ .body1 }}
    
    <h1>{{ .stitle2 }}</h1>
    {{ .body2 }}
</body>
</html>`

const simple_basic_MD = `---
$key_file: index.html
title: Welcome to GOat
---
stitle1: Introduction
body1: GOat is a modular web framework created as a competitor to Hugo. It supports various rendering types and can be used for both simple and complex projects.

stitle2: Modular and Flexible
body2: Each GOat project is tailored to its complexity level. For simple projects, 
you can use the basic structure. For more advanced projects, additional modules can be installed as needed.`