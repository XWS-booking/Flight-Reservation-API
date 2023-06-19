package dtos

type PageDto struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

func NewPageDto(pageNumber int, pageSize int) PageDto {
	return PageDto{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
}
