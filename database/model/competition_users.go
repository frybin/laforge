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


Table: competition_users
[ 0] id                                             UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] users_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "keAX\\YoOqjGA[wP^KmAHd]wJY",    "users_id": "]iQVSeYZKvrNRSvZHsnxy]VJ`",    "competitions_id": "cjOloGdn`wShNBa`gIQR]LgSf"}


Comments
-------------------------------------
[ 0] Warning table: competition_users does not have a primary key defined, setting col position 1 id as primary key
Warning table: competition_users primary key column id is nullable column, setting it as NOT NULL




*/

// CompetitionUsers struct is a row record of the competition_users table in the laforge-2 database
/*
type CompetitionUsers struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] users_id                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    UsersID null.String `gorm:"column:users_id;type:UUID;" json:"users_id" protobuf:"string,1,opt,name=users_id"`
    //[ 2] competitions_id                                UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    CompetitionsID null.String `gorm:"column:competitions_id;type:UUID;" json:"competitions_id" protobuf:"string,2,opt,name=competitions_id"`

}
*/

// TableName sets the insert table name for this struct type
func (c *CompetitionUsers) TableName() string {
	return "competition_users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CompetitionUsers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CompetitionUsers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CompetitionUsers) Validate(action Action) error {
	return nil
}
