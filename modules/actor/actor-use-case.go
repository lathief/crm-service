package actor

import (
	"fmt"
	"github.com/lathief/crm-service/config"
	"github.com/lathief/crm-service/constant"
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
	UpdateActor(updateActor request.UpdateActor, id int) error
	DeleteActor(id int) error
	UpdateFlagActor(actor ActorDTO, id int) error
	SearchApproval() ([]ApprovalDTO, error)
	SearchApprovalByStatus(status string) ([]ApprovalDTO, error)
	GetApprovalById(id int) (ApprovalDTO, error)
	ChangeStatusApproval(id int, status request.ApprovalStatus) error
}

func (au *useCaseActor) Register(actor request.AuthActor) error {
	existUsername, _ := au.ActorRepo.GetActorByName(actor.Username)
	if existUsername.Username != "" {
		return constant.ErrAdminUsernameExists
	}
	get, err := au.ActorRepo.GetRole(entity.ROLE_ADMIN)
	if err != nil {
		return constant.ErrRoleNotFound
	}
	ActorSave := entity.Actor{
		Username:   actor.Username,
		Password:   actor.Password,
		IsActive:   constant.False,
		IsVerified: constant.False,
		Role:       &get,
	}
	err = au.ActorRepo.CreateActor(&ActorSave)
	if err != nil {
		return err
	}
	getSuperAdmin, err := au.ActorRepo.GetActorByName(config.Config.SuperAccount.SuperName)
	if err != nil {
		return constant.ErrSuperAdminNotFound
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
		return "", constant.ErrAdminNotFound
	}
	match := security.ComparePass([]byte(account.Password), []byte(actor.Password))
	if match == false {
		return "", constant.ErrAdminPasswordNotMatch
	}
	if account.IsActive != constant.True {
		return "", constant.ErrAdminAccountNotActive
	}
	token := security.GenerateToken(account.ID, account.Username)
	return token, nil
}
func (au *useCaseActor) GetActorById(id int) (ActorDTO, error) {
	get, err := au.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return ActorDTO{}, constant.ErrAdminNotFound
	}
	getActor := ActorDTO{
		Username:   get.Username,
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
	err = au.ActorRepo.CountRowActor(&totalRows)
	if err != nil {
		return &helper.Pagination{}, err
	}
	pagination := helper.Pagination{
		Limit:     limit,
		Page:      page,
		Sort:      fmt.Sprintf("%s %s", filter["sortby"], filter["orderby"]),
		TotalRows: totalRows,
	}
	if filter["name"] != "" {
		result, err = au.ActorRepo.SearchActorByName(pagination, filter["name"])
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else {
		result, err = au.ActorRepo.GetAllActor(pagination)
		if err != nil {
			return &helper.Pagination{}, err
		}
	}
	var admins []ActorDTO
	data := result.Rows.([]*entity.Actor)
	for _, item := range data {
		var admin = ActorDTO{
			Role:       item.Role.Rolename,
			Username:   item.Username,
			IsActive:   string(item.IsActive),
			IsVerified: string(item.IsVerified),
		}
		admins = append(admins, admin)
	}
	result.Rows = admins
	return result, nil
}
func (au *useCaseActor) UpdateActor(updateActor request.UpdateActor, id int) error {
	_, err := au.ActorRepo.GetActorById(uint(id))
	if err != nil {
		return constant.ErrAdminNotFound
	}
	ActorUpdate := entity.Actor{
		Username: updateActor.Username,
		Password: security.HashPass(updateActor.Password),
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
		return constant.ErrAdminNotFound
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
		return constant.ErrAdminNotFound
	}
	getApproval, err := au.ApprovalRepo.GetApprovalByActorId(getActor.ID)
	if err != nil {
		return constant.ErrApprovalNotFound
	}
	if getApproval.Status != "approve" {
		return constant.ErrAdminNotApprove
	}
	ActorUpdate := entity.Actor{
		IsActive:   constant.BoolType(actor.IsActive),
		IsVerified: constant.BoolType(actor.IsVerified),
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
			ID: item.ID,
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
			ID: item.ID,
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
		return ApprovalDTO{}, constant.ErrApprovalNotFound
	}
	approvalDTO := ApprovalDTO{
		ID: get.ID,
		Admin: ActorDTO{
			Username:   get.Admin.Username,
			IsVerified: string(get.Admin.IsVerified),
			IsActive:   string(get.Admin.IsVerified),
		},
		Status: get.Status,
	}
	return approvalDTO, nil
}
func (au *useCaseActor) ChangeStatusApproval(id int, status request.ApprovalStatus) error {
	fmt.Println(status)
	_, err := au.ApprovalRepo.GetApprovalById(uint(id))
	if err != nil {
		return constant.ErrApprovalNotFound
	}
	approvalUpdate := entity.Approval{
		Status: status.Status,
	}
	err = au.ApprovalRepo.UpdateApproval(approvalUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
