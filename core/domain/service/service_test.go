package domain

import (
	"testing"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
)

func Test_IsConverterRequestEntityValid(t *testing.T) {

	var tests = []struct {
		name             string
		converterRequest me.ConverterRequest
		want             error
	}{
		{name: "empty", converterRequest: me.ConverterRequest{}, want: mo.ErrInvalidInput},
		{name: "non_values", converterRequest: me.ConverterRequest{Id: uuid.New(), Converter: mo.Converter{}}, want: mo.ErrInvalidInput},
		{name: "valid", converterRequest: me.ConverterRequest{Id: uuid.New(), Converter: mo.Converter{Item: 250, PackSizes: []int{250, 500, 1000, 2000, 5000}}}, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NewService().IsConverterRequestEntityValid(tt.converterRequest)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
