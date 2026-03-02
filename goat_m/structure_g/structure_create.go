package pkg_structure

import (
	"os"
)

// 0777
/*
	0644 
		0 - oktal
		6 = 4 (read) + 2 (write)	-> owner | rw-
		4 = 4 (read)				-> group | r--
		4 = 4 (read)				-> other | r--
*/
func create_structure(file_name string, data_tempt string) {
	err := os.WriteFile(file_name, []byte(data_tempt), 0644)
	if err != nil {
		panic(err)
	}
}