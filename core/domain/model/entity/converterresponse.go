package domain

import (
	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterResponse struct {
	Id       uuid.UUID   `json:"id"`
	Response mo.Response `json:"response"`
}

func (e ConverterResponse) ToJson() string {
	return tjson.ToJson(e)
}

func (a ConverterResponse) FromJson(i string) ConverterResponse {
	return tjson.FromJson[ConverterResponse](i)
}

func NewConverterResponse(id uuid.UUID, response mo.Response) ConverterResponse {
	return ConverterResponse{
		Id:       id,
		Response: response,
	}
}

func (o ConverterResponse) IsEmpty() bool {
	return o.ToJson() == ConverterResponse{}.ToJson()
}

func (o ConverterResponse) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o ConverterResponse) IsEqual(i ConverterResponse) bool {
	return o.ToJson() == i.ToJson()
}

func (o ConverterResponse) IsValid() error {
	if o.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if o.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if o.Response.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if err := o.Response.IsValid(); err != nil {
		return err
	}
	return nil
}
