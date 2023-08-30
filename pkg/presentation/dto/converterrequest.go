package presentation

import (
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterRequest struct {
	Item string `json:"item" form:"item" binding:"required"`
}

func (a ConverterRequest) ToJson() string {
	return tjson.ToJson(a)
}

func (e ConverterRequest) FromJson(i string) ConverterRequest {
	return tjson.FromJson[ConverterRequest](i)
}

func NewConverterRequest(item string) ConverterRequest {
	return ConverterRequest{
		Item: item,
	}
}

func (o ConverterRequest) IsEmpty() bool {
	return o.ToJson() == ConverterRequest{}.ToJson()
}

func (a ConverterRequest) ToConverterRequestEntity() me.ConverterRequest {
	return me.NewConverterRequest(
		mo.NewConverter(
			a.Item,
		),
	)
}

func (a *ConverterRequest) IsValid() error {
	if a.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if a.Item == "" {
		return mo.ErrInvalidInput
	}
	return nil
}
