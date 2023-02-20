package dao

import (
	"context"
	"time"

	"example.com/rest/example/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllFcConfigurationTemplate is a function to get a slice of record(s) from fc_configuration_template table in the main database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFcConfigurationTemplate(ctx context.Context, page, pagesize int64, order string) (results []*model.FcConfigurationTemplate, totalRows int, err error) {

	resultOrm := DB.Model(&model.FcConfigurationTemplate{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetFcConfigurationTemplate is a function to get a single record from the fc_configuration_template table in the main database
// error - ErrNotFound, db Find error
func GetFcConfigurationTemplate(ctx context.Context, argFcConfigurationTemplateID int32) (record *model.FcConfigurationTemplate, err error) {
	record = &model.FcConfigurationTemplate{}
	if err = DB.First(record, argFcConfigurationTemplateID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFcConfigurationTemplate is a function to add a single record to fc_configuration_template table in the main database
// error - ErrInsertFailed, db save call failed
func AddFcConfigurationTemplate(ctx context.Context, record *model.FcConfigurationTemplate) (result *model.FcConfigurationTemplate, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFcConfigurationTemplate is a function to update a single record from fc_configuration_template table in the main database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFcConfigurationTemplate(ctx context.Context, argFcConfigurationTemplateID int32, updated *model.FcConfigurationTemplate) (result *model.FcConfigurationTemplate, RowsAffected int64, err error) {

	result = &model.FcConfigurationTemplate{}
	db := DB.First(result, argFcConfigurationTemplateID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteFcConfigurationTemplate is a function to delete a single record from fc_configuration_template table in the main database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFcConfigurationTemplate(ctx context.Context, argFcConfigurationTemplateID int32) (rowsAffected int64, err error) {

	record := &model.FcConfigurationTemplate{}
	db := DB.First(record, argFcConfigurationTemplateID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
