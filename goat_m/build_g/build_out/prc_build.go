package prc_build

import (
	//
	"strings"

	struct_read "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/struct_l"
)

func Build_Page(md_obj *struct_read.MD_f, html_obj *struct_read.HTML_f) string {
	//
	out_temp := html_obj.F_content

	for k, v := range md_obj.Gmi_map {
		//
		target_key := "{ ." + k + " }"
		out_temp = strings.ReplaceAll(out_temp, target_key, v)
	}

	// out_temp = strings.ReplaceAll(out_temp, "<goat>", md_obj.F_content)
	// return out_temp
	return out_temp
}