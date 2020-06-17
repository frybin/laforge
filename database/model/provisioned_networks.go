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


Table: provisioned_networks
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] included_networks_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] included_teams_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "planned_checksum": "malUpigVnBYER\\`ml^C_BY`Ad",    "current_checksum": "TXmnq]_kjlgn`BeLtihA^HNv`",    "previous_checksum": "OIIL]LkAqonGkDyIII`RBUolt",    "id": "PDjYHKOu`DSqZJGoVidQHkku`",    "included_networks_id": "eN]qDgmljmFolCDpISjjFOyQM",    "included_teams_id": "KtDxvesPEj`mxJhDGjxwRhaXE",    "state": "UEWpmD\\^GgRaqbdMmauyZPHsq"}



*/

// ProvisionedNetworks struct is a row record of the provisioned_networks table in the laforge-2 database
/*
type ProvisionedNetworks struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] included_networks_id                           UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedNetworksID null.String `gorm:"column:included_networks_id;type:UUID;" json:"included_networks_id" protobuf:"string,1,opt,name=included_networks_id"`
    //[ 2] included_teams_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedTeamsID null.String `gorm:"column:included_teams_id;type:UUID;" json:"included_teams_id" protobuf:"string,2,opt,name=included_teams_id"`
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
func (p *ProvisionedNetworks) TableName() string {
	return "provisioned_networks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProvisionedNetworks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProvisionedNetworks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProvisionedNetworks) Validate(action Action) error {
	return nil
}
