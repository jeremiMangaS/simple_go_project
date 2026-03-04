package read_g

import (
	//
	"fmt"
	"io/fs"
	"path/filepath"
)

//
func Check_path(path_args string) { 
	if path_args == "." {
		//
		fmt.Print("From current folder")
		return
	} 
	fr_path := path_args
	read_fr(fr_path)
}

func read_fr(fr_path string) {
	ext_html := ".html"
	// ext-md := ".md"
	//
	// WalkDir has function as it parameters
	// path string : file in process
	// d fs.DirEntry : file detail
	// err error : error message
	// returning error in order to check it ssuccess or failed
	filepath.WalkDir(fr_path, func(path string, d fs.DirEntry, err error) error {
		//
		if err != nil {
			return err
		}
		// IsDir() : checking is the current thing is folder or not, returning true  or false
		// Ext() : taking file extension (exm : .txt, .gp, .md, etc)
		if !d.IsDir() && filepath.Ext(path) == ext_html {
			info, errr := d.Info()
			if errr != nil {
				return errr
			}
			f_name := info.Name()
			var html_obj *HTML_f
			html_obj.F_Path = path
			html_obj.F_name = f_name
			html_obj.F_Extension = ext_html
			F_read_lists.HTML_files[html_obj]
		}

		return err
	})
}

