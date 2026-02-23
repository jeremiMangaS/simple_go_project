package command

import "os"

func Process_input(flag string) string {
	// var result string
	for key, value := range command_map() {
		if flag == key {
			return value
		}
	}
	return help_l()
} 

func command_map() map[string]string {
	cmmnd_map := map[string] string {
		"--simple-basic" : simple_basic_handling(),
		"--simple" : simple_handling(),
	}
	return cmmnd_map
}

func simple_basic_handling() string {
	return "success simple basic"
}

func simple_handling() string {
	return"success simple"
}

func help_l() string {
	help := "command/flag not found, use 'goat --help' for more information"
	return help
}

func create_files(file_name string, data string) {
	// 0777
	/*
		0644 
			0 - oktal
			6 = 4 (read) + 2 (write)	-> owner | rw-
			4 = 4 (read)				-> group | r--
			4 = 4 (read)				-> other | r--
	*/
	err := os.WriteFile(file_name, []byte(data), 0644)
	check(err)
}

func check(e error) {
	// checking error
	if e != nil {
		panic(e)
	}
}