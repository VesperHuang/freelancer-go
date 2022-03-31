package route

import (
	"freelancer-go/assets"
	"freelancer-go/service/apigw/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	assetfs "github.com/moxiaomomo/go-bindata-assetfs"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{
		Asset:     assets.Asset,
		AssetDir:  assets.AssetDir,
		AssetInfo: assets.AssetInfo,
		Prefix:    root,
	}
	return &binaryFileSystem{
		fs,
	}
}

func Router() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	//將靜態文件 打包到 bin 文件
	//router.Use(static.Serve("/static/", BinaryFileSystem("static")))

	router.POST("/user/signup", handler.SignupHandler)
	// router.POST("/user/signin", handler.DoSigninHandler)

	return router
}
