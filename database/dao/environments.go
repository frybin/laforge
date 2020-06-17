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

// GetAllEnvironments is a function to get a slice of record(s) from environments table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllEnvironments(ctx context.Context, page, pagesize int64, order string) (environments []*model.Environments, totalRows int, err error) {

	environments = []*model.Environments{}

	environmentsOrm := DB.Model(&model.Environments{})
	environmentsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		environmentsOrm = environmentsOrm.Offset(offset).Limit(pagesize)
	} else {
		environmentsOrm = environmentsOrm.Limit(pagesize)
	}

	if order != "" {
		environmentsOrm = environmentsOrm.Order(order)
	}

	if err = environmentsOrm.Find(&environments).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return environments, totalRows, nil
}

// GetEnvironments is a function to get a single record from the environments table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetEnvironments(ctx context.Context, argId string) (record *model.Environments, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddEnvironments is a function to add a single record to environments table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddEnvironments(ctx context.Context, record *model.Environments) (result *model.Environments, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateEnvironments is a function to update a single record from environments table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateEnvironments(ctx context.Context, argId string, updated *model.Environments) (result *model.Environments, RowsAffected int64, err error) {

	result = &model.Environments{}
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

// DeleteEnvironments is a function to delete a single record from environments table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteEnvironments(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.Environments{}
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
