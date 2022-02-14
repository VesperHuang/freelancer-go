package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/micro/go-plugins/registry/kubernetes"
)

func SignupHandler(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	name := json["name"]
	firstName := json["firstName"]
	lastName := json["lastName"]
	mobile := json["mobile"]
	email := json["email"]
	password := json["password"]

	// resp, err := userCli.Signup(context.TODO(), &userProto.ReqSignup{
	// 	Username: username,
	// 	Password: password,
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"name":      name,
		"firstName": firstName,
		"lastName":  lastName,
		"mobile":    mobile,
		"email":     email,
		"password":  password,
	})
}
