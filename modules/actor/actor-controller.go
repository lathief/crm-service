package actor

import (
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/payload/response"
	"github.com/lathief/crm-service/utils/helper"
	"strconv"
)

type actorController struct {
	ActorUseCase UseCaseActor
}
type ActorController interface {
	Register(request.AuthActor) (response.Response, error)
	Login(request.AuthActor) (response.Response, error)
	GetActorById(actorId int) (response.Response, error)
	SearchActorByName(filter map[string]string) (response.Response, error)
	UpdateActor(updateActor request.UpdateActor, actorId int) (response.Response, error)
	DeleteActor(actorId int) (response.Response, error)
	SearchApproval(status string) (response.Response, error)
	GetApprovalById(approvalId int) (response.Response, error)
	UpdateFlagActor(Actor ActorDTO, actorId int) (response.Response, error)
	ChangeStatusApproval(approvalId int, status request.ApprovalStatus) (response.Response, error)
}

func (ac *actorController) Register(actor request.AuthActor) (response.Response, error) {
	err := ac.ActorUseCase.Register(actor)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(nil, "Create Actor Successfully", 201), err
}
func (ac *actorController) Login(actor request.AuthActor) (response.Response, error) {
	token, err := ac.ActorUseCase.Login(actor)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(response.ResponseLogin{Token: token}, "Login Successfully", 200), err
}
func (ac *actorController) SearchActorByName(filter map[string]string) (response.Response, error) {
	actors, err := ac.ActorUseCase.SearchActorByName(filter)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(actors, "Success Get Actors", 200), err
}
func (ac *actorController) GetActorById(actorId int) (response.Response, error) {
	user, err := ac.ActorUseCase.GetActorById(actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(user, "Success Get Actor By ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) UpdateActor(updateActor request.UpdateActor, actorId int) (response.Response, error) {
	err := ac.ActorUseCase.UpdateActor(updateActor, actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(nil, "Success Update Actor ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) DeleteActor(actorId int) (response.Response, error) {
	err := ac.ActorUseCase.DeleteActor(actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(nil, "Success Delete Actor ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) SearchApproval(status string) (response.Response, error) {
	if status == "" {
		res, err := ac.ActorUseCase.SearchApproval()
		if err != nil {
			return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
		}
		return response.HandleSuccessResponse(res, "Success Get All Approval Request", 200), err
	} else {
		res, err := ac.ActorUseCase.SearchApprovalByStatus(status)
		if err != nil {
			return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
		}
		return response.HandleSuccessResponse(res, "Success Get Approval Request with status "+status, 200), err
	}
}
func (ac *actorController) GetApprovalById(approvalId int) (response.Response, error) {
	approval, err := ac.ActorUseCase.GetApprovalById(approvalId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(approval, "success get approval request by id: "+strconv.Itoa(approvalId), 200), err
}
func (ac *actorController) UpdateFlagActor(Actor ActorDTO, actorId int) (response.Response, error) {
	err := ac.ActorUseCase.UpdateFlagActor(Actor, actorId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(nil, "Success Update Flag Actor ID: "+strconv.Itoa(actorId), 200), nil
}
func (ac *actorController) ChangeStatusApproval(approvalId int, status request.ApprovalStatus) (response.Response, error) {
	err := ac.ActorUseCase.ChangeStatusApproval(approvalId, status)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return response.HandleSuccessResponse(nil, "Success Change Status Approval ID: "+strconv.Itoa(approvalId), 200), nil
}
