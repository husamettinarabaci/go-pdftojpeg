package domain

import (
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type Converter struct {
	Item      int   `json:"item"`
	PackSizes []int `json:"pack_sizes"`
}

func (o Converter) ToJson() string {
	return tjson.ToJson(o)
}

func (a Converter) FromJson(i string) Converter {
	return tjson.FromJson[Converter](i)
}

func NewConverter(item int, packSizes []int) Converter {
	return Converter{
		Item:      item,
		PackSizes: packSizes,
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
	if o.Item <= 0 {
		return ErrInvalidInput
	}
	if o.PackSizes == nil {
		return ErrInvalidInput
	}
	if len(o.PackSizes) == 0 {
		return ErrInvalidInput
	}
	for _, v := range o.PackSizes {
		if v <= 0 {
			return ErrInvalidInput
		}
	}
	return nil
}

var DefaultPackSizes = []int{250, 500, 1000, 2000, 5000}
