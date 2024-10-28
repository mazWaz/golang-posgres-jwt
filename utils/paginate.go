package utils

import (
	"go-clean/db"
	"math"
)

type Pagination struct {
	TotalRecords int64 `json:"total_records"`
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	NextPage     *int  `json:"next_page"`
	PrevPage     *int  `json:"prev_page"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Paginate function now accepts filters and applies them dynamically
func Paginate(page int, limit int, filters map[string]interface{}, model interface{}, out interface{}) (*PaginatedResponse, error) {
	// Initialize the base query with the provided model
	query := db.Data.Model(model)

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1

	}

	// Apply filters dynamically
	for key, value := range filters {
		if value != "" && value != "%%" { // Exclude empty and unmodified LIKE pattern
			query = query.Where(key, value)
		}
	}

	var totalRecords int64
	query.Count(&totalRecords) // Count the total records after applying filters

	offset := (page - 1) * limit

	if err := query.Offset(offset).Limit(limit).Find(out).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	var nextPage *int
	if page < totalPages {
		next := page + 1
		nextPage = &next
	}

	var prevPage *int
	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}
	response := &PaginatedResponse{
		Data: out,
		Pagination: Pagination{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	}

	return response, nil
}
