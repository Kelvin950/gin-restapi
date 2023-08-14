package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
 
type Address struct{

	City string `json:"city"`
	Town string `json:"town"` 
	HouseNumber int `json:"house_no"`
}

type Person  struct{
  
	Name string  `json:"name"` 
	Age int 	`json:"age"`
	Address Address  `json:"address"`
 
}

func fetchAll(ctx *gin.Context){
   	ctx.JSON(http.StatusOK , gin.H{
		"hell":"2",
	})

	
}
func fetch(ctx *gin.Context){
	ctx.JSON(http.StatusOK , gin.H{
		"hell":"2",
	})
}

func delete(ctx *gin.Context){

	ctx.JSON(http.StatusOK , gin.H{
		"hell":"2",
	})
}

func create(ctx *gin.Context){
    
  fmt.Printf("%v" , ctx.Request.Body)
 newFile , err := os.OpenFile("data.json", os.O_CREATE |os.O_TRUNC |os.O_RDWR, 0064)
  
     if err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
		} )
	 }
    
 

	buffreader := bufio.NewReader(newFile)
fileinfo , err  := newFile.Stat() 
     if err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
		} )
	 }
var dataByte = make([]byte ,fileinfo.Size() )
 buffreader.Read(dataByte)	

 var personRead *[]Person

 if err =  json.Unmarshal(dataByte , personRead) ; err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
			"message":err.Error() ,
		} ) 
	 }

	 var person *Person
  if err =  ctx.BindJSON(&person) ; err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
			"message":err.Error() ,
		} ) 

  }
   
    
  *personRead =  append(*personRead , *person)  

 buffWriter :=  bufio.NewWriter(newFile) 

   
 
 dataWrite,err := json.Marshal(personRead) 

	if err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
			"message":err.Error() ,
		} )  
	}

    buffWritten , err  := buffWriter.Write(dataWrite)
    if err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed" , 
			"message":err.Error() ,
		} )  
	}

	
	buffWriter.Flush()
	

	ctx.JSON(http.StatusCreated , gin.H{
		"success":true ,
		 "message":gin.H{
         "data":buffWritten,
		 } ,
		 
	})
}



func main(){

 r:=  gin.Default()
 


 r.GET("/", fetchAll)
 r.GET("/:id", fetch)
 r.POST("/" ,create)
r.DELETE("/:id",delete)
 r.Run(":8000")
}