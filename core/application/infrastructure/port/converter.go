package application

import (
	"context"

	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
)

type ConverterPort interface {
	Convert(ctx context.Context, converterRequest me.ConverterRequest) (mo.Response, error)
}
