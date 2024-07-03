package dto

type Pageable struct {
	Page     int    `json:"page" validate:"required"`
	PageSize int    `json:"pageSize" validate:"required"`
	Sort     string `json:"sort" validate:"required"`
}
