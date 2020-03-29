package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
)

// NewContentNode creates a new ContentNode based on the provided struct
func NewContentNode(db *dgo.Dgraph, node models.NewContentNodeNode) (uid string, err error) {
	// Handle validation, etc
	// Making sure creation time is current
	node.Created = time.Now()

	return Mutator(db, node)
}

// ExpandContentNode is a function to get the data for a node and its children
func ExpandContentNode(db *dgo.Dgraph, uid string) (resp models.ContentNode, err error) {
	// Set up transaction
	txn := db.NewReadOnlyTxn()
	defer txn.Discard(context.Background())

	// Construct the query
	const q = `
	query ExpandContentNode($match: string) {
	  node(func: uid($match)) @filter(type(ContentNode)) {
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
		  created
		  edited
		  imageUri
		  sentiment
		  score
		  tags
		  dgraph.type
		  author {
			  uid
			  name
			  aviImgUri
		  }
	}
  `

	jsonResp, err := SimpleQuery(db, q, uid)

	// Decode the response
	var decode struct {
		Node []models.ContentNode
	}
	if err := json.Unmarshal(jsonResp, &decode); err != nil {
		log.Fatal(err)
	}

	if len(decode.Node) != 1 {
		return resp, fmt.Errorf("Query returned %d results, not 1 as expected. Did you query the wrong endpoint?", len(decode.Node))
	}

	// There's only ever one node
	return decode.Node[0], nil
}
