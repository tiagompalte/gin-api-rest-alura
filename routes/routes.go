package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagompalte/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ObterLista)
	r.GET("/alunos/:id", controllers.ObterPorId)
	r.POST("/alunos", controllers.Inserir)
	r.DELETE("/alunos/:id", controllers.Deletar)
	r.PUT("/alunos/:id", controllers.Atualizar)
	r.Run(":8000")
}
