package handler

import (
	"context"
	"log"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"

	cmn "freelancer-go/common"
	userProto "freelancer-go/service/account/proto"
)

var (
	userCli userProto.UserService
)

func init() {
	service := micro.NewService(
		micro.Flags(cmn.CustomFlags...),
	)
	service.Init()

	cli := service.Client()
	userCli = userProto.NewUserService("go.micro.service.user", cli)
}

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

	reqSignup := &userProto.ReqSignup{
		Name:       name.(string),
		FirstName:  firstName.(string),
		MiddleName: middleName.(string),
		LastName:   lastName.(string),
		Mobile:     mobile.(string),
		Email:      email.(string),
		Password:   password.(string),
	}
	fmt.Println("reqSignup => ",reqSignup)

	resp, err := userCli.Signup(context.TODO(), reqSignup)

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}
