package fail

import (
	"context"
	"errors"
)

func Retry(ctx context.Context, fun func(context.Context)) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	// 通过上下文传递重试次数
	retryCount, ok := ctx.Value("retryCount").(int)
	if !ok {
		retryCount = 0
	}
	if retryCount > 3 {
		return errors.New("retry too many times")
	}
	// 次数加一
	fun(context.WithValue(ctx, "retryCount", retryCount+1))
	return nil
}
