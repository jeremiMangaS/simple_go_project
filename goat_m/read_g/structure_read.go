package read_g

import (
	//
	"os"
	"fmt"
	"sync"
	"time"
	"io/fs"
	"path/filepath"
	struct_read "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/struct_l"
)

//
// func Check_path(path_args string) { 
// 	if path_args == "." {
// 		//
// 		fmt.Print("From current folder")
// 		return
// 	} 
// 	fr_path := path_args
// 	// read_fr(fr_path)
// 	central_analys_g(fr_path)
// }

func Central_analys_g(fr_path string) ([]struct_read.HTML_f, []struct_read.MD_f) {
	//
	if fr_path == "." {
		//
		fmt.Print("From current folder")
		return nil, nil
	}
	// SPEED TESTING
	start := time.Now()

	//
	var working_wg sync.WaitGroup
	var collect_wg sync.WaitGroup

	// channels
	path_chan := make(chan string, 100)
	html_result_chan := make(chan struct_read.HTML_f)
	md_result_chan := make(chan struct_read.MD_f)
	// defer close(html_result_chan)
	// defer close(md_result_chan)

	// extensions
	ext_html := ".html"
	ext_md := ".md"

	var html_obj_l []struct_read.HTML_f
	var md_obj_l []struct_read.MD_f
	
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
	}()
	
	// file object manager
	working_wg.Add(1)
	go func() {
		for path := range path_chan {
			ext_f := filepath.Ext(path)
			if ext_f == ext_html {
				new_obj_f := struct_read.HTML_f{
					Base_structure_Information: manage_obj_f(path, fr_path),
				}
				html_result_chan <- new_obj_f
			}
			if ext_f == ext_md {
				new_obj_f := struct_read.MD_f{
					Base_structure_Information: manage_obj_f(path, fr_path),
				}
				md_result_chan <- new_obj_f
			}
		}
		defer working_wg.Done()
	}()

	// grouping manager
	collect_wg.Add(2)
	go func() {
		for obj := range html_result_chan {
			html_obj_l = append(html_obj_l, obj)
		}
		defer collect_wg.Done()
	}()
	go func() {
		for obj := range md_result_chan {
			md_obj_l = append(md_obj_l, obj)
		}
		defer collect_wg.Done()
	}()

	working_wg.Wait()
	close(html_result_chan)
	close(md_result_chan)
	collect_wg.Wait()

	// defer fmt.Println(html_obj_l)
	// defer fmt.Println("==============================")
	// defer fmt.Println(md_obj_l)
	
	duration := time.Since(start)
	fmt.Println("TIME : ", duration)
	//
	return html_obj_l, md_obj_l
}


func manage_obj_f(path string, fr_path string) struct_read.Base_structure_Information {
	//
	// full path, file name, file extension, folder parent, is root
	f_name := filepath.Base(path)
	f_ext := filepath.Ext(path)
	fr_parent := filepath.Base(filepath.Dir(path))
	is_root := (fr_parent == fr_path)

	// read the content of the file
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	f_content := string(data)

	return struct_read.Base_structure_Information{
		F_Path: path,
		F_name: f_name,
		F_Extension: f_ext,
		Fr_Parent: fr_parent,
		Is_Root: is_root,
		F_content: f_content,
	}
}
