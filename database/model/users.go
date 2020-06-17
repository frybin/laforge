package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: users
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] roles_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []


JSON Sample
-------------------------------------
{    "email": "`S`ksUTaaSwEgR^ZwQJdUj]_^",    "roles_id": "eJM]bMs_CqdSrbpiQ^chAjsOE",    "id": "E\\^Cig^lGdbHlYUUUEibVxX\\C"}



*/

// Users struct is a row record of the users table in the laforge-2 database
/*
type Users struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Email null.String `gorm:"column:email;type:TEXT;" json:"email" protobuf:"string,1,opt,name=email"`
    //[ 2] roles_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    RolesID null.String `gorm:"column:roles_id;type:UUID;" json:"roles_id" protobuf:"string,2,opt,name=roles_id"`

}
*/

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}
