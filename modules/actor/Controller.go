package actor

import (
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/payload/response"
	"net/http"
	"strconv"
)

type actorController struct {
	ActorUseCase UseCaseActor
}
type ActorController interface {
	CreateActor(request.AuthActor) (response.Response, error)
	GetActorById(actorId int) (response.Response, error)
	UpdateActor(Actor ActorDTO, actorId int) (response.Response, error)
	DeleteActor(actorId int) (response.Response, error)
}

func (ac *actorController) CreateActor(actor request.AuthActor) (response.Response, error) {
	err := ac.ActorUseCase.CreateActor(actor)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), 500), err
	}
	return response.HandleSuccessResponse(nil, "Create Actor Successfully", 201), err
}

func (ac *actorController) GetActorById(actorId int) (response.Response, error) {
	user, err := ac.ActorUseCase.GetActorById(actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusNotFound), err
	}
	return response.HandleSuccessResponse(user, "Success Get Actor By ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) UpdateActor(Actor ActorDTO, actorId int) (response.Response, error) {
	err := ac.ActorUseCase.UpdateActor(Actor, actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Update Actor ID: "+strconv.Itoa(actorId), 200), err
}

func (ac *actorController) DeleteActor(actorId int) (response.Response, error) {
	err := ac.ActorUseCase.DeleteActor(actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Delete Actor ID: "+strconv.Itoa(actorId), 200), err
}
