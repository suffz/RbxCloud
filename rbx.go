package RbxCloud

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"errors"
	"net/http"
	"strings"
)

var (
	url              = "https://apis.roblox.com/datastores/v1/universes/%v"
	objects_url      = "https://apis.roblox.com/datastores/v1/universes/%v/standard-datastores/datastore/entries/entry"
	list_objects_url = "https://apis.roblox.com/datastores/v1/universes/%v/standard-datastores/datastore/entries"
	increment_url    = "https://apis.roblox.com/datastores/v1/universes/%v/standard-datastores/datastore/entries/entry/increment"
	version_url      = "https://apis.roblox.com/datastores/v1/universes/%v/standard-datastores/datastore/entries/entry/versions/version"
	datastores_url   = "https://apis.roblox.com/datastores/v1/universes/%v/standard-datastores"
)

func Init(global, key string) Application {
	return Application{
		UniversalID: global,
		APIKey:      key,
	}
}

func MD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (App *Application) ListEntries(Info Content) []Entry_Keys {
	var Returndata []Entry_Keys
Exit:
	for {
		Info.Check("list-entries")
		fmt.Println(fmt.Sprintf(list_objects_url, App.UniversalID) + Info.List.Params)
		req, _ := http.NewRequest("GET", fmt.Sprintf(list_objects_url, App.UniversalID)+Info.List.Params, nil)
		req.Header.Add("x-api-key", App.APIKey)
		if resp, err := http.DefaultClient.Do(req); err == nil {
			body, _ := io.ReadAll(resp.Body)
			var Data KeyData
			json.Unmarshal(body, &Data)
			if Data.NextPageCursor != "" {
				Info.List.Cursor = Data.NextPageCursor
				for _, keys := range Data.Keys {
					DB, Name := "", ""
					if strings.Contains(keys.Key, "/") {
						values := strings.Split(keys.Key, "/")
						switch len(values) {
						case 1:
							DB = values[0]
						case 2:
							DB = values[0]
							Name = values[1]
						default:
							DB = keys.Key
						}
					} else {
						DB = keys.Key
					}
					Returndata = append(Returndata, Entry_Keys{
						Datastore: DB,
						Name:      Name,
					})
				}
			} else {
				break Exit
			}
		}
	}
	return Returndata
}

func (App *Application) GetDatastores(I Content) (Info Datastores_Response) {
	I.Check("get-data")
	if req, err := http.NewRequest("GET", fmt.Sprintf(datastores_url, App.UniversalID)+I.GetData.Params, nil); err == nil {
		req.Header.Add("x-api-key", App.APIKey)
		if resp, err := http.DefaultClient.Do(req); err == nil {
			body, _ := io.ReadAll(resp.Body)
			json.Unmarshal(body, &Info)
		}
	}
	return
}

func (App *Application) DeleteEntry(Info Content) (bool, error) {
	Info.Check("delete-entry")
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf(objects_url, App.UniversalID)+Info.SetEntry.Params, nil)
	req.Header.Add("x-api-key", App.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		return resp.StatusCode == 204, nil
	}
	return false, err
}

func (App *Application) SetEntry(Info Content) (bool, error) {
	Info.Check("set-entry")
	req, _ := http.NewRequest("POST", fmt.Sprintf(objects_url, App.UniversalID)+Info.SetEntry.Params, bytes.NewBuffer([]byte(Info.SetEntry.Json.Content)))
	req.Header.Add("x-api-key", App.APIKey)
	req.Header.Add("content-md5", MD5(Info.SetEntry.Json.Content))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("roblox-entry-userids", fmt.Sprintf("[%v]", strings.Join(Info.SetEntry.UUIDS, ",")))
	req.Header.Add("roblox-entry-attributes", Info.SetEntry.Json.EntryJson)
	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		switch resp.StatusCode {
		case 200:
			return true, nil
		case 412:
			Info.SetEntry.IsNewDatabase = false
			return App.SetEntry(Info)
		}
	}
	return false, err
}

func (App *Application) GetEntry(Info Content) (Body []byte, err error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf(objects_url, App.UniversalID)+fmt.Sprintf("?datastoreName=%v&entryKey=%v", Info.Datastore_Name, Info.Entry_Key), nil)
	req.Header.Add("x-api-key", App.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		switch resp.StatusCode {
		case 200, 204:
			Body, err = io.ReadAll(resp.Body)
			return
		default:
			Body, err = io.ReadAll(resp.Body)
			return nil, errors.New(fmt.Sprintf(`Error: Unknown Status Code %v | %v`, resp.StatusCode, string(Body)))
		}
	} else {
		return nil, err
	}
}
