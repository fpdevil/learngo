package main

import (
	"encoding/json"
	"fmt"
)

// JSON Encoding
// Encoding JSON data using structs, maps and slices

type permissions map[string]bool

// encode the fields with differemt names now
// adding a field tag `json:"-"` to Password field inorder to make it
// available to other packages, but do not encode the value
// user type is exported now
type user struct {
	Name     string `json:"username"` // name is now exported
	Password string `json:"-"`        // password is now exported

	// here omitempty prevents encoding of zero values
	// it can also be just `json:",omitempty"` to only escape empty value
	// encoding still printing the original value encoded
	Permissions permissions `json:"perms,omitempty"` // convert permissions to named field for exporting
}

// adding a field tag `json:"-"` to Password field inorder to make it
// available to other packages, but do not encode the value
// user type is exported now
// type user struct {
// 	Name        string      // name is now exported
// 	Password    string      `json:"-"` // password is now exported
// 	Permissions permissions // convert permissions to named field for exporting
// }

// user type is exported now
// type user struct {
// 	Name        string      // name is now exported
// 	Password    string      // password is now exported
// 	Permissions permissions // convert permissions to named field for exporting
// }

// type user struct {
// 	Name        string // name is now exported
// 	Password    string // password is now exported
// 	permissions        // permissions is now yet exported and is anonymous
// }

// --- Unexported type original
// type user struct {
// 	name     string
// 	password string
// 	permissions
// 	// converted perssions to an anonymous type
// 	// anonymous types need not be another struct
// 	// permissions map[string]bool
// }

func main() {
	fmt.Println("24-structs/06-encoding")

	// create a couple of users in a list as below
	// users := []user{
	// 	{name: "mickey", password: "mouse001", permissions: nil},
	// 	{name: "gyro", password: "gearlooze101", permissions: permissions{"admin": true}},
	// 	{name: "tom", password: "jerry999", permissions: permissions{"write": true}},
	// }
	users := []user{
		{"mickey", "mouse001", nil},
		{"gyro", "gearlooze101", permissions{"admin": true}},
		{"tom", "jerry999", permissions{"write": true}},
		{"tintin", "snowy123", permissions{"read": true}},
	}

	// encoding the users type to json by marshalling
	// out, err := json.Marshal(users)

	// using MarshalIndent for pretty printing
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("%s\n", string(out))

	// This outputs encoded json: [{},{},{}]
	// It is because the package json only encodes `Exported Fields`.
	// as the fields are lower case and unexported
	// convert them to Upper case and try
	// fmt.Printf("encoded json: %s\n", string(out))

}
