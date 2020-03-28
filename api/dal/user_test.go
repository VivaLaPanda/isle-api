package dal

import (
	"testing"

	"github.com/VivaLaPanda/isle-api/api/models"
)

var testUser models.User = models.User{
	Email:      "test@email.com",
	Name:       "test-user",
	Reputation: 0,
	Spent:      0,
	DgraphType: []string{"User"},
	Role:       make([]*models.Role, 1),
}

func createUser() (roleUID, userUID string, err error) {
	db := newClient()

	role := models.Role{Text: "test-role"}
	roleUID, err = NewRole(db, role)
	role.UID = &roleUID
	testUser.Role[0] = &role

	if err != nil {
		return
	}

	userUID, err = NewUser(db, testUser)

	return
}

func TestGetUser(t *testing.T) {
	db := newClient()

	models.LoadSchema(db)

	roleUID, userUID, err := createUser()

	if err != nil {
		t.Errorf("Creating user during TestGetUser failed, err: %s", err)
		return
	}

	user, err := GetUser(db, userUID)
	if err != nil {
		t.Errorf("TestGetUser, err: %s", err)
		return
	}

	if user.Name != testUser.Name {
		t.Errorf("TestGetUser failed, user.Name != testUser.Name: %s != %s", user.Name, testUser.Name)
	}

	if *user.Role[0].UID != roleUID {
		t.Errorf("TestGetUser failed, ser.Role.UID != roleUID: %s != %s", *user.Role[0].UID, roleUID)
	}
}
