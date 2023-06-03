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
	CreateActor(c *gin.Context)
	GetActorById(c *gin.Context)
	UpdateActor(c *gin.Context)
	UpdateFlagActor(c *gin.Context)
	DeleteActor(c *gin.Context)
	SearchApproval(c *gin.Context)
	GetApprovalById(c *gin.Context)
	ChangeStatusApproval(c *gin.Context)
}

func (ar *actorRequestHandler) CreateActor(c *gin.Context) {
	actorReq := new(request.AuthActor)
	err := c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.CreateActor(*actorReq)
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

func (ar *actorRequestHandler) SearchApproval(c *gin.Context) {
	status := c.Query("status")
	res, err := ar.actorController.SearchApproval(status)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) GetApprovalById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.GetApprovalById(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) UpdateFlagActor(c *gin.Context) {
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
	res, err := ar.actorController.UpdateFlagActor(*actorReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) ChangeStatusApproval(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	str := c.Query("status")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.ChangeStatusApproval(id, str)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
