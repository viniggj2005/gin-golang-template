package routes

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viniggj2005/api-rest-go/dtos"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Println("Início da request:", c.Request.Method, c.Request.URL.Path)

		c.Next()
		latencia := time.Since(t)
		status := c.Writer.Status()
		fmt.Println("Fim da request:", status, " | tempo:", latencia)
	}
}

func HanddleRequests() {
	// io := NewSocketIOServer() //inicializa o socket
	melody := newWebsocketServer()
	r := gin.Default() //inicializa o gin
	r.Use(Logger())    //inicializa o logger
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Range"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
	})) //configurações do cors
	r.GET("/ws", func(c *gin.Context) {
		melody.HandleRequest(c.Writer, c.Request)
	})
	handleWebsocketconnection(melody)
	// r.POST("/users/:id", controllers.CreateUser)//exemplo de rota com parametro
	//r.GET("/users", controllers.FindAllUser)//exemplo de rota sem parametro

	// r.Static("/videos/hls", "./uploads/hls")//exemplo de rota para arquivos estaticos
	// r.GET("/socket.io/", gin.WrapH(io.HttpHandler())) //rota para a chamado do socket

	// r.DELETE("/uma-pessoa/:id", controllers.DeletaUmaPessoa)
	r.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //rota para o ginswagger

	//exemplo de validação , traduzindo as mensagens de validação
	r.POST("/validate-test", func(c *gin.Context) {
		var req dtos.UserDto

		if err := c.ShouldBindJSON(&req); err != nil {
			var errs validator.ValidationErrors
			if errors.As(err, &errs) {
				c.JSON(400, gin.H{"errors": errs.Translate(Trans)})
				return
			}
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, req)
	})

	r.Run("0.0.0.0:3000") //inicializa o servidor
}
