package main

import (
	"bufio"
	"fmt"
	"strings"
)

// content example
const file_sm = `$key_file : index.html
title : welcome to GOat
t_ct.judul1 : introduction
b_ct.body1 : GOat is a simple framework
continue from the body1 tag
t_ct.title2 : new title here`

// How to solve new line value ???

func main() {
	// short term memory so that it can refer to the previous key if the key 
	// is not found in the next row 
	var last_key string

	// map to store keys and values ​​so that they can be processed later as logic
	data_map := make(map[string]string)

	// NewReader wraps the data in strings.NewReaxer(data) so that it can 
	// be read by the scanner, the scanner puts the data into a buffer.
	scanner := bufio.NewScanner(strings.NewReader(file_sm))

	// reading the data from scanner buffer
	for scanner.Scan() {
		// fmt.Printf("%T", scanner)
		// fmt.Println(scanner, "\n\n")

		// converting the data from *bufio.Scanner*bufio.Scanner*bufio.Scanner*bufio.Scanner 
		// into string
		data_t := scanner.Text()
		// fmt.Printf("%T", data_t)
		// fmt.Println("\n\n", data_t)

		// ignore an empty line or the barrier "---" and just continue to the next logic
		if strings.TrimSpace(data_t) == "---" || strings.TrimSpace(data_t) == "" {
			continue
		}

		// split the data into the key and value to know between the keycode and value are 
		// the information that is filled in
		// key_value_data := strings.SplitN(data_t, ":", 2)

		// note : still using the same logic as before, just made a little adjustment
		if strings.Contains(data_t, ":") {
			key_value_data := strings.SplitN(data_t, ":", 2)
			key_data := strings.TrimSpace(key_value_data[0])
			value_data := strings.TrimSpace(key_value_data[1])
			// add/push the key and value to the map
			data_map[key_data] = value_data
			// refer to the current key / the last key for the next process
			last_key = key_data
		} else {
			// to solve the new line problem, so that the new line value will be added 
			// to the last key from last_key variable
			c_data := strings.TrimSpace(data_t)
			if data_t != "" && last_key != "" {
				// data_map[last_key] = c_data

				// add value from next row to previous locked
				data_map[last_key] = data_map[last_key] + " " + c_data
			}

		}

		// make sure that the data is really only divided into two 
		// if len(key_value_data) == 2 {
		// 	// key_data for key code
		// 	// value_data for value the information 
		// 	key_data := strings.TrimSpace(key_value_data[0])
		// 	value_data := strings.TrimSpace(key_value_data[1])

		// 	// add/push the key and value to the map
		// 	data_map[key_data] = value_data
		// }
	}

	fmt.Println("===================================")
	for key, value := range data_map {
		fmt.Printf("Key : %s  ||  value : %s\n", key, value)
	}
}