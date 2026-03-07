package read_g

import (
	//
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"
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

func central_analys_g(fr_path string) {
	//
	var waitg_g sync.WaitGroup
	waitg_g.Add(4)

	// channels
	path_chan := make(chan string)
	html_result_chan := make(chan HTML_f)
	md_result_chan := make(chan MD_f)
	// defer close(html_result_chan)
	// defer close(md_result_chan)

	// extensions
	ext_html := ".html"
	ext_md := ".md"

	var html_obj_l []HTML_f
	var md_obj_l []MD_f
	
	// structure explorer
	go func() {
		// defer fmt.Println("TESTING")
		filepath.WalkDir(fr_path, func(path string, d fs.DirEntry, err error) error {
			//
			if !d.IsDir() && (filepath.Ext(path) == ext_html || filepath.Ext(path) == ext_md) {
				path_chan <- path
				// fmt.Println(path)
			}
			return err
		})
		// closing the channel
		close(path_chan)
		defer waitg_g.Done()
	}()
	
	// file object manager
	go func() {
		for path := range path_chan {
			ext_f := filepath.Ext(path)
			if ext_f == ext_html {
				new_obj_f := manage_html_obj_f(path, fr_path)
				html_result_chan <- new_obj_f
			}
			if ext_f == ext_md {
				new_obj_f := manage_md_obj_f(path, fr_path)
				md_result_chan <- new_obj_f
			}
		}
		defer close(html_result_chan)
		defer close(md_result_chan)
		defer waitg_g.Done()
	}()

	// grouping manager
	go func() {
		for obj := range html_result_chan {
			html_obj_l = append(html_obj_l, obj)
		}
		defer waitg_g.Done()
	}()
	go func() {
		for obj := range md_result_chan {
			md_obj_l = append(md_obj_l, obj)
		}
		defer waitg_g.Done()
	}()
	
	waitg_g.Wait()

	defer fmt.Println(html_obj_l)
	defer fmt.Println("==============================")
	defer fmt.Println(html_obj_l)

	//
}

func manage_html_obj_f(path string, fr_path string) HTML_f {
	//
	// full path, file name, file extension, folder parent, is root
	f_name := filepath.Base(path)
	f_ext := filepath.Ext(path)
	fr_parent := filepath.Base(filepath.Dir(path))
	is_root := (fr_parent == fr_path)

	return HTML_f{
		//
		Base_structure_Information: Base_structure_Information{
			F_Path: path,
			F_name: f_name,
			F_Extension: f_ext,
			Fr_Parent: fr_parent,
			Is_Root: is_root,
		},
	}
}

func manage_md_obj_f(path string, fr_path string) MD_f {
	//
	// full path, file name, file extension, folder parent, is root
	f_name := filepath.Base(path)
	f_ext := filepath.Ext(path)
	fr_parent := filepath.Base(filepath.Dir(path))
	is_root := (fr_parent == fr_path)

	return MD_f{
		//
		Base_structure_Information: Base_structure_Information{
			F_Path: path,
			F_name: f_name,
			F_Extension: f_ext,
			Fr_Parent: fr_parent,
			Is_Root: is_root,
		},
	}
}