package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// NewContent will add a new ContentNode to the databse.
func NewContent(db *dgo.Dgraph) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parsing the body into a Go struct
		var node models.NewContentNode
		err := json.NewDecoder(r.Body).Decode(&node)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// Handle validation, etc
		// Making sure creation time is current
		node.Created = time.Now()

		// Marshal to json for dgraph
		dbQuery, err := json.Marshal(node)
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
			http.Error(w, err.Error(), 500)
			return
		}

		txn.Commit(context.Background())

		// Put the UID on the node data we're going to return
		// It's just the first (only) UID in the response
		for _, value := range out.GetUids() {
			node.UID = value
		}

		response := &ResponseMsg{
			Status:  http.StatusCreated,
			Message: "Created new content node",
			Result:  node,
		}

		// Marshall the response
		nodeJSON, err := json.Marshal(node)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(response.Status)
		w.Write(nodeJSON)
	})
}
