package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// IncrVisitCount 访问次数加一
func (r *Repository) IncrVisitCount(ctx context.Context, visitorID string) (int64, error) {
	val, err := r.IncrBy(ctx, fmt.Sprintf("advice:visitor.record:%s", visitorID), 1).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// GetVisitCount 获取总访问次数
func (r *Repository) GetVisitCount(ctx context.Context) (int64, error) {
	iter := r.Scan(ctx, 0, "advice:visitor.record:*", 0).Iterator()
	pipe := r.Pipeline()
	for iter.Next(ctx) {
		pipe.Get(ctx, iter.Val())
	}
	result, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}
	var count int64
	for _, v := range result {
		val, _ := v.(*redis.StringCmd).Int64()
		count += val
	}
	return count, nil
}
