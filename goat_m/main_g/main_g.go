package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Print("goat v.1.0.1")

	// taking the arguments and flags
	m_args := os.Args[1:]
	if len(m_args) <= 0 {
		help()
		return 
	}
	// taking the first argumen as main command
	m_cmd := m_args[0]

	switch m_cmd {
	case "init" :
		//
		init_handling()
	case "build" :
		//
		build_handling()
	case "--help" :
		//
		help_handling()
	default : 
		fmt.Println("Command/flag not found\nRun --help for more information")
	}
}

func help() {
	fmt.Println("command/flag not found\nRun --help for more information")
}