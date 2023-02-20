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


Table: fc_configuration_template
[ 0] fc_configuration_template_id                   INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] config_level                                   VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 2] setting                                        JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
[ 3] setting_secret                                 VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 4] is_enabled                                     BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
[ 5] create_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [CURRENT_TIMESTAMP]
[ 6] modify_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [CURRENT_TIMESTAMP]
[ 7] create_by                                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 8] modify_by                                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []


JSON Sample
-------------------------------------
{    "fc_configuration_template_id": 48,    "config_level": "oZDcUDCoUUcyylPKAtwXNFnRX",    "setting": "jsooYqRhFgWJybZURBpRMVeer",    "setting_secret": "meQTamLtNVSubfpsQrYjOhOEx",    "is_enabled": true,    "create_time": "2246-05-21T11:55:29.603113421+08:00",    "modify_time": "2103-02-15T02:31:24.512196174+08:00",    "create_by": "tZOiRtFgNhbKKYOgruQwvYwfr",    "modify_by": "HNpPaugsUWvgNMWRHfqoIKYBk"}



*/

// FcConfigurationTemplate struct is a row record of the fc_configuration_template table in the main database
type FcConfigurationTemplate struct {
	//[ 0] fc_configuration_template_id                   INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	FcConfigurationTemplateID int32 `gorm:"primary_key;column:fc_configuration_template_id;type:INT4;" json:"fc_configuration_template_id"`
	//[ 1] config_level                                   VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	ConfigLevel null.String `gorm:"column:config_level;type:VARCHAR;size:100;" json:"config_level"`
	//[ 2] setting                                        JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
	Setting null.String `gorm:"column:setting;type:JSONB;" json:"setting"`
	//[ 3] setting_secret                                 VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	SettingSecret null.String `gorm:"column:setting_secret;type:VARCHAR;" json:"setting_secret"`
	//[ 4] is_enabled                                     BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	IsEnabled null.Int `gorm:"column:is_enabled;type:BOOL;default:true;" json:"is_enabled"`
	//[ 5] create_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime null.Time `gorm:"column:create_time;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"create_time"`
	//[ 6] modify_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [CURRENT_TIMESTAMP]
	ModifyTime null.Time `gorm:"column:modify_time;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"modify_time"`
	//[ 7] create_by                                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	CreateBy null.String `gorm:"column:create_by;type:VARCHAR;size:100;" json:"create_by"`
	//[ 8] modify_by                                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	ModifyBy null.String `gorm:"column:modify_by;type:VARCHAR;size:100;" json:"modify_by"`
}


// TableName sets the insert table name for this struct type
func (f *FcConfigurationTemplate) TableName() string {
	return "fc_configuration_template"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FcConfigurationTemplate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FcConfigurationTemplate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FcConfigurationTemplate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FcConfigurationTemplate) TableInfo() *TableInfo {
	return fc_configuration_templateTableInfo
}
