package tags

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TagKey uint8

const (
	Undefined TagKey = iota
	ClientId
	Carrier
)

var keyValue = map[string]TagKey{
	"client_id": ClientId,
	"carrier":   Carrier,
}

func (k TagKey) String() string {
	switch k {
	case ClientId:
		return "client_id"
	case Carrier:
		return "carrier"
	}
	return "unknown"
}

func parseKey(s string) (TagKey, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := keyValue[s]
	if !ok {
		return Undefined, fmt.Errorf("%q is not a value tag key", s)
	}
	return value, nil
}

// MarshalText is enough.
//
// func (k TagKey) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(k.String())
// }

func (k TagKey) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
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

func (k *TagKey) UnmarshalText(text []byte) error {
	var err error
	if *k, err = parseKey(string(text)); err != nil {
		return err
	}
	return nil
}
