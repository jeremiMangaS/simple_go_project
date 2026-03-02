package pkg_cmd

import (
	//
)

func cmd_map() map[string]func() string {
	//
	// return the map immediately
	return map[string]func() string {
		//
		"--simple-basic" : simple_basic_handling,
		"--simple" : simple_handling,
	}
}

func simple_basic_handling() string {
	//
	return "success simple basic"
}

func simple_handling() string {
	//
	return "success simple"
}