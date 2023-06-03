package actor

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/middleware"
	"github.com/lathief/crm-service/repository"
	"gorm.io/gorm"
)

type ActorRoute struct {
	ActorHandler ActorRequestHandler
}

func NewRouter(db *gorm.DB) ActorRoute {
	return ActorRoute{
		ActorHandler: &actorRequestHandler{
			actorController: &actorController{
				ActorUseCase: &useCaseActor{
					ActorRepo: repository.ActorNewRepo(db),
				},
			},
		},
	}
}

func (ar *ActorRoute) Handle(router *gin.Engine) {
	basePath := "/actor"
	customer := router.Group(basePath)
	customer.POST("/register", ar.ActorHandler.Register)
	customer.POST("/login", ar.ActorHandler.Login)
	customer.Use(middleware.Authentication())
	customer.GET("/:id", middleware.Authorization(), ar.ActorHandler.GetActorById)
	customer.PUT("/:id", ar.ActorHandler.UpdateActor)
	customer.DELETE("/:id", ar.ActorHandler.DeleteActor)
}
