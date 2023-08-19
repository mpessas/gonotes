package tags

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TagKey string

const (
	Undefined TagKey = "undefined"
	ClientId  TagKey = "client_id"
	Carrier   TagKey = "carrier"
)

var keyValue = map[string]TagKey{
	"client_id": ClientId,
	"carrier":   Carrier,
}

func (k TagKey) String() string {
	switch k {
	case ClientId:
		return "ClientId"
	case Carrier:
		return "Carrier"
	}
	return "Undefined"
}

func parseKey(s string) (TagKey, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := keyValue[s]
	if !ok {
		return Undefined, fmt.Errorf("%q is not a value tag key", s)
	}
	return value, nil
}

func (k TagKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}

func (k *TagKey) UnmarshalJSON(data []byte) error {
	var key string
	var err error
	if err = json.Unmarshal(data, &key); err != nil {
		return err
	}
	if *k, err = parseKey(key); err != nil {
		return err
	}
	return nil
}
