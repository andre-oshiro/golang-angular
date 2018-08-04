package model

// UserStore is a Struct
type UserStore struct{}

// User is a Struct
type User struct {
	ID   string
	Name string
}

// Get function get User from Datastore
func (store *UserStore) Get() (*User, error) {
	// TODO: Get it from Datasore
	user := &User{
		ID:   "UserID",
		Name: "UserName",
	}
	return user, nil
}

// List function Get a list of Users from Datastore
func (store *UserStore) List() ([]*User, error) {
	// TODO: Get it from Datasore
	users := []*User{
		{
			ID:   "UserID1",
			Name: "UserName1",
		},
		{
			ID:   "UserID2",
			Name: "UserName2",
		},
	}
	return users, nil
}
