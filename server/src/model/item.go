package model

// ItemStore is a Struct
type ItemStore struct{}

// Item is a Struct
type Item struct {
	ID   string
	Name string
}

// Get function get Item from Datastore
func (store *ItemStore) Get() (*Item, error) {
	// TODO: Get it from Datasore
	item := &Item{
		ID:   "ItemID",
		Name: "ItemName",
	}
	return item, nil
}

// List function Get a list of Items from Datastore
func (store *ItemStore) List() ([]*Item, error) {
	// TODO: Get it from Datasore
	items := []*Item{
		{
			ID:   "ItemID1",
			Name: "ItemName1",
		},
		{
			ID:   "ItemID2",
			Name: "ItemName2",
		},
	}
	return items, nil
}
