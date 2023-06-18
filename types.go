package RbxCloud

import "time"

type Entry_Keys struct {
	Datastore string
	Name      string
}

type Datastores_Response struct {
	Datastores     []Datastores `json:"datastores"`
	NextPageCursor string       `json:"nextPageCursor"`
}

type Datastores struct {
	CreatedTime time.Time `json:"createdTime"`
	Name        string    `json:"name"`
}

type KeyData struct {
	Keys           []Keys `json:"keys"`
	NextPageCursor string `json:"nextPageCursor"`
}

type Keys struct {
	Key string `json:"key"`
}

type Application struct {
	UniversalID string
	APIKey      string
}

type Content struct {
	Entry_Key      string
	Datastore_Name string
	SetEntry       SetEntryContents
	DeleteEntry    Delete
	List           ListEntries
	GetData        Get_Data
}

type Get_Data struct {
	Params string
	Prefix string
	Limit  int
	Cursor string
}

type SetEntryContents struct {
	IsNewDatabase bool
	MatchVersion  string
	Params        string
	UUIDS         []string
	Json          SetEntryJson
}

type SetEntryJson struct {
	EntryJson string
	Content   string
}

type Delete struct {
	Scope  string
	Params string
}

type ListEntries struct {
	Params    string
	Cursor    string
	Scope     string
	Prefix    string
	Limit     int64
	AllScopes bool
}
