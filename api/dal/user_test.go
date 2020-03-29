package dal

import (
	"context"
	"testing"
	"time"

	"github.com/VivaLaPanda/isle-api/api/models"
	"github.com/dgraph-io/dgo/protos/api"
)

func init() {
	db := newClient()

	db.Alter(context.Background(), &api.Operation{
		DropAll: true,
	})
}

var now = time.Now()

var testUserA models.User = models.User{
	Email:      "test-a@email.com",
	Name:       "test-user-a",
	Reputation: 0,
	Spent:      0,
	Joined:     &now,
	DgraphType: []string{"User"},
	Role:       make([]*models.Role, 1),
}

var testUserB models.User = models.User{
	Email:      "test-b@email.com",
	Name:       "test-user-b",
	Reputation: 0,
	Spent:      0,
	Joined:     &now,
	DgraphType: []string{"User"},
	Role:       make([]*models.Role, 1),
}
var testInvite models.Invite = models.Invite{
	Code:       "test-invite",
	DgraphType: []string{"Invite"},
}
var testRole models.Role = models.Role{
	Text: "test-role",
}

func createUser() (err error) {
	db := newClient()

	roleUID, err := NewRole(db, testRole)
	testRole.UID = &roleUID
	testUserA.Role[0] = &testRole
	testUserB.Role[0] = &testRole

	if err != nil {
		return
	}

	userAUID, err := NewUser(db, testUserA, "")
	testUserA.UID = &userAUID

	inviteUID, err := NewInvite(db, testInvite, testUserA)
	testInvite.UID = &inviteUID
	if err != nil {
		return
	}

	userBUID, err := NewUser(db, testUserB, testInvite.Code)
	testUserB.UID = &userBUID
	if err != nil {
		return
	}

	return
}

func TestGetUser(t *testing.T) {
	db := newClient()

	models.LoadSchema(db)

	err := createUser()

	if err != nil {
		t.Errorf("Creating user during TestGetUser failed, err: %s", err)
		return
	}

	userA, err := GetUser(db, *testUserA.UID)
	if err != nil {
		t.Errorf("TestGetUser, err: %s", err)
		return
	}

	if userA.Name != testUserA.Name {
		t.Errorf("TestGetUser failed, userA.Name != testUser.Name: %s != %s", userA.Name, testUserA.Name)
	}

	if *userA.Role[0].UID != *testRole.UID {
		t.Errorf("TestGetUser failed, userA.Role.UID != roleUID: %s != %s", *userA.Role[0].UID, *testRole.UID)
	}

	userB, err := GetUser(db, *testUserB.UID)
	if err != nil {
		t.Errorf("TestGetUser, err: %s", err)
		return
	}

	if userB.InvitedBy[0].Name != userA.Name {
		t.Errorf("TestGetUser failed, userB.InvitedBy[0].Name != userA.Name: %s != %s", userB.InvitedBy[0].Name, userA.Name)
	}
}
