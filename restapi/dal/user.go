package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/VivaLaPanda/isle-api/models"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

func NewUser(c *dgo.Dgraph, inviteCode string, user models.User) (*models.User, error) {
	txn := c.NewTxn()
	defer txn.Discard(context.Background())

	// Check the invite code
	invVariables := map[string]string{"$code": inviteCode}
	const q = `
    query InviteCheck($code: string){
      matchingInvite(func: eq(invite.code, $code)) {
        uid
        invite.code
        invite.created
      }
    }
  `
	resp, err := txn.QueryWithVars(context.Background(), q, invVariables)
	if err != nil {
		return nil, fmt.Errorf("Failed to check invite code: %s\nerr:%s", inviteCode, err)
	}

	// Unmarshal the invite
	type Root struct {
		matchingInvite []models.Invite `json:"matchingInvite"`
	}
	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal invite code: \nerr:%s\njson:%s", err, resp.Json)
	}

	// Check if the invite existed
	if len(r.matchingInvite) != 1 {
		// TODO: Treat this error differently since it's the user's fault
		return nil, fmt.Errorf("Invite code does not exist: %s", inviteCode)
	}

	// Mutation to create the user
	// setting defaults on the user
	user.InvitedBy = r.matchingInvite[0].CreatedBy.UID
	user.Joined = time.Now()
	user.Reputation = &float64(100.0)
	user.Spent = 0

	userToCreate, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal user obj into json. Err: \n%s", err)
	}
	mutUserRes, err := txn.Mutate(context.Background(), &api.Mutation{SetJson: userToCreate})
	if err != nil {
		return nil, fmt.Errorf("Failed to add new user to DB. Err: \n%s", err)
	}

	// Remove the invite code that was used
	rmInv := `
    {
      delete {
        <%s> * * .
      }
    }
  `
	rmInvJson := []byte{fmt.Sprintf(rmInv, r.matchingInvite[0].UID)}
	_, err = txn.Mutate(context.Background(), &api.Mutation{SetJson: rmInvJson})
	if err != nil {
		return nil, fmt.Errorf("Failed to remove used invite code. Err: \n%s", err)
	}

	newUser, err := getUserByUid(txn, mutUserRes.Uids)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch new user data. Err: \n%s", err)
	}

	// Commit
	err = txn.Commit(context.Background())

	return newUser, nil
}

func getUserByUid(txn *dgo.Txn, uid string) (*models.User, error) {
	variables := map[string]string{"$uid": uid}
	const q = `
    query FetchUser($code: string){
      user(func: uid($uid)) {
        uid
        user
        user.name
        user.email
        user.aviImgUri
        user.reputation
        user.spent
        user.role
        user.invitedBy
        user.joined
      }
    }
  `

	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		panic("User node created but cannot be fetched")
	}

	// Unmarshal the invite
	type Root struct {
		user []models.User `json:"user"`
	}
	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal user: \nerr:%s\njson:%s", err, resp.Json)
	}

	return &r.user[0], nil
}
