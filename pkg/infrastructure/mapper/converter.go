package infrastructure

import (
	"sort"

	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-pdftojpeg/tool/json"
	tslice "github.com/husamettinarabaci/go-pdftojpeg/tool/slice"
)

var tSlice tslice.TSlice

type Converter struct {
	Item      int   `json:"item"`
	PackSizes []int `json:"pack_sizes"`
	Results   []int `json:"results"`
}

func (a Converter) ToJson() string {
	return tjson.ToJson(a)
}

func (e Converter) FromJson(i string) Converter {
	return tjson.FromJson[Converter](i)
}

func NewConverter(item int, packSizes []int) Converter {
	return Converter{
		Item:      item,
		PackSizes: packSizes,
	}
}

func FromConverterObject(converter mo.Converter) Converter {
	return NewConverter(
		converter.Item,
		converter.PackSizes,
	)
}

func (a Converter) IsValid() error {
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

func (a Converter) ToResponseObject() mo.Response {
	var packs []int
	var counts []int
	resMap := make(map[int]int)
	for _, pack := range a.Results {
		resMap[pack] += 1
	}
	for k, v := range resMap {
		packs = append(packs, k)
		counts = append(counts, v)
	}

	return mo.NewResponse(
		packs,
		counts,
	)
}

func (a *Converter) Convert() []int {
	a.Results = a.ConvertCumulative(a.Item, a.PackSizes)
	return a.Results
}

type Tuple struct {
	Packs int
	Combo []int
}

func (a Converter) ConvertCumulative(orderQty int, packSizes []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	maxPackSize := 0
	for _, pack := range packSizes {
		if pack > maxPackSize {
			maxPackSize = pack
		}
	}

	dp := make([]Tuple, orderQty+maxPackSize+1)
	for i := range dp {
		dp[i] = Tuple{orderQty + maxPackSize, []int{}}
	}
	dp[0] = Tuple{0, []int{}}

	for i := 0; i < orderQty+maxPackSize+1; i++ {
		for _, pack := range packSizes {
			if pack <= i {
				if 1+dp[i-pack].Packs < dp[i].Packs {
					newCombo := append([]int(nil), dp[i-pack].Combo...)
					newCombo = append(newCombo, pack)
					dp[i] = Tuple{1 + dp[i-pack].Packs, newCombo}
				}
			}
		}
	}

	// Filter to get valid combinations and then sort them
	maxPackSize = 0
	for _, pack := range packSizes {
		if pack > maxPackSize {
			maxPackSize = pack
		}
	}

	var validCombinations []Tuple
	for _, combo := range dp[orderQty : orderQty+maxPackSize+1] {
		if combo.Packs != orderQty+maxPackSize {
			validCombinations = append(validCombinations, combo)
		}
	}

	// Sort by the sum of packs
	sort.Slice(validCombinations, func(i, j int) bool {
		return tSlice.Sum(validCombinations[i].Combo) < tSlice.Sum(validCombinations[j].Combo)
	})

	sort.Sort(sort.Reverse(sort.IntSlice(validCombinations[0].Combo)))
	return validCombinations[0].Combo
}
