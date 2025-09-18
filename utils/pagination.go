package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	Search     string      `json:"search,omitempty"`
}

func Paginate(c *gin.Context, db *gorm.DB, model interface{}, out interface{}, searchField string) (*Pagination, error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var total int64
	query := db.Model(model)

	if search != "" && searchField != "" {
		query = query.Where(fmt.Sprintf("%s LIKE ?", searchField), "%"+search+"%")
	}

	query.Count(&total)

	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(out).Error; err != nil {
		return nil, err
	}

	return &Pagination{
		Data:       out,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
		Search:     search,
	}, nil

}
