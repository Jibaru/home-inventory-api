package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
	"strings"
)

func applyFilters(db *gorm.DB, filter repositories.QueryFilter) *gorm.DB {
	for _, group := range filter.ConditionGroups {
		var groupConditions []string
		var args []interface{}

		for _, condition := range group.Conditions {
			switch condition.Operator {
			case repositories.LikeComparisonOperator:
				groupConditions = append(groupConditions, condition.Field+" LIKE ?")
				args = append(args, condition.Value.(string))
			default:
				groupConditions = append(groupConditions, condition.Field+" "+string(condition.Operator)+" ?")
				args = append(args, condition.Value)
			}
		}

		groupQuery := strings.Join(groupConditions, " "+string(group.Operator)+" ")
		db = db.Where(groupQuery, args...)
	}

	return db
}
