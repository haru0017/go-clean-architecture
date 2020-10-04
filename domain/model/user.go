package model

// User はidとnameを持つ型
type User struct {
	ID int
	Name string
}

// Users はUserのスライス
type Users []User