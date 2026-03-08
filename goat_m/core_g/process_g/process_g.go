package process_g

import (
	//
	// "fmt"

	struct_read "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/struct_l"
	"github.com/jeremiMangaS/simple_go_project/goat_m/parser_g/html_g"
)

func Prc_parsing(html_obj_l []struct_read.HTML_f, md_obj_l []struct_read.MD_f) {
	//
	// inserting goa tag data into the html object
	for i := range html_obj_l {
		res := parser_html.G_tag(html_obj_l[i].F_content)
		html_obj_l[i].Goat_tag = res
	}
}
