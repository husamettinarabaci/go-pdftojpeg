package presentation

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pp "github.com/husamettinarabaci/go-pdftojpeg/core/application/presentation/port"
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
	file, _ := c.FormFile("file")
	uid := uuid.New()
	c.SaveUploadedFile(file, "./tmp/"+uid.String()+".pdf")
	var converterRequestDto dto.ConverterRequest
	converterRequestDto.Item = uid.String()
	converterResponse, err := api.commandHandler.Convert(c, converterRequestDto.ToConverterRequestEntity())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	os.Remove("./tmp/" + uid.String() + ".pdf")
	c.JSON(http.StatusOK, dto.FromResponseObject(converterResponse.Response))
}
