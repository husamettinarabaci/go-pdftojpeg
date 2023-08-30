package presentation

import (
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterResponse struct {
	Files [][]byte `json:"files"`
}

func (a ConverterResponse) ToJson() string {
	return tjson.ToJson(a)
}

func (e ConverterResponse) FromJson(i string) ConverterResponse {
	return tjson.FromJson[ConverterResponse](i)
}

func NewConverterResponse(files [][]byte) ConverterResponse {
	return ConverterResponse{
		Files: files,
	}
}

func (o ConverterResponse) IsEmpty() bool {
	return o.ToJson() == ConverterResponse{}.ToJson()
}

func FromResponseObject(response mo.Response) ConverterResponse {
	return NewConverterResponse(response.Files)
}
