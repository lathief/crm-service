package actor

import (
	"github.com/gin-gonic/gin"
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
					ActorRepo:    repository.ActorNewRepo(db),
					ApprovalRepo: repository.ApprovalNewRepo(db),
				},
			},
		},
	}
}

func (ar *ActorRoute) Handle(router *gin.Engine) {
	actorPath := "/actor"
	actor := router.Group(actorPath)
	actor.POST("", ar.ActorHandler.CreateActor)
	actor.GET("/:id", ar.ActorHandler.GetActorById)
	actor.PUT("/:id", ar.ActorHandler.UpdateActor)
	actor.PUT("/flag/:id", ar.ActorHandler.UpdateFlagActor)
	actor.DELETE("/:id", ar.ActorHandler.DeleteActor)

	approvePath := "/approval"
	approve := router.Group(approvePath)
	approve.GET("", ar.ActorHandler.SearchApproval)
	approve.GET("/:id", ar.ActorHandler.GetApprovalById)
	approve.PUT("/:id", ar.ActorHandler.ChangeStatusApproval)
	//approve.PUT("/:id", ar.ActorHandler.UpdateActor)
	//approve.DELETE("/:id", ar.ActorHandler.DeleteActor)
}
