package main

import (
	"flag"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"middle/app/user/rpc/internal/config"
	"middle/app/user/rpc/internal/svc"
	"middle/pkg/database"
)

var configFile = flag.String("f", "../etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化数据库
	err := error(nil)
	if database.DB, err = gormc.ConnectMysql(c.Mysql); err != nil {
		logx.Error(err)
		panic(err)
	}
	database.Connect(database.DB)

	// 注册命令
	svc.SetupCmd()
}
