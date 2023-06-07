package actor

import (
	"context"
	"github.com/lathief/crm-service/payload"
	"github.com/lathief/crm-service/utils/helper"
	"strconv"
)

type actorController struct {
	ActorUseCase UseCaseActor
}
type ActorController interface {
	Register(payload.AuthActor) (payload.Response, error)
	Login(payload.AuthActor) (payload.Response, error)
	GetActorById(actorId int) (payload.Response, error)
	SearchActorByName(c context.Context, filter map[string]string) (payload.Response, error)
	UpdateActor(updateActor payload.UpdateActor, actorId int) (payload.Response, error)
	DeleteActor(actorId int) (payload.Response, error)
	SearchApproval(status string) (payload.Response, error)
	GetApprovalById(approvalId int) (payload.Response, error)
	UpdateFlagActor(updateFlag payload.UpdateFlagActor, actorId int) (payload.Response, error)
	ChangeStatusApproval(approvalId int, status payload.ApprovalStatus) (payload.Response, error)
}

func (ac *actorController) Register(actor payload.AuthActor) (payload.Response, error) {
	err := ac.ActorUseCase.Register(actor)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Create Actor Successfully", 201), err
}
func (ac *actorController) Login(actor payload.AuthActor) (payload.Response, error) {
	token, err := ac.ActorUseCase.Login(actor)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(payload.ResponseLogin{Token: token}, "Login Successfully", 200), err
}
func (ac *actorController) SearchActorByName(c context.Context, filter map[string]string) (payload.Response, error) {
	actors, err := ac.ActorUseCase.SearchActorByName(c, filter)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(actors, "Success Get Actors", 200), err
}
func (ac *actorController) GetActorById(actorId int) (payload.Response, error) {
	user, err := ac.ActorUseCase.GetActorById(actorId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(user, "Success Get Actor By ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) UpdateActor(updateActor payload.UpdateActor, actorId int) (payload.Response, error) {
	err := ac.ActorUseCase.UpdateActor(updateActor, actorId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Update Actor ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) DeleteActor(actorId int) (payload.Response, error) {
	err := ac.ActorUseCase.DeleteActor(actorId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Delete Actor ID: "+strconv.Itoa(actorId), 200), err
}
func (ac *actorController) SearchApproval(status string) (payload.Response, error) {
	if status == "" {
		res, err := ac.ActorUseCase.SearchApproval()
		if err != nil {
			return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
		}
		return payload.HandleSuccessResponse(res, "Success Get All Approval Request", 200), err
	} else {
		res, err := ac.ActorUseCase.SearchApprovalByStatus(status)
		if err != nil {
			return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
		}
		return payload.HandleSuccessResponse(res, "Success Get Approval Request with status "+status, 200), err
	}
}
func (ac *actorController) GetApprovalById(approvalId int) (payload.Response, error) {
	approval, err := ac.ActorUseCase.GetApprovalById(approvalId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(approval, "success get approval request by id: "+strconv.Itoa(approvalId), 200), err
}
func (ac *actorController) UpdateFlagActor(updateFlag payload.UpdateFlagActor, actorId int) (payload.Response, error) {
	err := ac.ActorUseCase.UpdateFlagActor(updateFlag, actorId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Update Flag Actor ID: "+strconv.Itoa(actorId), 200), nil
}
func (ac *actorController) ChangeStatusApproval(approvalId int, status payload.ApprovalStatus) (payload.Response, error) {
	err := ac.ActorUseCase.ChangeStatusApproval(approvalId, status)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Change Status Approval ID: "+strconv.Itoa(approvalId), 200), nil
}
