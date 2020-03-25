package dal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// NewContentNode creates a new ContentNode based on the provided struct
func NewContentNode(db *dgo.Dgraph, node models.NewContentNodeNode) (uid string, err error) {
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
		return "", err
	}

	txn.Commit(context.Background())

	// Put the UID on the node data we're going to return
	// It's just the first (only) UID in the response
	for _, value := range out.GetUids() {
		return value, nil
	}

	log.Fatal("Mutation didn't return a node")
	return
}

// ExpandContentNode is a function to get the data for a node and its children
func ExpandContentNode(db *dgo.Dgraph, uid string) (resp models.ContentNode, err error) {
	// Set up transaction
	txn := db.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	// Construct the query
	const q = `
	query ExpandContentNode($id: string) {
	  node(func: uid($id)) {
		  uid
		  ...PostBody
		  ~parent {
				uid
		  ...PostBody
			~parent {
						uid
			  ...PostBody
			}
		  }
		}
	  }
		  
	  fragment PostBody {
		  title
		  body
		  author {
			  uid
			  name
		  }
	}
  `

	// Make variables map
	variables := map[string]string{"$id": uid}

	// Run the query
	out, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		return resp, err
	}

	txn.Commit(context.Background())

	// Decode the response
	var decode struct {
		Node []models.ContentNode
	}
	if err := json.Unmarshal(out.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}

	// There's only ever one node
	return decode.Node[0], nil
}
