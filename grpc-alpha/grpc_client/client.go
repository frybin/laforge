package main

import (
	"context"
	"flag"
	"time"

	pb "github.com/frybin/laforge/grpc-alpha/laforge_proto_agent"
	"github.com/kardianos/service"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address          = "localhost:50051"
	defaultName      = "Laforge Agent 1"
	certFile         = "server.crt"
	heartbeatSeconds = 5
)

var (
	logger service.Logger
)

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

// SendHeartBeat Example
func SendHeartBeat(c pb.LaforgeClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := &pb.HeartbeatRequest{Id: 12345}
	hostInfo, hostErr := host.Info()
	if hostErr == nil {
		(*request).Hostname = hostInfo.Hostname
		(*request).Uptime = hostInfo.Uptime
		(*request).Boottime = hostInfo.BootTime
		(*request).Numprocs = hostInfo.Procs
		(*request).Os = hostInfo.OS
		(*request).Hostid = hostInfo.HostID
	}
	mem, memErr := mem.VirtualMemory()
	if memErr == nil {
		(*request).Totalmem = mem.Total
		(*request).Freemem = mem.Free
		(*request).Usedmem = mem.Used
	}
	load, loadErr := load.Avg()
	if loadErr == nil {
		(*request).Load1 = load.Load1
		(*request).Load5 = load.Load5
		(*request).Load15 = load.Load15
	}
	r, err := c.GetHeartBeat(ctx, request)
	if err != nil {
		logger.Error("Error: %v", err)
	} else {
		logger.Infof("Response Message: %s", r.GetStatus())
	}

}

func genSendHeartBeat(p *program, c pb.LaforgeClient) {
	ticker := time.NewTicker(time.Duration(heartbeatSeconds) * time.Second)
	for {
		select {
		case <-ticker.C:
			SendHeartBeat(c)
		case <-p.exit:
			ticker.Stop()
		}
	}
}

func (p *program) run() error {
	logger.Infof("I'm running %v.", service.Platform())
	creds, credErr := credentials.NewClientTLSFromFile(certFile, "")
	if credErr != nil {
		logger.Errorf("Cred Error: %v", credErr)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))

	if err != nil {
		logger.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLaforgeClient(conn)

	// START VARS
	genSendHeartBeat(p, c)
	return nil
}
func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

// Service setup.
//   Define service config.
//   Create the service.
//   Setup the logger.
//   Handle service controls (optional).
//   Run the service.
func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "laforge-agent",
		DisplayName: "Laforge Client Agent",
		Description: "Laforge Client Agent",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target"},
		Option: options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Error(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		logger.Error(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				logger.Info(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			logger.Infof("Valid actions: %q\n", service.ControlAction)
			logger.Error(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
