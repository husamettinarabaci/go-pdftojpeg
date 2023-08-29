package domain

import (
	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterRequest struct {
	Id        uuid.UUID    `json:"id"`
	Converter mo.Converter `json:"converter"`
}

func (e ConverterRequest) ToJson() string {
	return tjson.ToJson(e)
}

func (a ConverterRequest) FromJson(i string) ConverterRequest {
	return tjson.FromJson[ConverterRequest](i)
}

func NewConverterRequest(id uuid.UUID, converter mo.Converter) ConverterRequest {
	return ConverterRequest{
		Id:        id,
		Converter: converter,
	}
}

func (o ConverterRequest) IsEmpty() bool {
	return o.ToJson() == ConverterRequest{}.ToJson()
}

func (o ConverterRequest) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o ConverterRequest) IsEqual(i ConverterRequest) bool {
	return o.ToJson() == i.ToJson()
}

func (o ConverterRequest) IsValid() error {
	if o.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if o.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if o.Converter.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if err := o.Converter.IsValid(); err != nil {
		return err
	}
	return nil
}
