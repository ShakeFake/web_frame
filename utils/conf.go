package utils

import (
	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

var (
	HttpPort string

	EtcdAddr string
	EtcdPort string
	EtcdUser string
	EtcdPass string
)

func Config() {
	conf, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		panic(err)
	}

	HttpPort = conf.MustValue("[Server]", "http_port", "8091")

	EtcdAddr = conf.MustValue("[Etcd]", "etcd_addr", "")
	EtcdPort = conf.MustValue("[Etcd]", "etcd_port", "")
	EtcdUser = conf.MustValue("[Etcd]", "etcd_user", "")
	EtcdPass = conf.MustValue("[Etcd]", "etcd_pass", "")

	log.Infof("HttpPort is:%v", HttpPort)

	log.Infof("EtcdAddr is:%v", EtcdAddr)
	log.Infof("EtcdPort is:%v", EtcdPort)
	log.Infof("EtcdUser is:%v", EtcdUser)
	log.Infof("EtcdPass is:%v", EtcdPass)

}
