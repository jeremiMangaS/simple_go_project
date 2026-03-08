package struct_l

type Base_structure_Information struct {
	//
	F_Path string			// File full path [v]
	F_name string			// File name [v]
	F_Extension string		// File extension (.html, .md) [v]
	Fr_Parent string		// Parent folder name 
	Is_Root bool 			// is it at root? 
	F_content string		// file content
}

type MD_f struct {
	// include
	Base_structure_Information
}

type HTML_f struct {
	// include
	Base_structure_Information
	Goat_tag string 		// goat gtml tag information
}

type F_read_lists struct {
	MD_files []MD_f
	HTML_files []HTML_f
}