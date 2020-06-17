package dao

import (
	"context"
	"time"

	"github.com/frybin/laforge/database/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllProvisionedSteps is a function to get a slice of record(s) from provisioned_steps table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProvisionedSteps(ctx context.Context, page, pagesize int64, order string) (provisionedsteps []*model.ProvisionedSteps, totalRows int, err error) {

	provisionedsteps = []*model.ProvisionedSteps{}

	provisionedstepsOrm := DB.Model(&model.ProvisionedSteps{})
	provisionedstepsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		provisionedstepsOrm = provisionedstepsOrm.Offset(offset).Limit(pagesize)
	} else {
		provisionedstepsOrm = provisionedstepsOrm.Limit(pagesize)
	}

	if order != "" {
		provisionedstepsOrm = provisionedstepsOrm.Order(order)
	}

	if err = provisionedstepsOrm.Find(&provisionedsteps).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return provisionedsteps, totalRows, nil
}

// GetProvisionedSteps is a function to get a single record from the provisioned_steps table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetProvisionedSteps(ctx context.Context, argId string) (record *model.ProvisionedSteps, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProvisionedSteps is a function to add a single record to provisioned_steps table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddProvisionedSteps(ctx context.Context, record *model.ProvisionedSteps) (result *model.ProvisionedSteps, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProvisionedSteps is a function to update a single record from provisioned_steps table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProvisionedSteps(ctx context.Context, argId string, updated *model.ProvisionedSteps) (result *model.ProvisionedSteps, RowsAffected int64, err error) {

	result = &model.ProvisionedSteps{}
	db := DB.First(result, argId)
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

// DeleteProvisionedSteps is a function to delete a single record from provisioned_steps table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProvisionedSteps(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.ProvisionedSteps{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
