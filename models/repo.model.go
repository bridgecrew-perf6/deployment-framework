package models

type RepositoryFile struct {
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	Name   string `json:"name"`
	Folder string `json:"folder"`
	Repo   string `json:"repo"`
}
