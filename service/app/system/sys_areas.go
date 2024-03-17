package system

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/gin-gonic/gin"
)

type BaseSystemAreas struct {
}
type Area struct {
	Id        int
	Title     string
	Code      int
	Ancestry  string
	CreatedAt string `gorm:"type:datetime"`
	UpdatedAt string `gorm:"type:datetime"`
}

type AreaResponse struct {
	Title    string         `json:"title"`
	Code     int            `json:"code"`
	Children []AreaResponse `json:"children"`
}

func (t *BaseSystemAreas) SystemAreasServices(c *gin.Context) (data []AreaResponse, err error) {
	var areas []Area
	// 错误处理1: 数据库查询
	if result := global.TREND_DB.Table("areas").Find(&areas); result.Error != nil {
		return nil, result.Error
	}
	areaMap := make(map[string]AreaResponse)

	for _, area := range areas {
		if area.Ancestry != "" {
			areaResponse := AreaResponse{
				Title: area.Title,
				Code:  area.Code,
			}

			if parent, ok := areaMap[area.Ancestry]; ok {
				parent.Children = append(parent.Children, areaResponse)
				areaMap[area.Ancestry] = parent
			} else {
				parent := AreaResponse{
					Children: []AreaResponse{areaResponse},
				}
				if len(area.Ancestry) == 0 {
					return nil, fmt.Errorf("Empty ancestry for area with code: %d", area.Code)
				}
				areaMap[area.Ancestry] = parent
			}

			continue
		}

		areaMap[fmt.Sprintf("%d", area.Code)] = AreaResponse{
			Title: area.Title,
			Code:  area.Code,
		}
	}

	areaResponses := make([]AreaResponse, 0, len(areaMap))

	for _, areaResponse := range areaMap {
		areaResponses = append(areaResponses, areaResponse)
	}

	return areaResponses, nil
}
