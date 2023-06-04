package repository

import (
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/utils/helper"
	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

type ActorInterfaceRepository interface {
	GetAllActor(pagination helper.Pagination, totalRows int64) (*helper.Pagination, error)
	GetActorById(id uint) (entity.Actor, error)
	GetActorByName(name string) (entity.Actor, error)
	CountRowActor(totalRows *int64) error
	SearchActorByName(pagination helper.Pagination, name string, totalRows int64) (*helper.Pagination, error)
	CreateActor(actor *entity.Actor) error
	UpdateActor(actor entity.Actor, id uint) error
	DeleteActor(id uint) error
	GetRole(name string) (entity.Role, error)
}

func ActorNewRepo(db *gorm.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (c *ActorRepository) CountRowActor(totalRows *int64) error {
	err := c.db.Model(&entity.Actor{}).Count(totalRows).Error
	return err
}
func (c *ActorRepository) GetAllActor(pagination helper.Pagination, totalRows int64) (*helper.Pagination, error) {
	var actors []*entity.Actor
	err := c.db.Scopes(helper.Paginate(actors, &pagination, totalRows)).Find(&actors).Error
	pagination.Rows = actors
	return &pagination, err
}
func (c *ActorRepository) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	err := c.db.Preload("Role").First(&actor, "id = ? ", id).Error
	return actor, err
}
func (c *ActorRepository) GetActorByName(name string) (entity.Actor, error) {
	var actor entity.Actor
	err := c.db.First(&actor, "username = ? ", name).Error
	return actor, err
}
func (c *ActorRepository) SearchActorByName(pagination helper.Pagination, name string, totalRows int64) (*helper.Pagination, error) {
	var actor []*entity.Actor
	err := c.db.Scopes(helper.Paginate(actor, &pagination, totalRows)).Where("username LIKE ?", "%"+name+"%").Find(&actor).Error
	pagination.Rows = actor
	return &pagination, err
}
func (c *ActorRepository) CreateActor(actor *entity.Actor) error {
	err := c.db.Model(&entity.Actor{}).Create(&actor).Error
	return err
}
func (c *ActorRepository) UpdateActor(actor entity.Actor, id uint) error {
	err := c.db.Model(&entity.Actor{}).Where("id = ?", id).Updates(entity.Actor{
		Username: actor.Username, Password: actor.Password, IsVerified: actor.IsVerified, IsActive: actor.IsActive}).Error
	return err
}
func (c *ActorRepository) DeleteActor(id uint) error {
	err := c.db.First(&entity.Actor{}).Where("id = ?", id).Delete(&entity.Actor{}).Error
	return err
}
func (c *ActorRepository) GetRole(name string) (entity.Role, error) {
	var role entity.Role
	err := c.db.First(&role, "rolename = ? ", name).Error
	return role, err
}
