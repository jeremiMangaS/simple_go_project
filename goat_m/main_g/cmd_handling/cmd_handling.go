package pkg_cmd

import (
	//
)

func Check_flag(flag string) string {
	key, value := init_map()[flag]
	if value == true {
		return key()
	}
	return "--help"
}