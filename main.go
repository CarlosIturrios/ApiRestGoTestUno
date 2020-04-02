package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/", PaginaPrincipal)      // URL para realizar metodo GET
	r.POST("/", PostPaginaPrincipal) // URL para realizar el metodo post
	r.GET("/query", QueryString)     // recibe un querystring para filtrado u obtencion de datos
	// /query?nombre=carlos&edad=20
	r.GET("/path/:nombre/:edad", PathParams) // URL para mandar parametros en la misma
	r.Run()                                  // Escucha las peticiones echas al localhost:8080
}

func PaginaPrincipal(c *gin.Context) {
	c.JSON(200, gin.H{
		"Mensaje": "Mensaje enviado desde GET",
	})
}

func PostPaginaPrincipal(c *gin.Context) {
	body := c.Request.Body
	fmt.Println(body)
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"Mensaje": string(value),
	})
}

func QueryString(c *gin.Context) {
	nombre := c.Query("nombre") // nombre = Carlos
	edad := c.Query("edad")     // edad = 20

	c.JSON(200, gin.H{
		"nombre": nombre,
		"edad":   edad,
	})
}

func PathParams(c *gin.Context) {
	nombre := c.Param("nombre")
	edad := c.Param("edad")

	c.JSON(200, gin.H{
		"nombre": nombre,
		"edad":   edad,
	})
}
