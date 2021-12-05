package services

import (
	"github.com/baqtiste/fizzbuzz/db"
	"github.com/baqtiste/fizzbuzz/models"
	"strconv"
)

type FizzbuzzService struct {
	RedisInstance *db.RedisInstance
}

func NewFizzbuzzService(ri *db.RedisInstance) *FizzbuzzService {
	return &FizzbuzzService{RedisInstance: ri}
}

func (fs *FizzbuzzService) ComputeFizzbuzz(mfb *models.Fizzbuzz) string {
	var output string

	for i := 1; i <= mfb.Limit; i++ {
		switch {
		case i % (mfb.Int1 * mfb.Int2) == 0:
			output += mfb.Str1 + mfb.Str2
		case i % mfb.Int1 == 0:
			output += mfb.Str1
		case i % mfb.Int2 == 0:
			output += mfb.Str2
		default:
			output += strconv.Itoa(i)
		}

		if i != mfb.Limit {
			output += ","
		}
	}

	return output
}