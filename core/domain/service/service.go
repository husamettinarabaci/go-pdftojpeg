package domain

import (
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (a Service) IsConverterRequestEntityValid(converterRequest me.ConverterRequest) error {
	return converterRequest.IsValid()
}
