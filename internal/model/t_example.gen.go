// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTExample = "t_example"

// TExample mapped from table <t_example>
type TExample struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name    string `gorm:"column:name" json:"name"` // Name
	CarType int32  `gorm:"column:car_type" json:"car_type"`
}

// TableName TExample's table name
func (*TExample) TableName() string {
	return TableNameTExample
}
