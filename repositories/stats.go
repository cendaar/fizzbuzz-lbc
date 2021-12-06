package repositories

import (
	"context"
	"github.com/cendaar/fizzbuzz/db"
	"github.com/go-redis/redis/v8"
)

const (
	keyPrefixRequestHashKey		= "request:"
	keyJSONRequest				= "json_request"
	keyRequestLeaderboard 		= "fizzbuzz_requests"
)

type StatsRepository struct {
	redis *db.RedisInstance
}

func NewStatsRepository(redis *db.RedisInstance) *StatsRepository {
	return &StatsRepository{redis: redis}
}

func (sr *StatsRepository) GetRequestLeader(ctx context.Context) (*string, int, error) {
	members, err := sr.redis.Client.ZRevRange(ctx, keyRequestLeaderboard, 0, 0).Result()
	if err != nil {
		return nil, 0, err
	}

	// if there are no member for this leaderboard
	if len(members) < 1 {
		return nil, 0, err
	}

	requestLeader, err := sr.getJSONRequestFromHashkey(ctx, members[0])
	if err != nil {
		return nil, 0, err
	}

	score, err := sr.redis.Client.ZMScore(ctx, keyRequestLeaderboard, members[0]).Result()
	if err != nil {
		return nil, 0, err
	}

	return &requestLeader, int(score[0]), nil
}

func (sr *StatsRepository) getJSONRequestFromHashkey(ctx context.Context, hashkey string) (string, error) {
	return sr.redis.Client.HGet(ctx, keyPrefixRequestHashKey+hashkey, keyJSONRequest).Result()
}

func (sr *StatsRepository) IncrementRequest(ctx context.Context, hashkey string) error {
	_, err := sr.redis.Client.ZIncrBy(ctx, keyRequestLeaderboard, 1, hashkey).Result()
	return err
}

func (sr *StatsRepository) InsertRequest(ctx context.Context, hashkey string, r []byte) error {
	_, err := sr.redis.Client.HSet(ctx, keyPrefixRequestHashKey+hashkey, keyJSONRequest, r).Result()
	if err != nil {
		return err
	}

	_, err = sr.redis.Client.ZAdd(ctx, keyRequestLeaderboard, &redis.Z{Score:  1, Member: hashkey}).Result()
	return err
}

func (sr *StatsRepository) IsRequestHashkeyExists(ctx context.Context, hashkey string) (bool, error) {
	return sr.redis.Client.HExists(ctx, keyPrefixRequestHashKey+hashkey, keyJSONRequest).Result()
}