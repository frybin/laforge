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


Table: competitions
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] name                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "ehHmVcRVU^Ox_ZLydHobaEZCR",    "name": "u^NYxiLywpu^GEc]pstPSoYeG"}



*/

// Competitions struct is a row record of the competitions table in the laforge-2 database
/*
type Competitions struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] name                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    Name string `gorm:"column:name;type:TEXT;" json:"name" protobuf:"string,1,opt,name=name"`

}
*/

// TableName sets the insert table name for this struct type
func (c *Competitions) TableName() string {
	return "competitions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Competitions) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Competitions) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Competitions) Validate(action Action) error {
	return nil
}
