package domain

import (
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type Converter struct {
	Item string `json:"item"`
}

func (o Converter) ToJson() string {
	return tjson.ToJson(o)
}

func (a Converter) FromJson(i string) Converter {
	return tjson.FromJson[Converter](i)
}

func NewConverter(item string) Converter {
	return Converter{
		Item: item,
	}
}

func (o Converter) IsEmpty() bool {
	return o.ToJson() == Converter{}.ToJson()
}

func (o Converter) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Converter) IsEqual(i Converter) bool {
	return o.ToJson() == i.ToJson()
}

func (o Converter) IsValid() error {
	if o.IsEmpty() {
		return ErrInvalidInput
	}
	if o.Item == "" {
		return ErrInvalidInput
	}
	return nil
}
