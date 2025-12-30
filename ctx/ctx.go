package common

import (
	"context"
	"fmt"
	"time"

	"github.com/marcuspeh/go-tools/logger"
)

func GetCtx(postfix string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	logID := fmt.Sprintf("%s_%s_%v",
		time.Now().UTC().Format(time.DateOnly),
		time.Now().UTC().Format(time.TimeOnly),
		postfix,
	)
	fmt.Println("LogID: ", logID)

	ctx = context.WithValue(ctx, logger.LogIDKey, logID)
	return ctx, cancel
}
