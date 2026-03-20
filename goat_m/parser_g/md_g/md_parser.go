package parser_md

import (
	// "regexp"
	"fmt"
	"strings"
)

//

func Goat_md_information_tag(f_content string) map[string]string {
	//
	const gmi_start = "---"
	const gmi_end = "---"

	start_tag_check := strings.Index(f_content, gmi_start)
	// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	if start_tag_check == -1 {
		return nil
	}

	start_tag_check += len(gmi_end)

	end_tag_check := strings.Index(f_content[start_tag_check:], gmi_end)
	if end_tag_check == -1 {
		return nil
	}

	f_content = f_content[start_tag_check : start_tag_check + end_tag_check]

	// re := regexp.MustCompile(`\s+`)
	// f_content = re.ReplaceAllString(f_content, "")
	f_content = strings.ReplaceAll(f_content, " ", "")
	f_content = strings.ReplaceAll(f_content, "\n", "")
	f_content = strings.ReplaceAll(f_content, "\r", "")
	f_content = strings.ReplaceAll(f_content, "\t", "")

	gmi_map := make(map[string]string)
	// Inserting the gmi into thhe map
	// start := 0
	// for i := 0; i < len(f_content); i++ {
	// 	if f_content[i] == ',' {
	// 		text_parsing := f_content[start : i]
	// 		start = i + 1
	// 		new_text := strings.SplitN(text_parsing, ":", 2)
	// 		gmi_map[new_text[0]] = new_text[1]
	// 	}
	// }
	split_content := strings.Split(f_content, ",")

	// s : split
	for _, s := range split_content {
		if s == "" {
			continue
		}
		split_colon := strings.SplitN(s, ":", 2)
		if len(split_colon) == 2 {
			key := strings.TrimSpace(split_colon[0])
			value := strings.TrimSpace(split_colon[1])
			gmi_map[key] = value
		}
	}

	fmt.Println(gmi_map)

	return gmi_map

	//
}