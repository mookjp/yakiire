package lib

import (
	"encoding/json"
)

// Doc represents documentation of Firestore
type Doc struct {
	data map[string]interface{}
}

// String returns JSON string of documentation
func (d *Doc) String() string {
	j, err := json.Marshal(d.data)
	if err != nil {
		panic(err)
	}
	return string(j)
}
