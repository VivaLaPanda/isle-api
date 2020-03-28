package dal

import (
	"log"
	"testing"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

var testParentNode models.NewContentNodeNode = models.NewContentNodeNode{
	ContentNode: models.ContentNode{
		Title:      "post-a",
		Body:       "body-a",
		Created:    time.Now(),
		Edited:     time.Now(),
		Score:      0,
		Sentiment:  1,
		DgraphType: []string{"ContentNode"},
	},
}

var testChildNode models.NewContentNodeNode = models.NewContentNodeNode{
	ContentNode: models.ContentNode{
		Title:      "post-b",
		Body:       "body-b",
		Created:    time.Now(),
		Edited:     time.Now(),
		Score:      0,
		Sentiment:  1,
		DgraphType: []string{"ContentNode"},
	},
}

func createNodes() (parentUID, childUID string, err error) {
	db := newClient()

	models.LoadSchema(db)

	parentUID, err = NewContentNode(db, testParentNode)

	if err != nil {
		return
	}

	testChildNode.ParentUID = &models.Edge{UID: parentUID}

	childUID, err = NewContentNode(db, testChildNode)

	return
}

func TestExpandContentNode(t *testing.T) {
	db := newClient()

	parentUID, childUID, err := createNodes()
	if err != nil {
		t.Errorf("Creating nodes during TestExpandContentNode failed, err: %s", err)
		return
	}

	actualParent, err := ExpandContentNode(db, parentUID)
	actualChild, err := ExpandContentNode(db, childUID)

	if err != nil {
		t.Errorf("TestExpandContentNode failed, err: %s", err)
		return
	}

	if actualParent.Title != testParentNode.Title {
		t.Errorf("TestExpandContentNode failed, testNode.Title != expected.Title: %s != %s", actualParent.Title, testParentNode.Title)
	}

	if actualParent.Children[0].UID != actualChild.UID {
		t.Errorf("TestExpandContentNode failed, actualParent.Children[0].UID != actualChild.UID: %s != %s", actualParent.Children[0].UID, actualChild.UID)
	}
}
