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


Table: host_definitions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] hostname                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] os                                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "hostname": "]aEvCmCjgif`bRAZ[sttmhUUi",    "os": "jqXLaqxQeNwIkUwjuIsrASgKZ",    "id": "qVmkWoUWWIfepXgJEf[nlpHPb"}



*/

// HostDefinitions struct is a row record of the host_definitions table in the laforge-2 database
/*
type HostDefinitions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] hostname                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Hostname null.String `gorm:"column:hostname;type:TEXT;" json:"hostname" protobuf:"string,1,opt,name=hostname"`
    //[ 2] os                                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Os null.String `gorm:"column:os;type:TEXT;" json:"os" protobuf:"string,2,opt,name=os"`

}
*/

// TableName sets the insert table name for this struct type
func (h *HostDefinitions) TableName() string {
	return "host_definitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HostDefinitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HostDefinitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HostDefinitions) Validate(action Action) error {
	return nil
}
