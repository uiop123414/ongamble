package models

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}
