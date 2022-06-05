package common

import (
	"fmt"
	"time"

	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Model represents a base model with primary id key and timestamps
type Model struct {
	ID int `gorm:"primary_key" json:"id" example:"1"`

	CreatedBy  int        `json:"createdBy"`
	CreatedAt  time.Time  `json:"createdAt" example:"2018-07-07T13:38:13+08:00"`
	ModifiedAt time.Time  `json:"modifiedAt" example:"2018-07-07T13:38:13+08:00"`
	DeletedAt  *time.Time `json:"-" example:"2018-07-07T13:38:13+08:00"`

	Extras *json.RawMessage `json:"extras,omitempty" sql:"type:jsonb" gorm:"default:'{}'"` // reserved
}

type gormLogger struct{}

func (*gormLogger) Print(v ...interface{}) {
	//if v[0] == "sql" {
	//	logrus.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Infof("(%s) %s", v[2], v[3])
	//}
	//if v[0] == "log" {
	//	logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Info(v[2])
	//}
}

// createCallback will set `CreatedAt`, `ModifiedAt` when creating
func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateCallback will set `ModifiedAt` when updating
func updateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedAt", time.Now())
	}
}

// deleteCallback will set `DeletedAt` when updating
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			deletedField, hasDeletedField := scope.FieldByName("Deleted")
			if hasDeletedField {
				scope.Raw(fmt.Sprintf(
					"UPDATE %v SET %v=%v,%v=%v%v%v",
					scope.QuotedTableName(),

					scope.Quote(deletedAtField.DBName),
					scope.AddToVars(time.Now()),
					scope.Quote(deletedField.DBName),
					scope.AddToVars(time.Now()),

					addExtraSpaceIfExist(scope.CombinedConditionSql()),
					addExtraSpaceIfExist(extraOption),
				)).Exec()
			} else {
				scope.Raw(fmt.Sprintf(
					"UPDATE %v SET %v=%v%v%v",
					scope.QuotedTableName(),
					scope.Quote(deletedAtField.DBName),
					scope.AddToVars(time.Now()),
					addExtraSpaceIfExist(scope.CombinedConditionSql()),
					addExtraSpaceIfExist(extraOption),
				)).Exec()
			}
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
