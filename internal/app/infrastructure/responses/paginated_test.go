package responses

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewPaginatedResponse(t *testing.T) {
	data := []string{"some test", "another test"}
	total := int64(1000)
	page := 1
	perPage := 2
	pageCount := len(data)
	path := "/tests"

	response := NewPaginatedResponse(data, total, page, perPage, pageCount, path)

	assert.Equal(t, data, response.Data)
	assert.Equal(t, total, response.Meta.Total)
	assert.Equal(t, page, response.Meta.Page)
	assert.Equal(t, perPage, response.Meta.PerPage)
	assert.Equal(t, pageCount, response.Meta.PageCount)
	assert.Equal(t, path+"?page="+strconv.Itoa(page)+"&per_page="+strconv.Itoa(perPage), response.Meta.Links.First)
	assert.Equal(t, path+"?page="+strconv.Itoa(pageCount)+"&per_page="+strconv.Itoa(perPage), response.Meta.Links.Last)
	assert.Equal(t, path+"?page="+strconv.Itoa(page-1)+"&per_page="+strconv.Itoa(perPage), response.Meta.Links.Prev)
	assert.Equal(t, path+"?page="+strconv.Itoa(page+1)+"&per_page="+strconv.Itoa(perPage), response.Meta.Links.Next)

	for _, v := range response.Data {
		assert.Contains(t, data, v)
	}
}
