package scv

type StringCase int

const (
	Snake StringCase = iota
	Kebab
	UpperCamel
	LowerCamel
)
