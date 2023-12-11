package models

type Photo struct {
	PhotoId  string
	owner    string
	likes    []string
	comments []string
}
