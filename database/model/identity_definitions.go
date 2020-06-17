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


Table: identity_definitions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "email": "FOcRFhsfUB`QmVkBThWgSgMbE",    "id": "PkWjMq^vlE_JTxOuIxfKxZIsP"}



*/

// IdentityDefinitions struct is a row record of the identity_definitions table in the laforge-2 database
/*
type IdentityDefinitions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Email null.String `gorm:"column:email;type:TEXT;" json:"email" protobuf:"string,1,opt,name=email"`

}
*/

// TableName sets the insert table name for this struct type
func (i *IdentityDefinitions) TableName() string {
	return "identity_definitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IdentityDefinitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IdentityDefinitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IdentityDefinitions) Validate(action Action) error {
	return nil
}
