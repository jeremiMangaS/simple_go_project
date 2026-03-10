package process_g

import (
	//
	// "fmt"

	"fmt"

	struct_read "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/struct_l"
	// "github.com/jeremiMangaS/simple_go_project/goat_m/parser_g"
	"github.com/jeremiMangaS/simple_go_project/goat_m/parser_g/html_g"
	// "github.com/jeremiMangaS/simple_go_project/goat_m/parser_g/md_g"
)

func Prc_parsing(html_obj_l []struct_read.HTML_f, md_obj_l []struct_read.MD_f) {
	//
	// inserting goa tag data into the html object

	for i := range html_obj_l {
		html_obj_l[i].Ghi_map = parser_html.Goat_html_information_tag(html_obj_l[i].F_content)
	}
	
	// inserting
	// for i := range md_obj_l {
	// 	md_obj_l[i].Gmi_map = parser_md.Goat_md_information_tag(md_obj_l[i].F_content)
	// }

	fmt.Println(html_obj_l)
	// fmt.Println(md_obj_l)


	// // MD 
	// // for i := range md_obj_l {
	// // 	fmt.Println(md_obj_l[i].Gmi_tag)
	// // }

	// // matching process
	// parser_g.Matching_prc(html_obj_l, md_obj_l)
}
