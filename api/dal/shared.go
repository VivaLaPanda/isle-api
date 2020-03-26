package dal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// NewUser creates a new User based on the provided struct
func Mutator(db *dgo.Dgraph, structToMutate interface{}) (uid string, err error) {
	// Handle validation, etc

	// Marshal to json for dgraph
	dbQuery, err := json.Marshal(structToMutate)
	if err != nil {
		log.Fatal(err)
	}

	// Set up transaction
	txn := db.NewTxn()
	defer txn.Discard(context.Background())

	// Run the query
	out, err := txn.Mutate(context.Background(), &api.Mutation{
		SetJson: dbQuery,
	})
	if err != nil {
		return "", err
	}

	txn.Commit(context.Background())

	// Get the UID to return
	for _, value := range out.GetUids() {
		return value, nil
	}

	log.Fatal("Mutation didn't return a result")
	return
}
