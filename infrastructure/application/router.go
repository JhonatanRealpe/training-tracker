package application

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setUpRoutes(dependencies *appDependencies) {
	router := gin.Default()

	router.POST("/training", dependencies.trainingHandler.SaveTraining)
	router.GET("/team", dependencies.teamHandler.GetTeam)
	router.POST("/players", dependencies.playersHandler.SavePlayers)
	router.GET("/players", dependencies.playersHandler.GetPlayers)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	defer dependencies.db.Close()
	router.Run(":8080")
}
