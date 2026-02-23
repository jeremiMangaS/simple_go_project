package structure 

import "os"

func Create_files(file_name string, data string) {
	// 0777
	/*
		0644 
			0 - oktal
			6 = 4 (read) + 2 (write)	-> owner | rw-
			4 = 4 (read)				-> group | r--
			4 = 4 (read)				-> other | r--
	*/
	err := os.WriteFile(file_name, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

// func check(e error) {
// 	// checking error
// 	if e != nil {
// 		panic(e)
// 	}
// }