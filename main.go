package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
 

func fetchAll(ctx *gin.Context){
   

	
}
func fetch(ctx *gin.Context){

}

func delete(ctx *gin.Context){

}

func create(ctx *gin.Context){
    
 
 newFile , err := os.OpenFile("data.json", os.O_CREATE |os.O_TRUNC |os.O_RDWR, 0064)


}



func main(){

 r:=  gin.Default()
 


 r.GET("/")
 r.GET("/:id")
 r.POST("/")
r.DELETE("/:id")
 r.Run(":8000")
}