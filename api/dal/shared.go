package dal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// Mutator is a wrapper to handle DGraph mutations
func Mutator(db *dgo.Dgraph, structToMutate interface{}) (uid string, err error) {
	// Handle validation, etc

	// Marshal to json for dgraph
	dbQuery, err := json.Marshal(structToMutate)
	if err != nil {
		return "", err
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

// SimpleQuery is a genericized function to wrap queries for a node by UID
func SimpleQuery(db *dgo.Dgraph, q string, match string) (resp []byte, err error) {
	// Set up transaction
	txn := db.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	// Make variables map
	variables := map[string]string{"$match": match}

	// Run the query
	out, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return resp, err
	}

	txn.Commit(context.Background())

	return out.GetJson(), err
}
