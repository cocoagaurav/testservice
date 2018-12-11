package model

type Post struct {
	Name        string `json:"name,omitempty"`
	EmailId     string `json:"emailid,omitempty"`
	Title       string `json:"title,omitempty"`
	Discription string `json:"discription,omitempty"`
}
