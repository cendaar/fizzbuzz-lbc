package models

import "net/http"

type Fizzbuzz struct {
	ID 		int		`json:"id"`
	Int1 	int		`json:"int1"`
	Int2	int		`json:"int2"`
	Limit	int		`json:"limit"`
	Str1	string	`json:"str1"`
	Str2	string	`json:"str2"`
	Count	int		`json:"count"`
}

func (fb *Fizzbuzz) Bind(r *http.Request) error {
	return nil
}