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

// GetAllHostDefinitions is a function to get a slice of record(s) from host_definitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllHostDefinitions(ctx context.Context, page, pagesize int64, order string) (hostdefinitions []*model.HostDefinitions, totalRows int, err error) {

	hostdefinitions = []*model.HostDefinitions{}

	hostdefinitionsOrm := DB.Model(&model.HostDefinitions{})
	hostdefinitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		hostdefinitionsOrm = hostdefinitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		hostdefinitionsOrm = hostdefinitionsOrm.Limit(pagesize)
	}

	if order != "" {
		hostdefinitionsOrm = hostdefinitionsOrm.Order(order)
	}

	if err = hostdefinitionsOrm.Find(&hostdefinitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return hostdefinitions, totalRows, nil
}

// GetHostDefinitions is a function to get a single record from the host_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetHostDefinitions(ctx context.Context, argId string) (record *model.HostDefinitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddHostDefinitions is a function to add a single record to host_definitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddHostDefinitions(ctx context.Context, record *model.HostDefinitions) (result *model.HostDefinitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateHostDefinitions is a function to update a single record from host_definitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateHostDefinitions(ctx context.Context, argId string, updated *model.HostDefinitions) (result *model.HostDefinitions, RowsAffected int64, err error) {

	result = &model.HostDefinitions{}
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

// DeleteHostDefinitions is a function to delete a single record from host_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteHostDefinitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.HostDefinitions{}
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
