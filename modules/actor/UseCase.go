package actor

import (
	"errors"
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/payload/request"
	"github.com/lathief/crm-service/repository"
	"github.com/lathief/crm-service/utils/security"
)

type useCaseActor struct {
	ActorRepo repository.ActorInterfaceRepository
}
type UseCaseActor interface {
	Login(actor request.AuthActor) error
	CreateActor(actor request.AuthActor) error
	GetActorById(id int) (ActorDTO, error)
	UpdateActor(Actor ActorDTO, id int) error
	DeleteActor(id int) error
}

func (uc *useCaseActor) CreateActor(actor request.AuthActor) error {
	existUsername, _ := uc.ActorRepo.GetActorByName(actor.Username)
	if existUsername.Username != "" {
		return errors.New("Username already in use")
	}
	get, _ := uc.ActorRepo.GetRole(entity.ROLE_ADMIN)
	ActorSave := entity.Actor{
		Username:   actor.Username,
		Password:   actor.Password,
		IsActive:   entity.False,
		IsVerified: entity.False,
		Role:       &get,
	}
	err := uc.ActorRepo.CreateActor(ActorSave)
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseActor) Login(actor request.AuthActor) error {
	account, _ := uc.ActorRepo.GetActorByName(actor.Username)
	if account.Username == "" {
		return errors.New("Account not found")
	}
	match := security.ComparePass([]byte(account.Password), []byte(actor.Password))
	if match == false {
		return errors.New("Password does not match")
	}
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
