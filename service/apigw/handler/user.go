package handler

import (
	"context"
	"log"
	"net/http"

	userProto "freelancer-go/service/account/proto"

	"github.com/gin-gonic/gin"
)

var (
	userCli userProto.UserService
)

func SignupHandler(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	name := json["name"]
	firstName := json["firstName"]
	middleName := json["middleName"]
	lastName := json["lastName"]
	mobile := json["mobile"]
	email := json["email"]
	password := json["password"]

	resp, err := userCli.Signup(context.TODO(), &userProto.ReqSignup{
		Name:       name.(string),
		FirstName:  firstName.(string),
		MiddleName: middleName.(string),
		LastName:   lastName.(string),
		Mobile:     mobile.(string),
		Email:      email.(string),
		Password:   password.(string),
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"name":       name,
	// 	"firstName":  firstName,
	// 	"middleName": middleName,
	// 	"lastName":   lastName,
	// 	"mobile":     mobile,
	// 	"email":      email,
	// 	"password":   password,
	// })

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}
