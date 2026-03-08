package parser_html

import (
	// "fmt"
	"strings"
)

//
// (map[string]string, string)
// func Parse_html(f_content string) string {
// 	//
// 	value_content := g_tag(f_content)

// 	// fmt.Println(value_content)
// 	return value_content
// }

func G_tag(f_content string) string {	
	const g_start = "<goat>"
	const g_end = "</goat>"
	
	start_tag_check := strings.Index(f_content, g_start)
	// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	if start_tag_check == -1 {
		return "EMPTY[]"
	}

	start_tag_check += len(g_start)

	end_tag_check := strings.Index(f_content[start_tag_check:], g_end)
	if  end_tag_check == -1 {
		return "EMPTY[END]"
	}

	return f_content[start_tag_check : start_tag_check + end_tag_check]
}

