package application

import (
	"context"

	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
)

type CommandPort interface {
	Convert(ctx context.Context, converterRequest me.ConverterRequest) (me.ConverterResponse, error)
}
