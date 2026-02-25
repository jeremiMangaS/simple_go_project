package main

import (
	"fmt"
	"os"
)

func command_map() map[string]string {
	// Map for holding all the command flags
	command_list := map[string]string{
		// Values of the command flags
		"--simple-basic": simple_bsc_handling(),
		"--simple":       simple_handling(),
	}
	return command_list
}

func main() {
	args := os.Args[1:]
	var result string
	// if len(args) == 0 { return }
	for key, value := range command_map() {
		if args[0] == key {
			result = value
		} //  else {
		// 	fmt.Println(args[0], " : ", key)
		// }
	}
	fmt.Print(result)
	// fmt.Printf("%T", args)
	// fmt.Printf("Result : \n%s : %s\n\n", command_map()["--simple-basic"], command_map()["--simple-basic"])
	// for key, value := range command_map() {
	// 	fmt.Println(key, " : ", value)
	// }
}

func simple_bsc_handling() string {
	return "succeess from simple basic"
}

func simple_handling() string {
	return "success from simple"
}
