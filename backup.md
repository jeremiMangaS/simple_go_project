package read_g

import (
	//
	"fmt"
	"io/fs"
	"path/filepath"
	// "time"
)

//
func Check_path(path_args string) { 
	if path_args == "." {
		//
		fmt.Print("From current folder")
		return
	} 
	fr_path := path_args
	// read_fr(fr_path)
	central_analys_g(fr_path)
}

func read_fr(fr_path string) {
	// start := time.Now()
	// ext_html := ".html"
	// ext_md := ".md"
	//

	// var html_slices []HTML_f
	// var md_slices []MD_f

	// WalkDir has function as it parameters
	// path string : file in process
	// d fs.DirEntry : file detail
	// err error : error message
	// returning error in order to check it ssuccess or failed
	filepath.WalkDir(fr_path, func(path string, d fs.DirEntry, err error) error {
		// //
		// if err != nil {
		// 	return err
		// }
		// // IsDir() : checking is the current thing is folder or not, returning true  or false
		// // Ext() : taking file extension (exm : .txt, .gp, .md, etc)

		// // HTML file handling
		// if !d.IsDir() && filepath.Ext(path) == ext_html {
		// 	info, errr := d.Info()
		// 	if errr != nil {
		// 		return errr
		// 	}
		// 	f_name := info.Name()
		// 	parent_fr := filepath.Base(filepath.Dir(path))
		// 	// instance
		// 	var html_obj HTML_f
		// 	html_obj.F_Path = path
		// 	html_obj.F_name = f_name
		// 	html_obj.F_Extension = ext_html
		// 	html_obj.Fr_Parent = parent_fr
		// 	html_obj.Is_Root = (parent_fr == fr_path)
		// 	// append
		// 	html_slices = append(html_slices, html_obj)
		// }

		// // MD file handling
		// if !d.IsDir() && filepath.Ext(path) == ext_md {
		// 	info, errr := d.Info()
		// 	if errr != nil {
		// 		return errr
		// 	}
		// 	f_name := info.Name()
		// 	parent_fr := filepath.Base(filepath.Dir(path))
		// 	// instance
		// 	var md_obj MD_f
		// 	md_obj.F_Path = path
		// 	md_obj.F_name = f_name
		// 	md_obj.F_Extension = ext_md
		// 	md_obj.Fr_Parent = parent_fr
		// 	md_obj.Is_Root = (parent_fr == fr_path)
		// 	// append
		// 	md_slices = append(md_slices, md_obj)
		// }

		if !d.IsDir() && (filepath.Ext(path) == ".html" || filepath.Ext(path) == ".md") {
			//
			infor_manage(path)
		}

		// duration := time.Since(start)
		// fmt.Println("TIME : ", duration)
		return err
	})

	// TESTING
	// fmt.Println(html_slices)
	// fmt.Println(md_slices)

	// if html_slices != nil && md_slices != nil {
	// 	var F_read_lists_obj F_read_lists
	// 	F_read_lists_obj.HTML_files = html_slices
	// 	F_read_lists_obj.MD_files = md_slices
	// 	fmt.Println(F_read_lists_obj)
	// }
	fmt.Println("NULL")
}

func central_analys_g(fr_path string) {
	//
	go read_fr(fr_path)
}

func infor_manage(f string) {
	//
	
}

// func read_fr(fr_path string) {
// 	//
// 	filepath.WalkDir(fr_path, func(path string, d fs.DirEntry, err error) error {
// 		//
// 		result := make(chan string)
// 		if !d.IsDir() && (filepath.Ext(path) == ".html" || filepath.Ext(path) == ".md") {
// 			//
// 			result <- path
// 		}

// 		return err
// 	})
// }







=========================================================================

func f_manage(fr_path string, path string, d fs.DirEntry, ext string) {
	//
	info, err := d.Info()
	
	if err != nil {
		panic(err)
	}
	
	f_name := info.Name()
	parent_fr := filepath.Base(filepath.Dir(path))

	if ext == ".html" {
		var html_obj HTML_f
		html_obj.F_name = f_name
		html_obj.F_Extension = ext
		html_obj.F_Path = path
		html_obj.Fr_Parent = parent_fr
		html_obj.Is_Root = (parent_fr == fr_path)
	}
}