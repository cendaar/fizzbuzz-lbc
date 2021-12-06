package models

import (
	"errors"
	"fmt"
	"net/http"
)

type Fizzbuzz struct {
	Int1 	int		`json:"int1,omitempty"`
	Int2	int		`json:"int2,omitempty"`
	Limit	int		`json:"limit,omitempty"`
	Str1	string	`json:"str1,omitempty"`
	Str2	string	`json:"str2,omitempty"`
}

func (fb *Fizzbuzz) Bind(r *http.Request) error {
	if fb.Int1 <= 0 || fb.Int2 <= 0 || fb.Limit <= 0 || fb.Str1 == "" || fb.Str2 == "" {
		return errors.New("invalid parameters")
	}

	return nil
}

func (fb *Fizzbuzz) GetKey() string {
	return fmt.Sprint(
		"int1:", fb.Int1,
		":int2:", fb.Int2,
		":limit:", fb.Limit,
		":str1:", fb.Str1,
		":str2:", fb.Str2)
}