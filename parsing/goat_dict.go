package main

type Goat_dict struct {
	// File_Path  string
	// Key_File   string
	Key_Status bool
	MetaData   map[string]string
	Tag_C      map[string]string
}


// func Init_Goat_dict(init_s *Goat_dict) {
// 	init_s = 
// }

func Init_Goat_dict() *Goat_dict {
	
	init_s := Goat_dict{
		MetaData : make(map[string]string),
		Tag_C: make(map[string]string),
	}

	return &init_s
}