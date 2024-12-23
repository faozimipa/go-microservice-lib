package pagination

import (
	"errors"
	"strconv"

	"github.com/faozimipa/go-microservice-lib/response"
)

type Page struct {
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

func ValidateAndGetPage(page, size, sortBy, sortOrder string) (*Page, error) {
	p := Page{SortBy: sortBy}

	if pageInt, err := strconv.Atoi(page); err != nil {
		return nil, errors.New("'page' parameter must be a number")
	} else {
		p.Page = pageInt
	}

	if sizeInt, err := strconv.Atoi(size); err != nil {
		return nil, errors.New("'size' parameter must be a number")
	} else {
		p.Size = sizeInt
	}

	if sortOrder == "asc" || sortOrder == "desc" {
		p.SortOrder = sortOrder
		return &p, nil
	} else {
		return nil, errors.New("'sortOrder' parameter must be 'asc' or 'desc'")
	}
}

func Paginator(p Page, total int) response.PaginationResponse {
	return response.PaginationResponse{
		PageNumber: p.Page,
		PageSize:   p.Size,
		Total:      total,
	}
}
