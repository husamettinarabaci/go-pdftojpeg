package presentation

import (
	"fmt"

	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterResponse struct {
	Packs []string `json:"packs"`
}

func (a ConverterResponse) ToJson() string {
	return tjson.ToJson(a)
}

func (e ConverterResponse) FromJson(i string) ConverterResponse {
	return tjson.FromJson[ConverterResponse](i)
}

func NewConverterResponse(packs []string) ConverterResponse {
	return ConverterResponse{
		Packs: packs,
	}
}

func (o ConverterResponse) IsEmpty() bool {
	return o.ToJson() == ConverterResponse{}.ToJson()
}

func FromResponseObject(response mo.Response) ConverterResponse {
	var packs []string
	for i := 0; i < len(response.Packs); i++ {
		packs = append(packs, fmt.Sprintf("%d x %d", response.Counts[i], response.Packs[i]))
	}
	return NewConverterResponse(
		packs,
	)
}
