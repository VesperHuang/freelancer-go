package client

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro"
	"github.com/mitchellh/mapstructure"

	"freelancer-go/service/dbproxy/orm"
	dbProto "freelancer-go/service/dbproxy/proto"
)

type UserMeta struct {
	Name       string
	FirstName  string
	MiddleName string
	LastName   string
	Mobile     string
	Email      string
	Password   string
}

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var (
	dbCli dbProto.DBProxyService
)

func Init(service micro.Service) {
	dbCli = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

func TableFileToFileMeta(tfile orm.TableFile) FileMeta {
	return FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
}

// execAction : 向dbproxy请求执行action
func execAction(funcName string, paramJson []byte) (*dbProto.RespExec, error) {
	return dbCli.ExecuteAction(context.TODO(), &dbProto.ReqExec{
		Action: []*dbProto.SingleAction{
			&dbProto.SingleAction{
				Name:   funcName,
				Params: paramJson,
			},
		},
	})
}

// parseBody : 转换rpc返回的结果
func parseBody(resp *dbProto.RespExec) *orm.ExecResult {
	if resp == nil || resp.Data == nil {
		return nil
	}
	resList := []orm.ExecResult{}
	_ = json.Unmarshal(resp.Data, &resList)
	// TODO:
	if len(resList) > 0 {
		return &resList[0]
	}
	return nil
}

func ToTableUser(src interface{}) orm.TableUser {
	user := orm.TableUser{}
	mapstructure.Decode(src, &user)
	return user
}

func ToTableFile(src interface{}) orm.TableFile {
	file := orm.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableFiles(src interface{}) []orm.TableFile {
	file := []orm.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableUserFile(src interface{}) orm.TableUserFile {
	ufile := orm.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func ToTableUserFiles(src interface{}) []orm.TableUserFile {
	ufile := []orm.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func GetFileMeta(filehash string) (*orm.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{filehash})
	res, err := execAction("/file/GetFileMeta", uInfo)
	return parseBody(res), err
}

func GetFileMetaList(limitCnt int) (*orm.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{limitCnt})
	res, err := execAction("/file/GetFileMetaList", uInfo)
	return parseBody(res), err
}

// OnFileUploadFinished : 新增/更新文件元信息到mysql中
func OnFileUploadFinished(fmeta FileMeta) (*orm.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location})
	res, err := execAction("/file/OnFileUploadFinished", uInfo)
	return parseBody(res), err
}

func UpdateFileLocation(filehash, location string) (*orm.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{filehash, location})
	res, err := execAction("/file/UpdateFileLocation", uInfo)
	return parseBody(res), err
}

func UserSignup(user UserMeta, encPasswd string) (*orm.ExecResult, error) {

	uInfo, _ := json.Marshal([]interface{}{user.Name, user.FirstName, user.MiddleName, user.LastName, encPasswd, user.Mobile, user.Email})
	res, err := execAction("/user/UserSignup", uInfo)
	return parseBody(res), err
}
