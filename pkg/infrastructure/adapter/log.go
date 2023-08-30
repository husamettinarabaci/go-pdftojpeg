package infrastructure

import (
	"context"

	mi "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/interface"
	tconfig "github.com/husamettinarabaci/go-pdftojpeg/tool/config"
)

type LogAdapter struct {
}

func NewLogAdapter() LogAdapter {
	adapter := LogAdapter{}
	return adapter
}

func (a LogAdapter) Log(ctx context.Context, source string, logData mi.Loggable) {
	if tconfig.GetLogConfigInstance().Logger.Console {

	}
}
