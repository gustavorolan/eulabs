package dto

import "math"

type Page struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	Sort       string      `json:"sort"`
	Data       interface{} `json:"data"`
	TotalPages int64       `json:"totalPages"`
	Total      int64       `json:"total"`
}

func NewPage(data interface{}, total int64, pageable *Pageable) Page {

	totalPages := int64(math.Ceil(float64(total) / float64(pageable.PageSize)))

	return Page{
		Page:       pageable.Page,
		PageSize:   pageable.PageSize,
		Sort:       pageable.Sort,
		Data:       data,
		TotalPages: totalPages,
		Total:      total,
	}
}
