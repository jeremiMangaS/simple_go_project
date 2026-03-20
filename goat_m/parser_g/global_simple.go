package parser_g

import (
	//
	"fmt"

	"strings"

	struct_read "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/struct_l"
	build_out "github.com/jeremiMangaS/simple_go_project/goat_m/build_g/build_out"
)

func Matching_prc(html_obj_l []struct_read.HTML_f, md_obj_l []struct_read.MD_f) {
	html_lookup := make(map[string]*struct_read.HTML_f)

	for i := range html_obj_l {
		key := strings.TrimSuffix(html_obj_l[i].F_name, ".html")
		html_lookup[key] = &html_obj_l[i]
	}

	for i := range md_obj_l {
		// 
		md := &md_obj_l[i]
		key_target := md.Gmi_map["$key_file"]

		k, v := html_lookup[key_target]
		
		if !v {
			fmt.Println("Skipped file : ", md.F_name, " <==> ", key_target)
			continue
		}

		if md.Fr_Parent == k.Fr_Parent {
			fmt.Println("Same Folder : " , md.Fr_Parent, " <===> ", k.Fr_Parent)
			//
			// fmt.Println(build_out.Build_Page(md, k), "\n")
			res := build_out.Build_Page(md, k)
			fmt.Println(res)
		}

		// if the file is in the different folder
		//
	}
}