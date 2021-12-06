package services

import (
	"context"
	"encoding/json"
	"github.com/cendaar/fizzbuzz/db"
	"github.com/cendaar/fizzbuzz/models"
	"github.com/cendaar/fizzbuzz/repositories"
	"strconv"
)

type StatsService struct {
	repository *repositories.StatsRepository
}

func NewStatsService(r *db.RedisInstance) *StatsService {
	return &StatsService{repository: repositories.NewStatsRepository(r)}
}

func (ss *StatsService) GetFormattedRequestLeader() (string, error) {
	ctx := context.Background()

	requestLeader, score, err := ss.repository.GetRequestLeader(ctx)
	if err != nil {
		return "", err
	}

	switch {
	case err != nil:
		return "", err
	case requestLeader == nil || score == 0:
		return "There are no request recorded yet.", nil
	case score == 1:
		return "The most used request is " + *requestLeader + " with a total of " + strconv.Itoa(score) + " hit.", nil
	default:
		return "The most used request is " + *requestLeader + " with a total of " + strconv.Itoa(score) + " hits.", nil
	}
}

func (ss *StatsService) HandleFizzbuzzRequest(fizzbuzz *models.Fizzbuzz) error {
	ctx := context.Background()
	hashkey := fizzbuzz.GetKey()
	JSONRequest, _ := json.Marshal(fizzbuzz)

	exists, err := ss.repository.IsRequestHashkeyExists(ctx, hashkey)
	if err != nil {
		return err
	}

	if !exists {
		err = ss.repository.InsertRequest(ctx, hashkey, JSONRequest)
	} else {
		err = ss.repository.IncrementRequest(ctx, hashkey)
	}

	return err
}