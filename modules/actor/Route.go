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
					ActorRepo:    repository.ActorNewRepo(db),
					ApprovalRepo: repository.ApprovalNewRepo(db),
				},
			},
		},
	}
}

func (ar *ActorRoute) Handle(router *gin.Engine) {
	router.POST("/login", ar.ActorHandler.Login)
	router.POST("/register", ar.ActorHandler.Register)

	adminPath := "/admin"
	adminRG := router.Group(adminPath)
	{
		adminRG.Use(middleware.Authentication())
		adminRG.GET("/:id", ar.ActorHandler.GetActorById)
		adminRG.GET("/search", ar.ActorHandler.Search)
		adminRG.PUT("/:id", middleware.AdminAuthorization(), ar.ActorHandler.UpdateActor)
		adminRG.DELETE("/:id", middleware.SuperAdminAuthorization(), ar.ActorHandler.DeleteActor)
	}
	approvePath := "/approval"
	approveRG := router.Group(approvePath)
	{
		approveRG.Use(middleware.Authentication())
		approveRG.GET("/search", middleware.SuperAdminAuthorization(), ar.ActorHandler.SearchApproval)
		approveRG.GET("/:id", middleware.SuperAdminAuthorization(), ar.ActorHandler.GetApprovalById)
		approveRG.PUT("/:id", middleware.SuperAdminAuthorization(), ar.ActorHandler.ChangeStatusApproval)
	}
}
