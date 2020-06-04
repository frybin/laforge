package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/frybin/laforge/grpc-alpha/laforge_proto"
)

const (
	address     = "localhost:50051"
	defaultName = "Laforge Agent 1"
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


/*  BASE LAFORGE */
// Fields Source: https://app.swaggerhub.com/apis/LaForge/LaforgeAPI/0.0.1-oas3#

// Request Competition - by string name, string id
func competition(c pb.LaforgeClient, name string) {
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

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	
	// competition
	comp_name := "Demo Comp"

	// END VARS

	ping(c, name)
	hostTest(c, name)
	competition(c, comp_name)

}
