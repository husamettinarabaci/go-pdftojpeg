package domain

import tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"

type Response struct {
	Packs  []int `json:"packs"`
	Counts []int `json:"counts"`
}

func (o Response) ToJson() string {
	return tjson.ToJson(o)
}

func (a Response) FromJson(i string) Response {
	return tjson.FromJson[Response](i)
}

func NewResponse(packs []int, counts []int) Response {
	return Response{
		Packs:  packs,
		Counts: counts,
	}
}

func (o Response) IsEmpty() bool {
	return o.ToJson() == Response{}.ToJson()
}

func (o Response) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Response) IsEqual(i Response) bool {
	return o.ToJson() == i.ToJson()
}

func (o Response) IsValid() error {
	if o.IsEmpty() {
		return ErrInvalidInput
	}
	if o.Packs == nil {
		return ErrInvalidInput
	}
	if len(o.Packs) == 0 {
		return ErrInvalidInput
	}
	for _, v := range o.Packs {
		if v <= 0 {
			return ErrInvalidInput
		}
	}
	for _, v := range o.Counts {
		if v <= 0 {
			return ErrInvalidInput
		}
	}
	return nil
}
