package main

import (
	"github.com/golobby/container/v3"

	pa "github.com/husamettinarabaci/go-pdftojpeg/core/application/presentation/adapter"
	as "github.com/husamettinarabaci/go-pdftojpeg/core/application/service"
	ds "github.com/husamettinarabaci/go-pdftojpeg/core/domain/service"
	ia "github.com/husamettinarabaci/go-pdftojpeg/pkg/infrastructure/adapter"
	cr "github.com/husamettinarabaci/go-pdftojpeg/pkg/presentation/controller/rest"
	tconfig "github.com/husamettinarabaci/go-pdftojpeg/tool/config"
)

var restConfig tconfig.RestConfig
var webConfig tconfig.WebConfig

func main() {
	restConfig.ReadConfig()
	webConfig.ReadConfig()
	var err error
	cont := container.New()

	//Domain PdfToJpeg Service
	err = cont.Singleton(func() ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure PdfToJpeg Converter Adapter
	err = cont.Singleton(func() ia.ConverterAdapter {
		return ia.NewConverterAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure PdfToJpeg Log Adapter
	err = cont.Singleton(func() ia.LogAdapter {
		return ia.NewLogAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Application PdfToJpeg Service
	err = cont.Singleton(func(s ds.Service, i ia.ConverterAdapter, l ia.LogAdapter) as.Service {
		return as.NewService(s, i, l)
	})
	if err != nil {
		panic(err)
	}

	//Application PdfToJpeg Query Adapter
	err = cont.Singleton(func(s as.Service) pa.QueryAdapter {
		return pa.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application PdfToJpeg Converter Adapter
	err = cont.Singleton(func(s as.Service) pa.ConverterAdapter {
		return pa.NewConverterAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var converterHandler pa.ConverterAdapter
	err = cont.Resolve(&converterHandler)
	if err != nil {
		panic(err)
	}

	cr.NewRestAPI(queryHandler, converterHandler).Serve(restConfig.Debug, restConfig.Restapi.Port)
}
