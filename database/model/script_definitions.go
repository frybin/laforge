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


Table: script_definitions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] source_files                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] runtime                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "lq]MlxhrOUqrqVgvWHCVgqkZA",    "type": "xOiZUrS]gtSH]SWZcPxWKKwc_",    "source_files": "oFVdYHEPigrEyUgtYedadQryI",    "runtime": "oRicBq\\x^eLriQA]l^Flg\\ajT"}



*/

// ScriptDefinitions struct is a row record of the script_definitions table in the laforge-2 database
/*
type ScriptDefinitions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] type                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Type string `gorm:"column:type;type:TEXT;" json:"type" protobuf:"string,1,opt,name=type"`
    //[ 2] source_files                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    SourceFiles null.String `gorm:"column:source_files;type:TEXT;" json:"source_files" protobuf:"string,2,opt,name=source_files"`
    //[ 3] runtime                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Runtime null.String `gorm:"column:runtime;type:TEXT;" json:"runtime" protobuf:"string,3,opt,name=runtime"`

}
*/

// TableName sets the insert table name for this struct type
func (s *ScriptDefinitions) TableName() string {
	return "script_definitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ScriptDefinitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ScriptDefinitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ScriptDefinitions) Validate(action Action) error {
	return nil
}
