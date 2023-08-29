package infrastructure

import (
	"context"

	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	mp "github.com/husamettinarabaci/go-pdftojpeg/pkg/infrastructure/mapper"
)

type ConverterAdapter struct {
}

func NewConverterAdapter() ConverterAdapter {
	adapter := ConverterAdapter{}

	return adapter
}

func (a ConverterAdapter) Convert(ctx context.Context, converterRequest me.ConverterRequest) (mo.Response, error) {
	converterMapper := mp.FromConverterObject(converterRequest.Converter)
	if err := converterMapper.IsValid(); err != nil {
		return converterMapper.ToResponseObject(), err
	}
	converterMapper.Convert()
	return converterMapper.ToResponseObject(), nil
}
