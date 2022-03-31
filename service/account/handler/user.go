package handler

import (
	"context"
	"fmt"

	"freelancer-go/common"
	cfg "freelancer-go/config"
	proto "freelancer-go/service/account/proto"
	dbCli "freelancer-go/service/dbproxy/client"
	"freelancer-go/util"
)

//to instantiate user service handler interface object
type User struct{}

func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) (error) {
	fmt.Println("user.go => userServiceHandler req =>", req)
	user := dbCli.UserMeta{
		Name:       req.Name,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Mobile:     req.Mobile,
		Email:      req.Email,
		Password:   req.Password,
	}

	fmt.Println("dbCli.UserMeta =>", user)

	//validate
	if len(user.Name) < 3 || len(user.Password) < 5 {
		res.Code = common.StatusParamInvalid
		res.Message = "register invalid"
		return nil
	}	

	encPassword := util.Sha1([]byte(req.Password + cfg.PasswordSalt))
	dbResp, err := dbCli.UserSignup(user, encPassword)
	if err == nil && dbResp.Suc {
		res.Code = common.StatusOK
		res.Message = "register success"
	} else {
		res.Code = common.StatusRegisterFailed
		res.Message = "register failed"
		fmt.Println("err:",err)
	}
	return nil
}

func (u *User) Signin(ctx context.Context, req *proto.ReqSignin, res *proto.RespSignin) error {

	return nil
}

// UserInfo ： 查询用户信息
func (u *User) UserInfo(ctx context.Context, req *proto.ReqUserInfo, res *proto.RespUserInfo) error {

	return nil
}
