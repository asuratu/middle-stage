package svc

import (
	"fmt"
	"os"

	"middle/app/user/rpc/database"
	"middle/pkg/cmd"
	"middle/pkg/cmd/make"
	"middle/pkg/console"

	"github.com/spf13/cobra"
)

func SetupCmd(c *ServiceContext) {
	var rootCmd = &cobra.Command{
		Use:   c.Config.Name,
		Short: "A simple forum project",
		Long:  `this is root cmd`,
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		make.CmdMake,
		database.CmdMigrate,
		database.CmdDBSeed,
	)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
