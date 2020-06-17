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


Table: builds
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] build_configs_id                               UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "current_checksum": "R_WSBQrlfohQ[AuZaq^ohwLth",    "previous_checksum": "GQLcZsq_XgWGdfKIajTsSxmci",    "id": "bhTVVMoYJFj^rsPTSMBjVJfVg",    "environments_id": "gXksQEatneXZd`fnL^tFTnVOG",    "build_configs_id": "HUKwDsjakc^nPZOXvwqC]QYbO",    "state": "wKOVbZ[FhVKfyJstrcieJDHwh",    "planned_checksum": "[e]oQu`RRIHTU]NMk]clVA\\bm"}



*/

// Builds struct is a row record of the builds table in the laforge-2 database
/*
type Builds struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    EnvironmentsID null.String `gorm:"column:environments_id;type:UUID;" json:"environments_id" protobuf:"string,1,opt,name=environments_id"`
    //[ 2] build_configs_id                               UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    BuildConfigsID null.String `gorm:"column:build_configs_id;type:UUID;" json:"build_configs_id" protobuf:"string,2,opt,name=build_configs_id"`
    //[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    State null.String `gorm:"column:state;type:TEXT;" json:"state" protobuf:"string,3,opt,name=state"`
    //[ 4] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,4,opt,name=planned_checksum"`
    //[ 5] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,5,opt,name=current_checksum"`
    //[ 6] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,6,opt,name=previous_checksum"`

}
*/

// TableName sets the insert table name for this struct type
func (b *Builds) TableName() string {
	return "builds"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Builds) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Builds) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Builds) Validate(action Action) error {
	return nil
}
