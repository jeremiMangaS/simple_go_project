package main

import (
	"fmt"
	"os"
	"simple_go_project/commands"
)


/*
	Data Read, Write and Execute
	Where : read - 4
			write - 2
			execute - 1
	combination : 
	0755 
		0 - for oktal
		7 = 4 (read) + 2 (write) + 1 (execute)	-> Owner | rwx
		5 = 4 (read) + 1 (execute)				-> group | r-x
		5 = 4 (read) + 1 (execute)				-> Other | r-x
*/

func main() {
	// Taking all the argumen/command as a string[] type
	m_args := os.Args[1:]

	if len(m_args) == 0 {
		fmt.Println("[Missing command] : try use 'goat init' or 'goat build'")
		return
	}

	/*
		Taking the first index of command as the main command ( m_command )
		-> goat
	*/
	m_command := m_args[0]
	// s_command := m_args[1]
	
	// fmt.Println(m_command, " : ", s_command)

	switch m_command {
	case "init" :
		init_command()
	case "build" :
		build_command()
	default :
		fmt.Print("This is default value")
	}
}

// func init_command() {
// 	// command input (full)
// 	m_args := os.Args[1:]
// 	// simple flag
// 	flag_smplbsc := "--simple-basic"
// 	flag_smpl := "--simple"
// 	// input flag from command 
// 	// taking the 2 index for the flag
// 	i_flag := m_args[1]
// 	//
// 	// creating project for simple basic flag
// 	if i_flag == flag_smplbsc {
// 		create_files("index.html", simple_HTML)
// 		create_files("index.md", simple_MD)
// 		create_files("index.json", "{}")
// 		fmt.Println("Successfully created project folder\nindex.html\nindex.md\nindex.json")
// 	// creating project for simple flag
// 	} else if i_flag == flag_smpl {
// 		os.Mkdir("folder_asset", 0755)
// 		os.Mkdir("folder_component", 0755)
// 		os.Mkdir("folder_page", 0755)
// 		create_files("main_inf.json", "{}")
// 		create_files("main_inf.html", simple_HTML)
// 		create_files("main_inf.md", simple_MD)	
// 		fmt.Println("Successfully created project folder\nfolder_assets/\nfolder_component/\nfolder_page/\nmain_inf.html\nmain_inf.md\nmain_inf.json")	
// 	}
// }

func init_command() {
	// taking input
	m_args := os.Args[1:]
	flag_cmmd := m_args[1]
	// fmt.Println("\n\n\nHASIL : ", flag_cmmd)
	result := command.Process_input(flag_cmmd)
	fmt.Println(result)
}


func build_command() {
	//
}