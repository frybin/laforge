package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/frybin/laforge/grpc-alpha/laforge_proto_agent"
	"google.golang.org/grpc/credentials"

	"os"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

)

const (
	port = ":50051"
	certFile = "server.crt"
	keyFile = "server.key"
)


type server struct {
	pb.UnimplementedLaforgeServer
}

/* TEST POSTGRES */

/*func CreatePGTable(){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	//figure out how to auto create table on start
	err = db.Exec(`CREATE TABLE LAFORGE;`)

	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully created new table")

}*/

func PingPG(){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully connected!")

	//create table after successful ping
	//CreatePGTable()
}

/* TEST MESSAGES */

//Ping Info
func (s *server) GetPing(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received: %v | ID: %v", in.GetName(), in.GetId())

	//postgres test
	PingPG()

	return &pb.PingReply{Name: "Hello " + in.GetName(), Id: in.GetId()}, nil
}

//HostTest Info
func (s *server) GetHostTest(ctx context.Context, in *pb.HostTestRequest) (*pb.HostTestReply, error) {
	log.Printf("Got Host: %v | ID: %v | IP: %s | OS: %s", in.GetName(), in.GetId(), in.GetIp(), in.GetOs())
	return &pb.HostTestReply{Name: in.GetName(), Id: in.GetId(), Ip: in.GetIp(), Os: in.GetOs()}, nil
}

/*  BASE LAFORGE */
// Fields Source: https://app.swaggerhub.com/apis/LaForge/LaforgeAPI/0.0.1-oas3#

//Competition Info
/*
func (s *server) GetCompetition(ctx context.Context, in *pb.CompetitionRequest) (*pb.CompetitionReply, error) {
	// set defaults if name or id is omitted due to oneof from client request
	name := "N/A"
	id := "000000"

	switch {
	case in.GetName() != "":
		name = in.GetName()
	case in.GetId() != "":
		id = in.GetId()
	}

	log.Printf("Client Send - Competition Name: %v | ID: %v", in.GetName(), in.GetId())
	log.Printf("Server Change - Competition Name: %v | ID: %v", name, id)

	//demo response
	comp := pb.CompetitionReply{Name: name, Id: id, Environments: []int32{1, 2, 3}, Users: []int32{11, 22, 33}, BuildConfigs: []int32{111, 222, 333}}
	return &comp, nil
}

// Environment Info
func (s *server) GetEnvironment(ctx context.Context, in *pb.EnvironmentRequest) (*pb.EnvironmentReply, error) {
	// set defaults if name or id is omitted due to oneof from client request
	name := "N/A"
	id := in.GetId()

	switch {
	case in.GetName() != "":
		name = in.GetName()
	}

	log.Printf("Client Send - Environment Name: %v | ID: %v", in.GetName(), in.GetId())
	
	//demo response
	env := pb.EnvironmentReply{Id: id, CompetitionId: 123456, OwnerId: 1111, Name: name, State: "Not Running", Attrs: []string{"local", "internal only"}, Networks: []int32{1, 2, 3}, Teams: []int32{11, 22, 33}}
	return &env, nil
}
*/


func main() {


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//secure server
	creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	s := grpc.NewServer(grpc.Creds(creds))

	//insecure server
	//s := grpc.NewServer()

	fmt.Println("Starting Laforge Server on port " + port)

	pb.RegisterLaforgeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}