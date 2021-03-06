package config

import "fmt"

var (
	MySQLSource = "root:password@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
	//MySQLSource = "root:password@tcp(172.22.19.39:13306)/fileserver?charset=utf8"
)

func UpdateDBHost(host string) {
	MySQLSource = fmt.Sprintf("root:password@tcp(%s)/fileserver?charset=utf8", host)
}
