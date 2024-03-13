package models

import (
	"github.com/qiniu/qmgo/field"
)

type User struct {
	field.DefaultField `bson:",inline"` 
	Username           string `bson:"username"`
	Email              string `bson:"email"`
	Password           string `bson:"password"`
}

// Implement CustomFields() field.CustomFieldsBuilder 
// And define the custom fields
func (u *User) CustomFields() field.CustomFieldsBuilder {
    return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt").SetId("MyId")
}
