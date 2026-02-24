package main

import (
	"bufio"
	"fmt"
	"strings"
)

const file_sm = `$key_file : index.html
title : welcome to GOat
t_ct.judul1 : introduction
b_ct.body1 : GOat is a simple framework`

func main() { 
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
		key_value_data := strings.SplitN(data_t, ":", 2)

		// make sure that the data is really only divided into two 
		if len(key_value_data) == 2 {
			// key_data for key code
			// value_data for value the information 
			key_data := strings.TrimSpace(key_value_data[0])
			value_data := strings.TrimSpace(key_value_data[1])

			// add/push the key and value to the map
			data_map[key_data] = value_data
		}
	}

	fmt.Println("===================================")
	for key, value := range data_map {
		fmt.Printf("Key : %s  ||  value : %s\n", key, value)
	}
}