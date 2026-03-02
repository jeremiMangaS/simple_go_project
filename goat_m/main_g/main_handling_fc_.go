// ================================================== [ HANDLING FUNCTION ] ==================================================
package main

import (
	//
	"fmt"
	"os"
	"github.com/jeremiMangaS/simple_go_project/goat_m/main_g/cmd_handling"
)



func init_handling() {
	//
	// taking the flag
	m_args := os.Args[1:]
	// checking
	if len(m_args) <= 1 {
		fmt.Print("no flag found")
		return
	}

	m_cmd := m_args[1]
	result := pkg_cmd.Check_flag(m_cmd)
	fmt.Print(result)
}



func build_handling() {
	fmt.Println("build...")
}




func help_handling() {
	fmt.Println("--help : testing")
}
