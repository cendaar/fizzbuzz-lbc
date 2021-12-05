package models

import (
	"net/http"
)

type Fizzbuzz struct {
	ID 		int		`json:"id,omitempty"`
	Int1 	int		`json:"int1,omitempty"`
	Int2	int		`json:"int2,omitempty"`
	Limit	int		`json:"limit,omitempty"`
	Str1	string	`json:"str1,omitempty"`
	Str2	string	`json:"str2,omitempty"`
	Count	int		`json:"count,omitempty"`
}

func (fb *Fizzbuzz) Bind(r *http.Request) error {
	return nil
}