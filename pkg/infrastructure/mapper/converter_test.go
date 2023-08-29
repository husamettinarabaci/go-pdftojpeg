package infrastructure

import (
	"reflect"
	"testing"
)

func Test_ConvertCumulative(t *testing.T) {

	var packSizes = []int{250, 500, 1000, 2000, 5000}
	var tests = []struct {
		name          string
		orderQuantity int
		want          []int
	}{
		{name: "Items ordered:     1", orderQuantity: 1, want: []int{250}},
		{name: "Items ordered:   250", orderQuantity: 250, want: []int{250}},
		{name: "Items ordered:   251", orderQuantity: 251, want: []int{500}},
		{name: "Items ordered:   501", orderQuantity: 501, want: []int{500, 250}},
		{name: "Items ordered:   751", orderQuantity: 751, want: []int{1000}},
		{name: "Items ordered: 12001", orderQuantity: 12001, want: []int{5000, 5000, 2000, 250}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converterulator := NewConverter(tt.orderQuantity, packSizes)
			converterulator.Convert()
			ans := converterulator.Results

			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
