package uti

import (
	"flag"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/common/log"
)

var ConsulClient *api.Client

var ServerId string
var ServerName string
var ServerPort int

func ServerReg(name string,port int){
	ServerName=name
	ServerPort=port
}

func init(){

	config:=api.DefaultConfig()
	config.Address="127.0.0.1:8500"
	client,_:=api.NewClient(config)
	ConsulClient=client
	ServerId="UserServers"+uuid.New().String()
}

func ConsulReg()  {
	name:=flag.String("name","","服务名")
	port:=flag.Int("port",0,"端口名")
	flag.Parse()

	ServerReg(*name,*port)    // 命令行获取服务名和端口
	// 修改Consul的配置
	reg:=api.AgentServiceRegistration{}
	reg.ID=ServerId
	reg.Name=ServerName
	reg.Address="127.0.0.1"       // 网关
	reg.Port=ServerPort
	reg.Tags=[]string{"primary"}

	check:= api.AgentServiceCheck{}
	check.Interval="5s"
	check.HTTP="http://127.0.0.1:8080/health"
	reg.Check=&check

	ConsulClient.Agent().ServiceRegister(&reg)
}

// 删除consul
func UnRegister()  {
	err:=ConsulClient.Agent().ServiceDeregister("101")
	if err!=nil{
		log.Error(err)
	}
}
