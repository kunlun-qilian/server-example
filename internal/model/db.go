package model

import (
	"database/sql/driver"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/go-courier/sqlx/v2/datatypes"
	"time"
)

var DB = sqlx.NewDatabase("")

type PrimaryID struct {
	ID uint `db:"F_id,autoincrement" json:"-"`
}

type OperationTime struct {
	CreatedAt datatypes.MySQLTimestamp `db:"F_created_at" json:"createdAt"`
	UpdatedAt datatypes.MySQLTimestamp `db:"F_updated_at" json:"updatedAt"`
	DeletedAt datatypes.MySQLTimestamp `db:"F_deleted_at,default=0" json:"-"`
}

func (o *OperationTime) SetNow() {
	o.UpdatedAt = datatypes.MySQLTimestamp(time.Now())
}

func (o *OperationTime) SetNowForCreate() {
	o.SetNow()
	o.CreatedAt = o.UpdatedAt
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tags []Tag

func (tags *Tags) Scan(v interface{}) error {
	return datatypes.JSONScan(v, tags)
}

func (tags Tags) Value() (driver.Value, error) {
	return datatypes.JSONValue(tags)
}

func (Tags) DataType(driverName string) string {
	return "text"
}

type ConditionRules struct {
}

func NewCondRules() *ConditionRules {
	return &ConditionRules{}
}

func (rule *ConditionRules) When(epr bool, condition builder.SqlCondition) builder.SqlCondition {
	if epr {
		return condition
	}
	return nil
}
