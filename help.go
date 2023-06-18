package RbxCloud

import (
	"fmt"
	"strings"
)

func StripWhitespace(s string) string {
	return strings.Replace(strings.Replace(s, "\t", "", -1), "\n", "", -1)
}

func (Info *Content) Check(t string) {
	switch t {
	case "get-data":
		if Info.GetData.Params == "" {
			Info.GetData.Params = fmt.Sprintf("?prefix=%v&limit=%v&cursor=%v", Info.GetData.Prefix, Info.GetData.Limit, Info.GetData.Cursor)
		}
	case "list-entries":
		if Info.SetEntry.Params == "" {
			Info.List.Params = fmt.Sprintf("?datastoreName=%v&scope=%v&prefix=%v&limit=%v&allScopes=%v&cursor=%v", Info.Datastore_Name, Info.List.Scope, Info.List.Prefix, Info.List.Limit, Info.List.AllScopes, Info.List.Cursor)
		}
	case "delete-entry":
		if Info.DeleteEntry.Params == "" {
			Info.DeleteEntry.Params = fmt.Sprintf("?datastoreName=%v&entryKey=%v&scope=%v", Info.Datastore_Name, Info.Entry_Key, Info.DeleteEntry.Scope)
		}
	case "set-entry":
		if Info.SetEntry.Params == "" {
			Info.SetEntry.Params = fmt.Sprintf("?datastoreName=%v&entryKey=%v&exclusiveCreate=%v", Info.Datastore_Name, Info.Entry_Key, Info.SetEntry.IsNewDatabase)
			if Info.SetEntry.MatchVersion != "" {
				Info.SetEntry.Params += "&matchVersion=" + Info.SetEntry.MatchVersion
			}
		}
		if Info.SetEntry.Json.EntryJson == "" {
			Info.SetEntry.Json.EntryJson = "{}"
		} else {
			Info.SetEntry.Json.EntryJson = StripWhitespace(Info.SetEntry.Json.EntryJson)
		}
		if len(Info.SetEntry.UUIDS) == 0 {
			Info.SetEntry.UUIDS = append(Info.SetEntry.UUIDS, Info.Entry_Key)
		}
	}
}
