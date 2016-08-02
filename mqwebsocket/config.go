package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	l4g "github.com/alecthomas/log4go"
)

type RedisConfig struct {
	Host string `json:"Host,omitempty"`
	Port int    `json:"Port,omitempty"`
	Addr string `json:"addr,omitempty"`
}

type HttpsConfig struct {
	Port     int    `json:"Port,omitempty"`
	CertFile string `json:"CertFile,omitempty"`
	KeyFile  string `json:"KeyFile,omitempty"`
}

type Config struct {
	LogConfigFile       string      `json:"logconfig,omitempty"`           //l4g配置文件
	ListenIp            string      `json:"listenip,omitempty"`            //需要绑定的ip
	HttpPort            int         `json:"http_port,omitempty"`           //http端口
	Redis               RedisConfig `json:"redis,omitempty"`               //redis地址
	DBDns               string      `json:"dns,omitempty"`                 //数据库连接地址
	WebNotifyChannel    string      `json:"web_notify_channel,omitempty"`  //redis通知pub/sub的通道
	BasePath            string      `json:"basepath,omitempty"`            //url基础路径
	BackWebsocketAddr   string      `json:"back_websocket_addr,omitempty"` //url基础路径
	BackWebsocketOrigin string      `json:"back_websocket_origin",omitempty`
	Daemon              bool        `json:"Daemon,omitempty"` //
}

func (this *Config) Init(configpath string) error {

	this.HttpPort = 8001
	this.ListenIp = "0.0.0.0"
	this.LogConfigFile = "./etc/log4go.xml"
	if len(configpath) <= 0 {
		return errors.New("config file is empty.")
	}
	//读取配置文件内容
	configbuf, oerr := ioutil.ReadFile(configpath)
	if oerr != nil {
		fmt.Println("config init, read config failed, error:%s.", oerr.Error())
		return oerr
	}

	//解析文件内容
	if err := json.Unmarshal(configbuf, this); err != nil {
		fmt.Println("config init, parse config failed, error:%s.", err.Error())
		return err
	}

	l4g.Info("load config success, config:%v", this)
	return this.checkConfig()
}

func (this *Config) checkConfig() error {

	return nil
}
