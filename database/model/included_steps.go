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


Table: included_steps
[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] included_hosts_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] step_definitions_id                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 3] step_definition_type                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] step_offset                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "ZbsJoApNi_dPMJKkxseydFDaO",    "included_hosts_id": "nltoF\\ouLdOksk^sqcXuT_jop",    "step_definitions_id": "kZKvNwLkgjVKFcDgebfRYPsJF",    "step_definition_type": "UWlKS[Zw`sH[wkIiSeVmlHjht",    "step_offset": 15}



*/

// IncludedSteps struct is a row record of the included_steps table in the laforge-2 database
/*
type IncludedSteps struct {
    //[ 0] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
    ID string `gorm:"primary_key;column:id;type:UUID;" json:"id" protobuf:"string,0,opt,name=id"`
    //[ 1] included_hosts_id                              UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    IncludedHostsID null.String `gorm:"column:included_hosts_id;type:UUID;" json:"included_hosts_id" protobuf:"string,1,opt,name=included_hosts_id"`
    //[ 2] step_definitions_id                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
    StepDefinitionsID null.String `gorm:"column:step_definitions_id;type:UUID;" json:"step_definitions_id" protobuf:"string,2,opt,name=step_definitions_id"`
    //[ 3] step_definition_type                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
    StepDefinitionType null.String `gorm:"column:step_definition_type;type:TEXT;" json:"step_definition_type" protobuf:"string,3,opt,name=step_definition_type"`
    //[ 4] step_offset                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
    StepOffset null.Int `gorm:"column:step_offset;type:INT8;" json:"step_offset" protobuf:"int32,4,opt,name=step_offset"`

}
*/

// TableName sets the insert table name for this struct type
func (i *IncludedSteps) TableName() string {
	return "included_steps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IncludedSteps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IncludedSteps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IncludedSteps) Validate(action Action) error {
	return nil
}
