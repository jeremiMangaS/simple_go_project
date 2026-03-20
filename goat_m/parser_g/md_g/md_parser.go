package parser_md

import (
	// "regexp"
	"fmt"
	"strings"
)

//

func Goat_md_information_tag(f_content string) map[string]string {
	//
	// const gmi_start = "---"
	// const gmi_end = "---"

	const gmi_separation = "---"

	parts := strings.SplitN(f_content, gmi_separation, 3)

	if len(parts) < 3 {
		return nil
	}

	gmi_map := make(map[string]string)

	config_line := strings.Split(parts[1], "\n")
	for _, line := range config_line {
		line = strings.TrimSpace(line)
		if strings.Contains(line, ":") {
			kv_part := strings.SplitN(line, ":", 2)
			gmi_map[strings.TrimSpace(kv_part[0])] = strings.TrimSpace(kv_part[1])
		}
	}

	content_line := strings.Split(parts[2], "\n")
	for _, content_block := range content_line {
		content_block = strings.TrimSpace(content_block)
		if strings.Contains(content_block, ":") {
			kv_part := strings.SplitN(content_block, ":", 2)
			key := strings.TrimSpace(kv_part[0])
			value := strings.TrimSpace(kv_part[1])
			gmi_map[key] = value
		}
	}



	// start_tag_check := strings.Index(f_content, gmi_start)
	// // Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	// if start_tag_check == -1 {
	// 	return nil
	// }

	// start_tag_check += len(gmi_end)

	// end_tag_check := strings.Index(f_content[start_tag_check:], gmi_end)
	// if end_tag_check == -1 {
	// 	return nil
	// }

	// f_content = f_content[start_tag_check : start_tag_check + end_tag_check]

	// gmi_map := make(map[string]string)

	// lines := strings.Split(f_content, "\n")
	// for _, line := range lines {
	// 	line = strings.TrimSpace(line)
	// 	if line == "" || !strings.Contains(line, ":") {
	// 		continue
	// 	}

	// 	parts := strings.SplitN(line, ":", 2)
	// 	key := strings.TrimSpace(parts[0])
	// 	value := strings.TrimSpace(parts[1])
	// 	value = strings.TrimSuffix(value, ",")

	// 	gmi_map[key] = value
	// }

	fmt.Println(gmi_map)

	return gmi_map

	//
}