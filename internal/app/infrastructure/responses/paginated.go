package responses

import "strconv"

type PaginatedResponse[T comparable] struct {
	Data []T `json:"data"`
	Meta struct {
		Total     int64 `json:"total"`
		Page      int   `json:"page"`
		PerPage   int   `json:"per_page"`
		PageCount int   `json:"page_count"`
		Links     struct {
			First string `json:"first"`
			Last  string `json:"last"`
			Prev  string `json:"prev"`
			Next  string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

func NewPaginatedResponse[T comparable](
	data []T,
	total int64,
	page, perPage, pageCount int,
	path string,
) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Data: data,
		Meta: struct {
			Total     int64 `json:"total"`
			Page      int   `json:"page"`
			PerPage   int   `json:"per_page"`
			PageCount int   `json:"page_count"`
			Links     struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			} `json:"links"`
		}{
			Total:     total,
			Page:      page,
			PerPage:   perPage,
			PageCount: pageCount,
			Links: struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			}{
				First: path + "?page=1&per_page=" + strconv.Itoa(perPage),
				Last:  path + "?page=" + strconv.Itoa(pageCount) + "&per_page=" + strconv.Itoa(perPage),
				Prev:  path + "?page=" + strconv.Itoa(page-1) + "&per_page=" + strconv.Itoa(perPage),
				Next:  path + "?page=" + strconv.Itoa(page+1) + "&per_page=" + strconv.Itoa(perPage),
			},
		},
	}
}
