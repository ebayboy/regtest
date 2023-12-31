package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
)

var serviceConfig = &service.Config{Name: "serviceName",
	DisplayName: "service Display Name",
	Description: "service description",
}

/*
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
*/

func registryTest() {
	// 创建：指定路径的项
	// 路径：HKEY_CURRENT_USER\Software\Hello Go
	//Microsoft\Windows\CurrentVersion\Internet Settings
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings\Hello Go`, registry.ALL_ACCESS)

	defer key.Close()
	// 判断是否已经存在了
	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}

	// 写入：32位整形值
	key.SetDWordValue(`32位整形值`, uint32(123456))
	// 写入：64位整形值
	key.SetQWordValue(`64位整形值`, uint64(123456))
	// 写入：字符串
	key.SetStringValue(`字符串`, `hello`)
	// 写入：字符串数组
	key.SetStringsValue(`字符串数组`, []string{`hello`, `world`})
	// 写入：二进制
	key.SetBinaryValue(`二进制`, []byte{0x11, 0x22})

	// 读取：字符串
	s, _, _ := key.GetStringValue(`字符串`)
	println(s)

	// 读取：一个项下的所有子项
	keys, _ := key.ReadSubKeyNames(0)
	for _, key_subkey := range keys {
		// 输出所有子项的名字
		println(key_subkey)
	}

	// 创建：子项
	subkey, _, _ := registry.CreateKey(key, `子项`, registry.ALL_ACCESS)
	defer subkey.Close()

	// 删除：子项
	// 该键有子项，所以会删除失败
	// 没有子项，删除成功
	registry.DeleteKey(key, `子项`)
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
	registryTest()
}
