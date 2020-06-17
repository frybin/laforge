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


Table: environments
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] roles_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 5] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "name": "TB`Ke[ZQ[eGdyRlC^Aglv[Jow",    "attrs": "jdJ^ZCxGXiIY_q`lWiKwtKjyh",    "planned_checksum": "xuPNYquDBWyXU`NwCDdhMIBlC",    "current_checksum": "MgYBtJpIivnh[OHnCCNjTjkjI",    "previous_checksum": "tTI`xHge[piJOhKKMddMcaOj_",    "id": "OmYjkkI_OHBk\\PFRmNPHQvRsN",    "competitions_id": "[WmhsNWYQ]]UlWhcCGCTnrcHD",    "roles_id": "Cx]r`YxMY^uHmlieNhKvV^OPA"}



*/

// Environments struct is a row record of the environments table in the laforge-2 database
/*
type Environments struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    CompetitionsID null.String `gorm:"column:competitions_id;type:UUID;" json:"competitions_id" protobuf:"string,1,opt,name=competitions_id"`
    //[ 2] roles_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    RolesID null.String `gorm:"column:roles_id;type:UUID;" json:"roles_id" protobuf:"string,2,opt,name=roles_id"`
    //[ 3] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Name null.String `gorm:"column:name;type:TEXT;" json:"name" protobuf:"string,3,opt,name=name"`
    //[ 4] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
    Attrs null.String `gorm:"column:attrs;type:JSON;" json:"attrs" protobuf:"string,4,opt,name=attrs"`
    //[ 5] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,5,opt,name=planned_checksum"`
    //[ 6] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,6,opt,name=current_checksum"`
    //[ 7] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,7,opt,name=previous_checksum"`

}
*/

// TableName sets the insert table name for this struct type
func (e *Environments) TableName() string {
	return "environments"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Environments) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Environments) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Environments) Validate(action Action) error {
	return nil
}
