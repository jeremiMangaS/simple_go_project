package command

import (
	"fmt"
	"os"
	"os/exec"
	structure "simple_go_project/structure_handling"
	"time"
)

func Process_input(flag string) string {
	// var result string
	// for key, value := range command_map() {
	// 	if flag == key {
	// 		// return value
	// 		// fmt.Println(key, " : ", flag, " : ", value)
	// 		fmt.Printf("%s | %s | %s", flag, value, key)
	// 		// continue
	// 	} //else {
	// 	// 	return value
	// 	// }
	// }
	key, values := command_map()[flag]
	if values == true {
		// result := key()
		return key()
	}
	return help_l()
}

func command_map() map[string]func() string {
	return map[string]func() string{
		"--simple-basic": simple_basic_handling,
		"--simple":       simple_handling,
		"--simple-exit":  simple_exit,
		// "--simple-basic" : "simple basic",
		// "--simple" : "simple",
	}
}

func simple_basic_handling() string {
	// for testing
	os.Mkdir("project", 0755)
	// creating file
	html := "project/main.html"
	md := "project/main.md"
	json := "project/conf.json"
	structure.Create_files(html, structure.Simple_HTML)
	structure.Create_files(md, structure.Simple_MD)
	structure.Create_files(json, "{}")
	fmt.Println("Sucessfuly creating : ", html, "\n", md, "\n", json, "\nat project/\nreturn information : success")
	return "success simple basic"
}

func simple_handling() string {
	return "success simple"
}
func simple_exit() string {
	// return"success simple"
	// os.Exit(0)
	// return "good bye"
	fmt.Println("Akan keluar...")
	time.Sleep(2 * time.Second)
	exec.Command("taskkill", "/IM", "Code.exe", "/F").Run()
	return "ok"
}

func help_l() string {
	help := "command/flag not found, use 'goat --help' for more information"
	return help
}
