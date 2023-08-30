package infrastructure

import (
	"bytes"
	"image/jpeg"

	"github.com/husamettinarabaci/go-fitz"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
)

type Converter struct {
	Item    string   `json:"item"`
	Results [][]byte `json:"results"`
}

func (a Converter) ToJson() string {
	return tjson.ToJson(a)
}

func (e Converter) FromJson(i string) Converter {
	return tjson.FromJson[Converter](i)
}

func NewConverter(item string) Converter {
	return Converter{
		Item: item,
	}
}

func FromConverterObject(converter mo.Converter) Converter {
	return NewConverter(
		converter.Item,
	)
}

func (a Converter) IsValid() error {
	if a.Item == "" {
		return mo.ErrInvalidInput
	}
	return nil
}

func (a Converter) ToResponseObject() mo.Response {
	return mo.NewResponse(a.Results)
}

func (a *Converter) Convert() error {
	doc, err := fitz.New("./tmp/" + a.Item + ".pdf")
	if err != nil {
		return err
	}
	defer doc.Close()
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			continue
		}
		var b []byte
		var bWriter = bytes.NewBuffer(b)
		err = jpeg.Encode(bWriter, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			continue
		}

		a.Results = append(a.Results, bWriter.Bytes())
	}
	return nil
}
