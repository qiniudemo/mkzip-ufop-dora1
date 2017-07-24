package main

import (
	"fmt"
	"github.com/qiniu/api.v6/conf"
	"github.com/qiniu/log"
	"os"
	"ufop"
	"ufop/mkzip"
)

const (
	VERSION = "1.4"
)

func help() {
	fmt.Printf("Usage: qufop <UfopConfig>\r\n\r\nVERSION: %s\r\n", VERSION)
}

func setQiniuHosts() {
	conf.RS_HOST = "http://rs.qiniu.com"
}

func main() {
	log.SetOutput(os.Stdout)
	setQiniuHosts()

	args := os.Args
	argc := len(args)

	var configFilePath string

	switch argc {
	case 2:
		configFilePath = args[1]
	default:
		help()
		return
	}

	//load config
	ufopConf := &ufop.UfopConfig{}
	confErr := ufopConf.LoadFromFile(configFilePath)
	if confErr != nil {
		log.Error("load config file error,", confErr)
		return
	}

	ufopServ := ufop.NewServer(ufopConf)

	if err := ufopServ.RegisterJobHandler("mkzip.conf", &mkzip.Mkzipper{}); err != nil {
		log.Error(err)
	}

	//listen
	ufopServ.Listen()
}
