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


Table: remote_file_definitions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] source                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "bqNZxFuHHVwHXeyolFhnrVqaV",    "type": "yn`TpXkHrPdej\\ZUUgliUVBhd",    "source": "VH]kWpp`u`ErCZqx\\imCVA^xX"}



*/

// RemoteFileDefinitions struct is a row record of the remote_file_definitions table in the laforge-2 database
/*
type RemoteFileDefinitions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Type string `gorm:"column:type;type:TEXT;" json:"type" protobuf:"string,1,opt,name=type"`
    //[ 2] source                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Source null.String `gorm:"column:source;type:TEXT;" json:"source" protobuf:"string,2,opt,name=source"`

}
*/

// TableName sets the insert table name for this struct type
func (r *RemoteFileDefinitions) TableName() string {
	return "remote_file_definitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RemoteFileDefinitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RemoteFileDefinitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RemoteFileDefinitions) Validate(action Action) error {
	return nil
}
