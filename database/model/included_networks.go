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


Table: included_networks
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] network_defintions_id                          UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] attrs                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "planned_checksum": "xuDnNkXYUmmD[QKAUnrxJCcUb",    "current_checksum": "JWX]NeyEXDMebKkToyJ_k^xbG",    "previous_checksum": "DHckdmLqkNWZMfmicSn^qBK\\P",    "attrs": "\\esKsZB^EuNTkIHYRpgpnTb`E",    "id": "RsPOaReogFfQ^UH[jVDHKlSwA",    "environments_id": "rfAGFBeyqfn_dI[KAedpfCqyl",    "network_defintions_id": "Gom_giUubNnFopKtmTxkMJ]OR"}



*/

// IncludedNetworks struct is a row record of the included_networks table in the laforge-2 database
/*
type IncludedNetworks struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] environments_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    EnvironmentsID null.String `gorm:"column:environments_id;type:UUID;" json:"environments_id" protobuf:"string,1,opt,name=environments_id"`
    //[ 2] network_defintions_id                          UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    NetworkDefintionsID null.String `gorm:"column:network_defintions_id;type:UUID;" json:"network_defintions_id" protobuf:"string,2,opt,name=network_defintions_id"`
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
func (i *IncludedNetworks) TableName() string {
	return "included_networks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IncludedNetworks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IncludedNetworks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IncludedNetworks) Validate(action Action) error {
	return nil
}
