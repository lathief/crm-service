package repository

import (
	"github.com/lathief/crm-service/entity"
	"gorm.io/gorm"
)

type ApprovalRepository struct {
	db *gorm.DB
}
type ApprovalInterfaceRepository interface {
	SearchApproval() ([]entity.Approval, error)
	SearchApprovalByStatus(status string) ([]entity.Approval, error)
	GetApprovalById(id uint) (entity.Approval, error)
	GetApprovalByActorId(id uint) (entity.Approval, error)
	CreateApproval(approval entity.Approval) error
	UpdateApproval(approval entity.Approval, id uint) error
}

func ApprovalNewRepo(db *gorm.DB) *ApprovalRepository {
	return &ApprovalRepository{db: db}
}

func (a *ApprovalRepository) SearchApproval() ([]entity.Approval, error) {
	var approval []entity.Approval
	err := a.db.Preload("Admin").Find(&approval).Error
	return approval, err
}
func (a *ApprovalRepository) SearchApprovalByStatus(status string) ([]entity.Approval, error) {
	var approval []entity.Approval
	err := a.db.Preload("Admin").Where("status = ?", status).Find(&approval).Error
	return approval, err
}
func (a *ApprovalRepository) GetApprovalById(id uint) (entity.Approval, error) {
	var approval entity.Approval
	err := a.db.Preload("Admin").First(&approval, "id = ? ", id).Error
	return approval, err
}
func (a *ApprovalRepository) CreateApproval(approval entity.Approval) error {
	err := a.db.Model(&entity.Approval{}).Create(&approval).Error
	return err
}
func (a *ApprovalRepository) UpdateApproval(approval entity.Approval, id uint) error {
	err := a.db.Model(&entity.Approval{}).Where("id = ?", id).Updates(entity.Approval{
		Status: approval.Status}).Error
	return err
}
func (a *ApprovalRepository) GetApprovalByActorId(id uint) (entity.Approval, error) {
	var approval entity.Approval
	err := a.db.Preload("Admin").Where("admin_id = ?", id).Find(&approval).Error
	return approval, err
}
