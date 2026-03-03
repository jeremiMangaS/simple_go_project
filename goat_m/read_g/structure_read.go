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
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			fmt.Println("HTML ==> ", path)
		} else if !d.IsDir() && filepath.Ext(path) == ".md" {
			fmt.Println("MD ==> ", path)
		} else if !d.IsDir() && filepath.Ext(path) == ".json" {
			fmt.Println("JSON ==> ", path)
		}

		return err
	})
}