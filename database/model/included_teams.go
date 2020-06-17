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


Table: included_teams
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] number                                         INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] tags                                           JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 4] enabled                                        BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 5] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "tags": "W_WMSpC[Jr[_GwScoUYkmIkKX",    "enabled": false,    "planned_checksum": "HTG]cNkfjeHxNjo_OX\\uEGuvB",    "current_checksum": "ewemk[mPBabZyjhJMNiIqjq`r",    "previous_checksum": "[GnKo_l[NYM`xs^OcGPOwjQ^^",    "id": "\\aovSpsbWukcBXrK\\kaILrxdr",    "environments_id": "lGSL[QWnu]xilrnoi_FlSYNnQ",    "number": 97}



*/

// IncludedTeams struct is a row record of the included_teams table in the laforge-2 database
/*
type IncludedTeams struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    EnvironmentsID null.String `gorm:"column:environments_id;type:UUID;" json:"environments_id" protobuf:"string,1,opt,name=environments_id"`
    //[ 2] number                                         INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
    Number int32 `gorm:"column:number;type:INT2;" json:"number" protobuf:"int32,2,opt,name=number"`
    //[ 3] tags                                           JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
    Tags null.String `gorm:"column:tags;type:JSON;" json:"tags" protobuf:"string,3,opt,name=tags"`
    //[ 4] enabled                                        BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
    Enabled null.Int `gorm:"column:enabled;type:BOOL;" json:"enabled" protobuf:"bool,4,opt,name=enabled"`
    //[ 5] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,5,opt,name=planned_checksum"`
    //[ 6] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,6,opt,name=current_checksum"`
    //[ 7] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,7,opt,name=previous_checksum"`

}
*/

// TableName sets the insert table name for this struct type
func (i *IncludedTeams) TableName() string {
	return "included_teams"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IncludedTeams) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IncludedTeams) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IncludedTeams) Validate(action Action) error {
	return nil
}
