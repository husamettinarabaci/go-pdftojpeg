package application

import (
	as "github.com/husamettinarabaci/go-pdftojpeg/core/application/service"
)

type QueryAdapter struct {
	service as.Service
}

func NewQueryAdapter(s as.Service) QueryAdapter {
	return QueryAdapter{
		service: s,
	}
}
