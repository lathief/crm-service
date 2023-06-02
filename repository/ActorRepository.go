package repository

import (
	"github.com/lathief/crm-service/entity"
	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

type ActorInterfaceRepository interface {
	CreateActor(actor entity.Actor) error
	GetActorById(id uint) (entity.Actor, error)
	GetActorByName(name string) (entity.Actor, error)
	UpdateActor(actor entity.Actor, id uint) error
	DeleteActor(id uint) error
	GetRole(name string) (entity.Role, error)
}

func ActorNewRepo(db *gorm.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (c *ActorRepository) CreateActor(actor entity.Actor) error {
	err := c.db.Model(&entity.Actor{}).Create(&actor).Error
	return err
}
func (c *ActorRepository) GetActorByName(name string) (entity.Actor, error) {
	var actor entity.Actor
	err := c.db.First(&actor, "username = ? ", name).Error
	return actor, err
}
func (c *ActorRepository) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	err := c.db.Preload("Role").First(&actor, "id = ? ", id).Error
	return actor, err
}
func (c *ActorRepository) UpdateActor(actor entity.Actor, id uint) error {
	err := c.db.Model(&entity.Actor{}).Where("id = ?", id).Updates(entity.Actor{
		Username: actor.Username, Password: actor.Password}).Error
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
