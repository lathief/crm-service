package actor

import (
	"errors"
	"fmt"
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/repository"
)

type useCaseActor struct {
	ActorRepo    repository.ActorInterfaceRepository
	ApprovalRepo repository.ApprovalInterfaceRepository
}
type UseCaseActor interface {
	CreateActor(actor request.AuthActor) error
	GetActorById(id int) (ActorDTO, error)
	UpdateActor(Actor ActorDTO, id int) error
	DeleteActor(id int) error
	UpdateFlagActor(actor ActorDTO, id int) error

	SearchApproval() ([]ApprovalDTO, error)
	SearchApprovalByStatus(status string) ([]ApprovalDTO, error)
	GetApprovalById(id int) (ApprovalDTO, error)
	ChangeStatusApproval(id int, status string) error
}

func (uc *useCaseActor) CreateActor(actor request.AuthActor) error {
	get, _ := uc.ActorRepo.GetRole(entity.ROLE_ADMIN)
	ActorSave := entity.Actor{
		Username:   actor.Username,
		Password:   actor.Password,
		IsActive:   entity.False,
		IsVerified: entity.False,
		Role:       &get,
	}
	err := uc.ActorRepo.CreateActor(&ActorSave)
	if err != nil {
		return err
	}
	getSuperAdmin, err := uc.ActorRepo.GetActorById(uint(1))
	if err != nil {
		return err
	}
	fmt.Println(getSuperAdmin)
	newApproval := entity.Approval{
		Admin_id:      ActorSave.ID,
		Admin:         &ActorSave,
		Superadmin:    &getSuperAdmin,
		Superadmin_id: getSuperAdmin.ID,
		Status:        "pending",
	}
	err = uc.ApprovalRepo.CreateApproval(newApproval)
	if err != nil {
		return err
	}
	fmt.Println(newApproval)
	return nil
}
func (uc *useCaseActor) GetActorById(id int) (ActorDTO, error) {
	get, err := uc.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return ActorDTO{}, err
	}
	getActor := ActorDTO{
		Username:   get.Username,
		Password:   get.Password,
		IsVerified: string(get.IsVerified),
		IsActive:   string(get.IsActive),
		Role:       get.Role.Rolename,
	}
	return getActor, nil
}
func (uc *useCaseActor) UpdateActor(Actor ActorDTO, id int) error {
	_, err := uc.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	ActorUpdate := entity.Actor{
		Username: Actor.Username,
		Password: Actor.Password,
	}
	err = uc.ActorRepo.UpdateActor(ActorUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseActor) DeleteActor(id int) error {
	_, err := uc.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	err = uc.ActorRepo.DeleteActor(uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseActor) UpdateFlagActor(actor ActorDTO, id int) error {
	getActor, err := uc.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	getApproval, err := uc.ApprovalRepo.GetApprovalByActorId(getActor.ID)
	if err != nil {
		return err
	}
	if getApproval.Status != "active" {
		return errors.New("Approve dulu bang super admin")
	}
	ActorUpdate := entity.Actor{
		Username:   getActor.Username,
		Password:   getActor.Password,
		IsActive:   entity.BoolType(actor.IsActive),
		IsVerified: entity.BoolType(actor.IsVerified),
	}
	fmt.Println(getActor.Username)
	err = uc.ActorRepo.UpdateActor(ActorUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}

func (uc *useCaseActor) SearchApproval() ([]ApprovalDTO, error) {
	gets, err := uc.ApprovalRepo.SearchApproval()
	if err != nil {
		return nil, err
	}
	var appovalsDTO []ApprovalDTO

	for _, item := range gets {
		fmt.Println(item.Status)
		approvalDTO := ApprovalDTO{
			Admin: ActorDTO{
				Username:   item.Admin.Username,
				IsVerified: string(item.Admin.IsVerified),
				IsActive:   string(item.Admin.IsVerified),
			},
			Status: item.Status,
		}
		appovalsDTO = append(appovalsDTO, approvalDTO)
	}
	return appovalsDTO, nil
}
func (uc *useCaseActor) SearchApprovalByStatus(status string) ([]ApprovalDTO, error) {
	gets, err := uc.ApprovalRepo.SearchApprovalByStatus(status)
	if err != nil {
		return nil, err
	}
	var appovalsDTO []ApprovalDTO
	for _, item := range gets {
		approvalDTO := ApprovalDTO{
			Admin: ActorDTO{
				Username:   item.Admin.Username,
				IsVerified: string(item.Admin.IsVerified),
				IsActive:   string(item.Admin.IsVerified),
			},
			Status: item.Status,
		}
		appovalsDTO = append(appovalsDTO, approvalDTO)
	}
	return appovalsDTO, nil
}
func (uc *useCaseActor) GetApprovalById(id int) (ApprovalDTO, error) {
	get, err := uc.ApprovalRepo.GetApprovalById(uint(id))
	if err != nil {
		return ApprovalDTO{}, err
	}
	approvalDTO := ApprovalDTO{
		Admin: ActorDTO{
			Username:   get.Admin.Username,
			IsVerified: string(get.Admin.IsVerified),
			IsActive:   string(get.Admin.IsVerified),
		},
		Status: get.Status,
	}
	return approvalDTO, nil
}
func (uc *useCaseActor) ChangeStatusApproval(id int, status string) error {
	_, err := uc.ApprovalRepo.GetApprovalById(uint(id))
	if err != nil {
		return err
	}
	approvalUpdate := entity.Approval{
		Status: status,
	}
	err = uc.ApprovalRepo.UpdateApproval(approvalUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
