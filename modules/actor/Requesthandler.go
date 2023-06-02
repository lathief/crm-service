package actor

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/payload/response"
	"net/http"
	"strconv"
)

type actorRequestHandler struct {
	actorController ActorController
}
type ActorRequestHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetActorById(c *gin.Context)
	UpdateActor(c *gin.Context)
	DeleteActor(c *gin.Context)
}

func (ar *actorRequestHandler) Register(c *gin.Context) {
	actorReq := new(request.AuthActor)
	err := c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.Register(*actorReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) Login(c *gin.Context) {
	actorReq := new(request.AuthActor)
	err := c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.Login(*actorReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) GetActorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.GetActorById(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) UpdateActor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	actorReq := new(ActorDTO)
	err = c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.UpdateActor(*actorReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) DeleteActor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.DeleteActor(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
