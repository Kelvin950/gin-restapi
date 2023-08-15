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
 

	var persons []Person 
 
	file ,err :=  os.OpenFile("data.json" , os.O_RDONLY  , 0644) 

	if  err == nil{
        
	 buffReader := bufio.NewReader(file) 
    
	 fileinfo ,err  :=  file.Stat() 
 
	   if err == nil {
	  databyte := make([]byte , fileinfo.Size())
  buffReader.Read(databyte)
        
      err =  json.Unmarshal(databyte ,&persons)
	  if err == nil {

		ctx.JSON(http.StatusOK , gin.H{
			"success":true , 
			"data": persons ,
		})
	  }else {
			ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed1d1" , 
			"message":err.Error() ,
		} )
	  }

	   }else{
			ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed1d1" , 
			"message":err.Error() ,
		} )
	   }
	


	}else{
			ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed1d1" , 
			"message":err.Error() ,
		} )
	}

 

	
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
    


 newFile , err := os.OpenFile("data.json", os.O_CREATE|os.O_RDWR, 0644)
  
     if err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failedcr" , 
			"message":err.Error() ,
		} )
	 }
    
 
defer newFile.Close()
	buffreader := bufio.NewReader(newFile)

     if err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed1d1" , 
			"message":err.Error() ,
		} )
	 }

	var sizeofByt int64 

	if fileinfo , err := newFile.Stat() ; fileinfo.Size() > 0 {
		sizeofByt = fileinfo.Size()
	}else if err != nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed12" , 
			"message":err.Error() ,
		} ) 
	}else  {
		sizeofByt = 0
	}
	 dataByte :=make([]byte ,sizeofByt )
 buffreader.Read(dataByte)	
fmt.Printf("qwq%s\n" , dataByte)
 var personRead []Person

 if err =  json.Unmarshal(dataByte , &personRead) ; err !=nil{

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed1234" , 
			"message":err.Error() ,
		} ) 
	 }

	 var person *Person
  if err =  ctx.BindJSON(&person) ; err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failedxfd" , 
			"message":err.Error() ,
		} ) 

  }
   fmt.Println(person)
    
  personRead =  append(personRead , *person)  


  

os.Remove(newFile.Name()) ;

newFile ,err=  os.Create("data.json")
 if err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failedxfd" , 
			"message":err.Error() ,
		} ) 

  }
 buffWriter :=  bufio.NewWriter(newFile) 

   
 
 dataWrite,err := json.MarshalIndent(&personRead, " ", " ") 

	if err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failedwqdw" , 
			"message":err.Error() ,
		} )  
	}

    buffWritten , err  := buffWriter.Write(dataWrite)
    if err !=nil{
		ctx.JSON(http.StatusInternalServerError , gin.H{
			"success":"failed11" , 
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