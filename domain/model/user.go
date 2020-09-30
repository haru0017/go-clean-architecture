package model

// User はidとnameを持つ型
type User struct {
	id int
	name string
}

// Users はUserのスライス
type Users []User