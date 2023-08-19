package tags

import (
	"encoding/json"
	"testing"
)

type TagSet struct {
	Key TagKey `json:"key"`
}

func TestJSON(t *testing.T) {
	ts := &TagSet{ClientId}
	data, err := json.Marshal(ts)
	if err != nil {
		t.Error("Serialization failed")
	}
	if string(data) != "{\"key\":\"client_id\"}" {
		t.Errorf("Wrong serialized value %s", string(data))
	}

	var deTs TagSet
	if err = json.Unmarshal(data, &deTs); err != nil {
		t.Errorf("De-serialization failed: %s", err)
	}
	if deTs.Key != ClientId {
		t.Errorf("Wrong de-serialized value: %s", deTs.Key)
	}

}

func TestJSONMap(t *testing.T) {
	tsMap := make(map[TagKey]string)
	tsMap[ClientId] = "123"
	data, err := json.Marshal(tsMap)
	if err != nil {
		t.Error("Serialization failed")
		return
	}
	if string(data) != "{\"client_id\":\"123\"}" {
		t.Errorf("Wrong serialized value %s", string(data))
		return
	}

	var deTsMap map[TagKey]string
	if err = json.Unmarshal(data, &deTsMap); err != nil {
		t.Errorf("De-serialization failed: %s", err)
		return
	}
	if deTsMap[ClientId] != "123" {
		t.Errorf("Wrong serialized value: %s", deTsMap[ClientId])
		return
	}
}
