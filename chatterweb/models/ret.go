package models

type Page struct {
	Begin        int
	End          int
	PageSize     int
	PageCount    int
	TotalRecords int
	PageNo       int
}
