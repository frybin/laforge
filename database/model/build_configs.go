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


Table: build_configs
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] provider                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 3] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "RKlvbHgm\\R\\IZGcpGESIZrc[n",    "provider": "WHGdkF\\a_kibJUcYbS\\CkFAAq",    "attrs": "FMB[[]xNmyIZeOdAoh]ZckE]g",    "competitions_id": "wXgZPmHKHajpVvEIqg_^dJc\\O"}



*/

// BuildConfigs struct is a row record of the build_configs table in the laforge-2 database
/*
type BuildConfigs struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] provider                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Provider null.String `gorm:"column:provider;type:TEXT;" json:"provider" protobuf:"string,1,opt,name=provider"`
    //[ 2] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
    Attrs null.String `gorm:"column:attrs;type:JSON;" json:"attrs" protobuf:"string,2,opt,name=attrs"`
    //[ 3] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    CompetitionsID null.String `gorm:"column:competitions_id;type:UUID;" json:"competitions_id" protobuf:"string,3,opt,name=competitions_id"`

}
*/

// TableName sets the insert table name for this struct type
func (b *BuildConfigs) TableName() string {
	return "build_configs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BuildConfigs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BuildConfigs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BuildConfigs) Validate(action Action) error {
	return nil
}
