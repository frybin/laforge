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


Table: provisioned_hosts
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] included_hosts_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] provisioned_networks_id                        UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] ip_address                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] conn_info                                      JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 6] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "JAAvXC`B]tiEiOYNXvcgs^A^d",    "included_hosts_id": "juDJNVMZcksaYydIcv^hRDAhB",    "provisioned_networks_id": "XrkiEOjmmqZ[`inFxeuvvmXQA",    "previous_checksum": "`HAJUrOvxFjdvtXbu[ZROHaoT",    "state": "fpA]Q^CLhKtLmor[Vkne]_T]K",    "ip_address": "cMVuIoRpQWk_souciv^uyVVCI",    "conn_info": "aSwedjiP^uELwJ`WixTiJmcBX",    "planned_checksum": "ajPvtCkw_SRjqFuBgU`q`YXwg",    "current_checksum": "j\\KYKmiFkgf`T\\rPgGrx]nEfl"}



*/

// ProvisionedHosts struct is a row record of the provisioned_hosts table in the laforge-2 database
/*
type ProvisionedHosts struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] included_hosts_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedHostsID null.String `gorm:"column:included_hosts_id;type:UUID;" json:"included_hosts_id" protobuf:"string,1,opt,name=included_hosts_id"`
    //[ 2] provisioned_networks_id                        UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    ProvisionedNetworksID null.String `gorm:"column:provisioned_networks_id;type:UUID;" json:"provisioned_networks_id" protobuf:"string,2,opt,name=provisioned_networks_id"`
    //[ 3] state                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    State null.String `gorm:"column:state;type:TEXT;" json:"state" protobuf:"string,3,opt,name=state"`
    //[ 4] ip_address                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    IPAddress null.String `gorm:"column:ip_address;type:TEXT;" json:"ip_address" protobuf:"string,4,opt,name=ip_address"`
    //[ 5] conn_info                                      JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
    ConnInfo null.String `gorm:"column:conn_info;type:JSON;" json:"conn_info" protobuf:"string,5,opt,name=conn_info"`
    //[ 6] planned_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PlannedChecksum null.String `gorm:"column:planned_checksum;type:TEXT;" json:"planned_checksum" protobuf:"string,6,opt,name=planned_checksum"`
    //[ 7] current_checksum                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    CurrentChecksum null.String `gorm:"column:current_checksum;type:TEXT;" json:"current_checksum" protobuf:"string,7,opt,name=current_checksum"`
    //[ 8] previous_checksum                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    PreviousChecksum null.String `gorm:"column:previous_checksum;type:TEXT;" json:"previous_checksum" protobuf:"string,8,opt,name=previous_checksum"`

}
*/

// TableName sets the insert table name for this struct type
func (p *ProvisionedHosts) TableName() string {
	return "provisioned_hosts"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProvisionedHosts) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProvisionedHosts) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProvisionedHosts) Validate(action Action) error {
	return nil
}
