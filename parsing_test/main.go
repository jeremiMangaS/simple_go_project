package main

import (
	"fmt"
	"strings"
)

// example of the md structure
const e_content = `---
$key_file : index.html
title : welcome
---
t_ct.title1 : Introduction
b_ct.bodyc1 : GOat is a simple framework....`

func main() {
	// split the ---
	parts := strings.Split(e_content, "---")

	// checking for the format
	if len(parts) < 3 {
		fmt.Print("Error, format not complete")
	}

	data_head := strings.TrimSpace(parts[1])
	data_body := strings.TrimSpace(parts[2])

	fmt.Println("Head : \n", data_head)
	fmt.Println("Body : \n", data_body)
	parsing_data(data_head, data_body)
}

func parsing_data(d_head string, d_body string) {
	fmt.Print("========================================\n")
	line_hf := strings.Split(d_head, "\n")
	line_bf := strings.Split(d_body, "\n")


	fmt.Println("head \n", line_hf)
	fmt.Println("body\n", line_bf)


	for _, data := range line_hf {
		// fmt.Println(data, "\n")
		data_p := strings.SplitN(data, ":", 2)
		if len(data_p) == 2 {
			key_c := strings.TrimSpace(data_p[0])
			value_c := strings.TrimSpace(data_p[1])
			fmt.Println(key_c, " : ", value_c, "\n")
		}
	}
}