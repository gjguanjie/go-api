package register

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
	"os"
	"strconv"
)

func Register(consulHost string, consulPort string, svcHost string, svcPort string,logger log.Logger) (registar sd.Registrar) {
	var client consul.Client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + consulPort
	consulClient,err := api.NewClient(consulConfig)
	if err != nil {
		logger.Log("create consul client error")
		os.Exit(1)
	}
	client = consul.NewClient(consulClient)
	healthCheck := api.AgentServiceCheck{
		HTTP: "http://" + svcHost + ":" + svcPort + "/health",
		Interval: "10s",
		Timeout: "1s",
		Notes: "Consul check service health status",
	}
	port,_ := strconv.Atoi(svcPort)
	reg := api.AgentServiceRegistration{
		ID: "arithmetic" + uuid.New(),
		Name: "arithmetic",
		Address: svcHost,
		Port: port,
		Tags: []string{"arithmetic"},
		Check:   &healthCheck,
	}
	registar  = consul.NewRegistrar(client, &reg,logger)


}