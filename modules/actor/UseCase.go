package actor

import (
	"errors"
	"github.com/lathief/crm-service/config"
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/repository"
	"github.com/lathief/crm-service/utils/helper"
	"github.com/lathief/crm-service/utils/security"
	"strconv"
)

type useCaseActor struct {
	ActorRepo    repository.ActorInterfaceRepository
	ApprovalRepo repository.ApprovalInterfaceRepository
}
type UseCaseActor interface {
	Register(actor request.AuthActor) error
	Login(actor request.AuthActor) (string, error)
	GetActorById(id int) (ActorDTO, error)
	SearchActorByName(filter map[string]string) (*helper.Pagination, error)
	UpdateActor(Actor ActorDTO, id int) error
	DeleteActor(id int) error
	UpdateFlagActor(actor ActorDTO, id int) error
	SearchApproval() ([]ApprovalDTO, error)
	SearchApprovalByStatus(status string) ([]ApprovalDTO, error)
	GetApprovalById(id int) (ApprovalDTO, error)
	ChangeStatusApproval(id int, status string) error
}

func (au *useCaseActor) Register(actor request.AuthActor) error {
	existUsername, _ := au.ActorRepo.GetActorByName(actor.Username)
	if existUsername.Username != "" {
		return errors.New("Username already in use")
	}
	get, _ := au.ActorRepo.GetRole(entity.ROLE_ADMIN)
	ActorSave := entity.Actor{
		Username:   actor.Username,
		Password:   actor.Password,
		IsActive:   entity.False,
		IsVerified: entity.False,
		Role:       &get,
	}
	err := au.ActorRepo.CreateActor(&ActorSave)
	if err != nil {
		return err
	}
	getSuperAdmin, err := au.ActorRepo.GetActorByName(config.Config.SuperAccount.SuperName)
	if err != nil {
		return err
	}
	newApproval := entity.Approval{
		Admin_id:      ActorSave.ID,
		Admin:         &ActorSave,
		Superadmin:    &getSuperAdmin,
		Superadmin_id: getSuperAdmin.ID,
		Status:        "pending",
	}
	err = au.ApprovalRepo.CreateApproval(newApproval)
	if err != nil {
		return err
	}
	return nil
}
func (au *useCaseActor) Login(actor request.AuthActor) (string, error) {
	account, _ := au.ActorRepo.GetActorByName(actor.Username)
	if account.Username == "" {
		return "", errors.New("Account not found")
	}
	match := security.ComparePass([]byte(account.Password), []byte(actor.Password))
	if match == false {
		return "", errors.New("Password does not match")
	}
	token := security.GenerateToken(account.ID, account.Username)
	return token, nil
}
func (au *useCaseActor) GetActorById(id int) (ActorDTO, error) {
	get, err := au.ActorRepo.GetActorById(uint(id))
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
func (au *useCaseActor) SearchActorByName(filter map[string]string) (*helper.Pagination, error) {
	var result *helper.Pagination
	var totalRows int64
	var err error
	page, err := strconv.Atoi(filter["page"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	limit, err := strconv.Atoi(filter["limit"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	pagination := helper.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  "Id desc",
	}
	err = au.ActorRepo.CountRowActor(&totalRows)
	if err != nil {
		return &helper.Pagination{}, err
	}
	if filter["name"] != "" {
		result, err = au.ActorRepo.SearchActorByName(pagination, filter["name"], totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else {
		result, err = au.ActorRepo.GetAllActor(pagination, totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
	}
	return result, nil
}
func (au *useCaseActor) UpdateActor(Actor ActorDTO, id int) error {
	_, err := au.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	ActorUpdate := entity.Actor{
		Username: Actor.Username,
		Password: Actor.Password,
	}
	err = au.ActorRepo.UpdateActor(ActorUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (au *useCaseActor) DeleteActor(id int) error {
	_, err := au.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	err = au.ActorRepo.DeleteActor(uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (au *useCaseActor) UpdateFlagActor(actor ActorDTO, id int) error {
	getActor, err := au.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return err
	}
	getApproval, err := au.ApprovalRepo.GetApprovalByActorId(getActor.ID)
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
	err = au.ActorRepo.UpdateActor(ActorUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (au *useCaseActor) SearchApproval() ([]ApprovalDTO, error) {
	gets, err := au.ApprovalRepo.SearchApproval()
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
func (au *useCaseActor) SearchApprovalByStatus(status string) ([]ApprovalDTO, error) {
	gets, err := au.ApprovalRepo.SearchApprovalByStatus(status)
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
func (au *useCaseActor) GetApprovalById(id int) (ApprovalDTO, error) {
	get, err := au.ApprovalRepo.GetApprovalById(uint(id))
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
func (au *useCaseActor) ChangeStatusApproval(id int, status string) error {
	_, err := au.ApprovalRepo.GetApprovalById(uint(id))
	if err != nil {
		return err
	}
	approvalUpdate := entity.Approval{
		Status: status,
	}
	err = au.ApprovalRepo.UpdateApproval(approvalUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
