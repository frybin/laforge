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

// GetAllStepDefinitions is a function to get a slice of record(s) from step_definitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllStepDefinitions(ctx context.Context, page, pagesize int64, order string) (stepdefinitions []*model.StepDefinitions, totalRows int, err error) {

	stepdefinitions = []*model.StepDefinitions{}

	stepdefinitionsOrm := DB.Model(&model.StepDefinitions{})
	stepdefinitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		stepdefinitionsOrm = stepdefinitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		stepdefinitionsOrm = stepdefinitionsOrm.Limit(pagesize)
	}

	if order != "" {
		stepdefinitionsOrm = stepdefinitionsOrm.Order(order)
	}

	if err = stepdefinitionsOrm.Find(&stepdefinitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return stepdefinitions, totalRows, nil
}

// GetStepDefinitions is a function to get a single record from the step_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetStepDefinitions(ctx context.Context, argId string) (record *model.StepDefinitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddStepDefinitions is a function to add a single record to step_definitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddStepDefinitions(ctx context.Context, record *model.StepDefinitions) (result *model.StepDefinitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateStepDefinitions is a function to update a single record from step_definitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateStepDefinitions(ctx context.Context, argId string, updated *model.StepDefinitions) (result *model.StepDefinitions, RowsAffected int64, err error) {

	result = &model.StepDefinitions{}
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

// DeleteStepDefinitions is a function to delete a single record from step_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteStepDefinitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.StepDefinitions{}
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
