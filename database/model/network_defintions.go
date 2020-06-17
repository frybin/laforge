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


Table: network_defintions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] domain                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "bwROlhoSIem[FU^T_[rFoLRKN",    "domain": "o^gpJxcZ[OIJv[lUinSXaysxc"}



*/

// NetworkDefintions struct is a row record of the network_defintions table in the laforge-2 database
/*
type NetworkDefintions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] domain                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Domain null.String `gorm:"column:domain;type:TEXT;" json:"domain" protobuf:"string,1,opt,name=domain"`

}
*/

// TableName sets the insert table name for this struct type
func (n *NetworkDefintions) TableName() string {
	return "network_defintions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NetworkDefintions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NetworkDefintions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NetworkDefintions) Validate(action Action) error {
	return nil
}
