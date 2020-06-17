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


Table: command_definition
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "mFTmm[sfngTUbNIy_sQCtZsVM",    "type": "ewU[ktuSad_[E]mKrraARilKs",    "name": "hUucibrRWslGsIC\\OFIugixGu"}



*/

// CommandDefinition struct is a row record of the command_definition table in the laforge-2 database
/*
type CommandDefinition struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Type string `gorm:"column:type;type:TEXT;" json:"type" protobuf:"string,1,opt,name=type"`
    //[ 2] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Name null.String `gorm:"column:name;type:TEXT;" json:"name" protobuf:"string,2,opt,name=name"`

}
*/

// TableName sets the insert table name for this struct type
func (c *CommandDefinition) TableName() string {
	return "command_definition"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CommandDefinition) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CommandDefinition) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CommandDefinition) Validate(action Action) error {
	return nil
}
