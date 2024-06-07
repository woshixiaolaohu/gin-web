package internal

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/event"
	opt "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type mongo struct {
}

var Mongo = new(mongo)

func (m *mongo) GetClientOptions() []options.ClientOptions {
	cmdMonitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			zap.L().Info(fmt.Sprintf("[MongoDB][RequestID:%d][database:%s] %s\n", startedEvent.RequestID, startedEvent.DatabaseName, startedEvent.CommandName), zap.String("business", "mongo"))
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			zap.L().Info(fmt.Sprintf("[MongoDB][RequestID:%d] [%d] %s\n", succeededEvent.RequestID, succeededEvent.DurationNanos, succeededEvent.Reply), zap.String("business", "mongo"))
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			zap.L().Info(fmt.Sprintf("[MongoDB][RequestID:%d] [%d] %s\n", failedEvent.RequestID, failedEvent.DurationNanos, failedEvent.Failure), zap.String("business", "mongo"))
		},
	}
	return []options.ClientOptions{
		{
			ClientOptions: &opt.ClientOptions{
				Monitor: cmdMonitor,
			},
		},
	}
}
