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


Table: provisioned_teams
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] included_teams_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] build_configs_id                               UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "current_checksum": "loSRhnZMHrPWpUyyqZcmiE\\dr",    "previous_checksum": "bLuMbntr[arOPkhDd\\sHHteFN",    "id": "cykkey_wnjJtUD]tTpff]yjfn",    "included_teams_id": "fiwZqtsquCR`ZVtBJAcy[nMVQ",    "build_configs_id": "^Ty[nJKLLdaSWdYSSuf`oiTpB",    "state": "PC[EPP`QmsNKLoQe\\qpisYNAJ",    "planned_checksum": "Pk\\IC[nSUQddDtUDs[UtT_ttX"}



*/

// ProvisionedTeams struct is a row record of the provisioned_teams table in the laforge-2 database
/*
type ProvisionedTeams struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] included_teams_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedTeamsID null.String `gorm:"column:included_teams_id;type:UUID;" json:"included_teams_id" protobuf:"string,1,opt,name=included_teams_id"`
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
func (p *ProvisionedTeams) TableName() string {
	return "provisioned_teams"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProvisionedTeams) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProvisionedTeams) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProvisionedTeams) Validate(action Action) error {
	return nil
}
