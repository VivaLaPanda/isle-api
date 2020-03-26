package main

import (
	"flag"
	"log"

	"github.com/VivaLaPanda/isle-api/api/router"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// Various runtime flags
var dbAddress = flag.String("dbAddress", "127.0.0.1:9080", "HTTP address of the dgraph database")
var serverPort = flag.Int("serverPort", 9090, "Port to run the REST api on")

func main() {
	flag.Parse()

	dbClient := newClient(*dbAddress)

	LoadSchema(dbClient)
	router.ServeAPI(*serverPort, dbClient)
}

func newClient(dbAddress string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(dbAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
