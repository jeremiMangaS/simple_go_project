// ================================================== [ HANDLING FUNCTION ] ==================================================
package main

import (
	//
	"fmt"
	"os"
	"github.com/jeremiMangaS/simple_go_project/goat_m/main_g/cmd_handling"
	// "github.com/jeremiMangaS/simple_go_project/goat_m/build_g"
	"github.com/jeremiMangaS/simple_go_project/goat_m/read_g"
	core_prc "github.com/jeremiMangaS/simple_go_project/goat_m/core_g/process_g"
)



func init_handling() {
	//
	m_args := os.Args[1:]
	// checking
	if len(m_args) <= 1 {
		fmt.Print("no flag found")
		return
	}

	// taking the flag
	m_cmd := m_args[1]
	result := pkg_cmd.Check_flag(m_cmd)
	fmt.Print(result)
}



func build_handling() {
	// fmt.Println("build...")
	//
	m_args := os.Args[1:]
	// checking
	if len(m_args) <= 1 {
		fmt.Print("no path found")
		return
	}

	// taking the path
	path_args := m_args[1]
	html_obj_l, md_obj_l := read_g.Central_analys_g(path_args)
	//
	core_prc.Prc_parsing(html_obj_l, md_obj_l)
}




func help_handling() {
	fmt.Println("--help : testing")
}
