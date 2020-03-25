package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VivaLaPanda/isle-api/api/dal"
	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo"
	"github.com/gorilla/mux"
)

// NewContentNode will add a new ContentNode to the databse.
func NewContentNode(db *dgo.Dgraph) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parsing the body into a Go struct
		var node models.NewContentNodeNode
		err := json.NewDecoder(r.Body).Decode(&node)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		newNodeUID, err := dal.NewContentNode(db, node)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		node.UID = newNodeUID

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

// ExpandContentNode will add a new ContentNode to the databse.
func ExpandContentNode(db *dgo.Dgraph) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the UID of the node we're querying
		uid := mux.Vars(r)["uid"]

		contentNode, err := dal.ExpandContentNode(db, uid)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		// Structure our response
		response := &ResponseMsg{
			Status:  http.StatusFound,
			Message: "Fetched node successfully",
			Result:  contentNode,
		}

		// Marshall the response
		respJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(response.Status)
		w.Write(respJSON)
	})
}
