package main

import (
	"fmt"
	"strings"
)

const e_content = `---
$key_file ; index.html
title : welcome
---
t_ct.title1 : Introduction
b_ct.bodyc1 : GOat is a simple framework....`

func main() {
	parts := strings.Split(e_content, "---")
	
}