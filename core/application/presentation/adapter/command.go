package application

import (
	"context"

	as "github.com/husamettinarabaci/go-pdftojpeg/core/application/service"
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
)

type CommandAdapter struct {
	service as.Service
}

func NewCommandAdapter(s as.Service) CommandAdapter {
	return CommandAdapter{
		service: s,
	}
}

func (a CommandAdapter) Convert(ctx context.Context, converterRequest me.ConverterRequest) (me.ConverterResponse, error) {
	return a.service.Convert(ctx, converterRequest)
}
