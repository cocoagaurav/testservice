package model

type User struct {
	AuthId   int    `json:"authid,omitempty"`
	EmailId  string `json:"emailid,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
}

//Method to create a new user that returns Pointer to newly created obj
