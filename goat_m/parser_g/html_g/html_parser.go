package parser_html

import (
	//
	// "fmt"
	// "os"
	"regexp"
	"strings"
)
func Goat_html_information_tag(f_content string) map[string]string {
	//
	const ghi_start = "<goat>"
	const ghi_end = "</goat>"
	
	start_tag_check := strings.Index(f_content, ghi_start)
	// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	if start_tag_check == -1 {
		return nil
	}

	start_tag_check += len(ghi_start)

	end_tag_check := strings.Index(f_content[start_tag_check:], ghi_end)
	if  end_tag_check == -1 {
		return nil
	}

	f_content = f_content[start_tag_check : start_tag_check + end_tag_check] 

	re := regexp.MustCompile(`\s+`)
	f_content = re.ReplaceAllString(f_content, "")

	// fmt.Println(f_content)

	ghi_map := make(map[string]string)
	// Inserting the ghi into the map
	start := 0
	for i := 0; i < len(f_content); i++ {
		//
		if f_content[i] == ',' {
			text_parsing := f_content[start : i]
			start = i + 1
			// fmt.Println(text_parsing)
			new_text := strings.SplitN(text_parsing, ":", 2)
			// fmt.Println("INDEX 0 : ", new_text[0])
			// fmt.Println("INDEX 1 : ", new_text[1])
			if len(new_text) == 2 {
				ghi_map[new_text[0]] = new_text[1]
			}
		}
	}
	// fmt.Println(ghi_map)
	return ghi_map
}