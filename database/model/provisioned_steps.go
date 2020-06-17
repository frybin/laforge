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


Table: provisioned_steps
[ 0] id                                             UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] provisioned_hosts_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "jRqwWyqNbaBreHPO]RZPpegSS",    "state": "gXlBvM\\sYRWLvvfGJsuQKGHjq",    "planned_checksum": "nkGQlvLRyaSxcFNT`UQIwimQu",    "current_checksum": "XCpeyScTbtFfHxBVdhGvKxfjE",    "previous_checksum": "CijiVY[_I\\iPrwIiqefgxwnkv",    "provisioned_hosts_id": "K`xgdKF^EFsNQR]\\SFChQDBsD"}


Comments
-------------------------------------
[ 0] Warning table: provisioned_steps does not have a primary key defined, setting col position 1 id as primary key
Warning table: provisioned_steps primary key column id is nullable column, setting it as NOT NULL




*/

// ProvisionedSteps struct is a row record of the provisioned_steps table in the laforge-2 database
/*
type ProvisionedSteps struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    State null.String `gorm:"column:state;type:TEXT;" json:"state" protobuf:"string,1,opt,name=state"`
    //[ 2] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,2,opt,name=planned_checksum"`
    //[ 3] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,3,opt,name=current_checksum"`
    //[ 4] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,4,opt,name=previous_checksum"`
    //[ 5] provisioned_hosts_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    ProvisionedHostsID null.String `gorm:"column:provisioned_hosts_id;type:UUID;" json:"provisioned_hosts_id" protobuf:"string,5,opt,name=provisioned_hosts_id"`

}
*/

// TableName sets the insert table name for this struct type
func (p *ProvisionedSteps) TableName() string {
	return "provisioned_steps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProvisionedSteps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProvisionedSteps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProvisionedSteps) Validate(action Action) error {
	return nil
}
