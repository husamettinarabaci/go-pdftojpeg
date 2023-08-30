package application

import (
	"context"

	ip "github.com/husamettinarabaci/go-pdftojpeg/core/application/infrastructure/port"
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mi "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/interface"
	ds "github.com/husamettinarabaci/go-pdftojpeg/core/domain/service"
)

type Service struct {
	domainService ds.Service
	converter     ip.ConverterPort
	logger        ip.LogPort
}

func NewService(domainService ds.Service, converterPort ip.ConverterPort, logger ip.LogPort) Service {
	return Service{
		domainService: domainService,
		converter:     converterPort,
		logger:        logger,
	}
}

func (a Service) Log(ctx context.Context, operationName string, logData mi.Loggable) {
	a.logger.Log(ctx, operationName, logData)
}

func (a Service) Convert(ctx context.Context, converterRequest me.ConverterRequest) (me.ConverterResponse, error) {
	operationName := "ExecuteConverter"
	a.Log(ctx, operationName, converterRequest)
	if err := a.domainService.IsConverterRequestEntityValid(converterRequest); err != nil {
		return me.ConverterResponse{
			Id: converterRequest.Converter.Item,
		}, err
	}
	response, err := a.converter.Convert(ctx, converterRequest)
	converterResponse := me.NewConverterResponse(converterRequest.Converter.Item, response)
	a.Log(ctx, operationName, converterResponse)
	return converterResponse, err
}
