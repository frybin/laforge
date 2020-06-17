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


Table: included_hosts
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] included_networks_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] host_definitions_id                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "attrs": "KPSnydV\\HgAEBRl`Ja[diamB]",    "id": "\\mHkwmvtJ^yRTQvmtxYmwaKrn",    "included_networks_id": "GDqVHJa[BDLkaLoDNwMfL_mv_",    "host_definitions_id": "[PBtvpwcR]X[HjmgIn[^dvLWq",    "planned_checksum": "hfICaX^cVmDbBlcb^JRltfFML",    "current_checksum": "m_tW`EBeXX^msAr[lBsxE[wZL",    "previous_checksum": "LsEuEb_en`qXIL[r`^ZL_UiEw"}



*/

// IncludedHosts struct is a row record of the included_hosts table in the laforge-2 database
/*
type IncludedHosts struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] included_networks_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedNetworksID null.String `gorm:"column:included_networks_id;type:UUID;" json:"included_networks_id" protobuf:"string,1,opt,name=included_networks_id"`
    //[ 2] host_definitions_id                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    HostDefinitionsID null.String `gorm:"column:host_definitions_id;type:UUID;" json:"host_definitions_id" protobuf:"string,2,opt,name=host_definitions_id"`
    //[ 3] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,3,opt,name=planned_checksum"`
    //[ 4] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,4,opt,name=current_checksum"`
    //[ 5] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,5,opt,name=previous_checksum"`
    //[ 6] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
    Attrs null.String `gorm:"column:attrs;type:JSON;" json:"attrs" protobuf:"string,6,opt,name=attrs"`

}
*/

// TableName sets the insert table name for this struct type
func (i *IncludedHosts) TableName() string {
	return "included_hosts"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IncludedHosts) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IncludedHosts) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IncludedHosts) Validate(action Action) error {
	return nil
}
