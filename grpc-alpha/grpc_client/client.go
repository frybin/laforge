package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/frybin/laforge/grpc-alpha/laforge_proto_agent"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "Laforge Agent 1"
	certFile    = "server.crt"
)

/* TEST MESSAGES */

func ping(c pb.LaforgeClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPing(ctx, &pb.PingRequest{Name: name, Id: 111111})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s | ID: %v", r.GetName(), r.GetId())

}

func hostTest(c pb.LaforgeClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetHostTest(ctx, &pb.HostTestRequest{Name: name, Id: 123124, Ip: "1.1.1.1", Os: "Ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Host Info: %s, ID: %v, IP: %s, OS: %s", r.GetName(), r.GetId(), r.GetIp(), r.GetOs())
}

// SendHeartBeat Example
func SendHeartBeat(c pb.LaforgeClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	hostInfo, _ := host.Info()
	mem, _ := mem.VirtualMemory()
	load, _ := load.Avg()
	request := &pb.HeartbeatRequest{Id: 12345, Hostname: hostInfo.Hostname, Uptime: hostInfo.Uptime, Boottime: hostInfo.BootTime, Numprocs: hostInfo.Procs, Os: hostInfo.OS, Hostid: hostInfo.HostID, Load1: load.Load1, Load5: load.Load5, Load15: load.Load15, Totalmem: mem.Total, Freemem: mem.Free, Usedmem: mem.Used}
	r, err := c.GetHeartBeat(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response Message: %s", r.GetStatus())

}

/*  BASE LAFORGE */
// Fields Source: https://app.swaggerhub.com/apis/LaForge/LaforgeAPI/0.0.1-oas3#

// Request Competition - by string name, string id
/*func competition(c pb.LaforgeClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// COMP_NAME/ID are interchangeable due to "oneof" proto definition
	comp_id := &pb.CompetitionRequest_Id{Id: "1234"}
	r, err := c.GetCompetition(ctx, &pb.CompetitionRequest{Comp: comp_id})

	//comp_name := &pb.CompetitionRequest_Name{Name: name}
	//r, err := c.GetCompetition(ctx, &pb.CompetitionRequest{Comp: comp_name})

	if err != nil {
			log.Fatalf("could not greet: %v", err)
	}
	//print server demo response
	log.Printf("Competition Name: %v | ID: %v | Env: %v | Users: %v | Build Config: %v", r.GetName(), r.GetId(), r.GetEnvironments(), r.GetUsers(), r.GetBuildConfigs())
}

// Request Environment - by string name, string state, int32 id, int32 competition_id, int32 owner_id
func environment(c pb.LaforgeClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ENV_NAME/ID are interchangeable due to "oneof" proto definition
	env_id := &pb.EnvironmentRequest_Id{Id: 1234}
	r, err := c.GetEnvironment(ctx, &pb.EnvironmentRequest{Env: env_id})

	//env_name := &pb.EnvironmentRequest_Name{Name: name}
	//r, err := c.GetEnvironment(ctx, &pb.EnvironmentRequest{Env: env_name})

	if err != nil {
			log.Fatalf("could not greet: %v", err)
	}
	//print server demo response
	log.Printf("Environment Name: %v | ID: %v | Comp ID: %v | Owner ID: %v | State: %s | Atts: %v | Networks: %v | Teams: %v", r.GetName(), r.GetId(), r.GetCompetitionId(), r.GetOwnerId(), r.GetState(), r.GetAttrs(), r.GetNetworks(), r.GetTeams())
}
*/

func main() {
	// Set up a connection to the server.
	//secure connection
	creds, credErr := credentials.NewClientTLSFromFile(certFile, "")
	if credErr != nil {
		log.Fatalf("Cred Error: %v", credErr)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))

	//insecure connection
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLaforgeClient(conn)

	// START VARS
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//comp_name := "Demo Comp"
	//env_name := "Test Environment"

	// END VARS

	ping(c, name)
	hostTest(c, name)
	SendHeartBeat(c, name)
	//competition(c, comp_name)
	//environment(c, env_name)

}
