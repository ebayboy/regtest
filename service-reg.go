package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

var serviceConfig = &service.Config{Name: "serviceName",
	DisplayName: "service Display Name",
	Description: "service description",
}

func regAdd() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SAM\XIAO`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, err := k.GetStringValue("x")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", s)
}

func cmdsAdd() {
	log.Info("cmdsAdd ...")

	u := fmt.Sprintf("http://127.0.0.1/proxy.js")
	log.Info("u:", u)

	err := exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Internet Settings`, "/v", "AutoConfigURL", "/t", "REG_SZ", "/d", u, "/f").Run()
	if err != nil {
		log.Errorf("set AutoConfigURL error, %s\n", err.Error())
		return
	}

	log.Info("cmdsAdd done!")
}
func main() {

	file := "c:\\" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	// 设置日志级别
	log.SetLevel(log.DebugLevel)

	// 将文件设置为log输出的文件
	log.SetOutput(logFile)

	// 构建服务对象
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	cmd := os.Args[1]

	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	}
	if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}

	// install, uninstall, start, stop 的另一种实现方式
	// err = service.Control(s, os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	cmdsAdd()
}
