package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pp "github.com/husamettinarabaci/go-pdftojpeg/core/application/presentation/port"
	mo "github.com/husamettinarabaci/go-pdftojpeg/core/domain/model/object"
	dto "github.com/husamettinarabaci/go-pdftojpeg/pkg/presentation/dto"
)

type RestAPI struct {
	engine         *gin.Engine
	commandHandler pp.CommandPort
	queryHandler   pp.QueryPort
}

func NewRestAPI(qh pp.QueryPort, ch pp.CommandPort) *RestAPI {
	api := &RestAPI{
		commandHandler: ch,
		queryHandler:   qh,
	}
	api.engine = gin.New()
	api.engine.POST("/api/converter", api.Convert)
	return api
}

func (api *RestAPI) Serve(debug bool, port string) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	api.engine.Run(":" + port)
}

func (api *RestAPI) Convert(c *gin.Context) {
	var converterRequestDto dto.ConverterRequest
	if err := c.ShouldBindJSON(&converterRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	converterRequestDto.FillPackSizes()
	if err := converterRequestDto.IsValid(); err == nil {
		converterResponse, err := api.commandHandler.Convert(c, converterRequestDto.ToConverterRequestEntity())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.FromResponseObject(converterResponse.Response))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": mo.ErrInvalidInput.Error()})
	}
}
