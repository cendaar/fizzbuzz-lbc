package services

import (
	"github.com/cendaar/fizzbuzz/models"
	"strconv"
)

type FizzbuzzService struct {
}

func NewFizzbuzzService() *FizzbuzzService {
	return &FizzbuzzService{}
}

func (fs *FizzbuzzService) ComputeFizzbuzz(fb *models.Fizzbuzz) string {
	var output string

	for i := 1; i <= fb.Limit; i++ {
		switch {
		case i % (fb.Int1 * fb.Int2) == 0:
			output += fb.Str1 + fb.Str2
		case i % fb.Int1 == 0:
			output += fb.Str1
		case i % fb.Int2 == 0:
			output += fb.Str2
		default:
			output += strconv.Itoa(i)
		}

		if i != fb.Limit {
			output += ","
		}
	}

	return output
}