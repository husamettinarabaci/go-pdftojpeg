package presentation

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type ConverterRequest struct {
	Item         int    `json:"item" form:"item" binding:"required"`
	PackSizes    []int  `json:"pack_sizes"`
	PackSizesStr string `json:"pack_sizes_str" form:"pack_sizes_str"`
}

func (a ConverterRequest) ToJson() string {
	return tjson.ToJson(a)
}

func (e ConverterRequest) FromJson(i string) ConverterRequest {
	return tjson.FromJson[ConverterRequest](i)
}

func NewConverterRequest(item int, packSizes []int) ConverterRequest {
	return ConverterRequest{
		Item:      item,
		PackSizes: packSizes,
	}
}

func (o ConverterRequest) IsEmpty() bool {
	return o.ToJson() == ConverterRequest{}.ToJson()
}

func (a ConverterRequest) ToConverterRequestEntity() me.ConverterRequest {
	if a.PackSizes == nil {
		a.PackSizes = mo.DefaultPackSizes
	}
	return me.NewConverterRequest(
		uuid.New(),
		mo.NewConverter(
			a.Item,
			a.PackSizes,
		),
	)
}

func (a *ConverterRequest) FillPackSizes() {
	if a.PackSizesStr != "" {
		sizes := strings.Split(a.PackSizesStr, ",")
		a.PackSizes = make([]int, 0)
		for _, v := range sizes {
			if v != "" {
				size, err := strconv.Atoi(v)
				if err != nil {
					a.PackSizes = make([]int, 0)
					break
				}
				a.PackSizes = append(a.PackSizes, size)
			}
		}
	}
	if a.PackSizes == nil {
		a.PackSizes = mo.DefaultPackSizes
	}
	if len(a.PackSizes) == 0 {
		a.PackSizes = mo.DefaultPackSizes
	}
}

func (a *ConverterRequest) IsValid() error {
	if a.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if a.Item <= 0 {
		return mo.ErrInvalidInput
	}
	if a.PackSizes == nil {
		return mo.ErrInvalidInput
	}
	if len(a.PackSizes) == 0 {
		return mo.ErrInvalidInput
	}
	for _, v := range a.PackSizes {
		if v <= 0 {
			return mo.ErrInvalidInput
		}
	}
	return nil
}
