// ---------------------------------------
// 2021/09/23 - Create By Kee
// ---------------------------------------
// URL 资源获取命令控制
// --------- Request Options ---------
// 	-X --method		请求方法(GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH...)
// 	-d --data		<DATA>请求数据
// 	-F --form		<KEY=VALUE>以表单形式提交数据（？待定）
// 	-H --header		<HEADER>设置请求头
// 	-A --user-agent	<STRING>设置请求代理头; 默认为: Go-http-client/1.1, 参数为AUTO时随机生成
//  -p --parallel   <NUMBER>设置并发请求数量
// 	-------- 连接和协议选项 --------
// 	--http10		使用 HTTP/1.0 协议
// 	--http11		使用 HTTP/1.1 协议
// 	--http2			使用 HTTP/2 协议
//  --insecure		忽略证书验证
//  --timeout		<SECONDS>设置请求超时时间
//  --max-time		<SECONDS>设置请求最大时间
//	-------- 代理认证选项 --------
//	--proxy			<PROXY_URL>设置代理地址
//	--proxy-user	<USER:PASSWORD>设置代理认证信息
//	--proxy-header	<HEADER>设置代理请求头
//	-u --user		<USER:PASSWORD>设置认证信息
// 	-------- Output Options --------
// 	-I --head 		仅获取响应头
//	-o --output		<FILE>将响应内容输出到文件
//	-v --verbose	输出详细信息

package cmd

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/urfave/cli/v2"
)

type Cmd struct {
	app *cli.App
}

var (
	version  = "0.0.1"
	cmdOpts  = COpts{}
	cmdFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "methods",
			Aliases: []string{"X2s"},
			Value:   "GET",
			Usage:   "Request Method <GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH...>",
			//Destination: &cmdOpts.Method,
		},
	}
)

func cmdAction(ctx *cli.Context) error {
	// fmt.Println("The Args Count:", ctx.NArg())
	// fmt.Println("args-0:", ctx.Args().Get(0))

	// fetch cmdOpts options
	keys := getTypeKeys(cmdOpts)
	ref := reflect.ValueOf(&cmdOpts).Elem()
	for _, key := range keys {
		val := ref.FieldByName(key)
		fmt.Println(key, ":", val)
	}

	//fmt.Println("args-Methon:", cmdOpts.Method, cmdOpts)

	// // get keys
	// fmt.Println("get keys: ", getTypeDesc(cmdOpts))
	// desc := getTypeDesc(cmdOpts)
	// fmt.Println("covert: ", typeDescToCmdFlags(cmdOpts, desc))
	return nil
}

func NewCli() *Cmd {
	app := &cli.App{
		Name:                 "go-curl",
		Usage:                "A simple command line tool for sending HTTP requests",
		Version:              version,
		EnableBashCompletion: true,
		Action:               cmdAction,
	}
	return &Cmd{app}
}

func (cmd *Cmd) Run() {
	//cmd.app.Flags = cmdFlags
	// cmd.app.Commands = cmds

	desc := getTypeDesc(&cmdOpts)
	flags := typeDescToCmdFlags(cmdOpts, desc)
	cmd.app.Flags = flags

	if err := cmd.app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
