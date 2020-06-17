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

// GetAllRemoteFileDefinitions is a function to get a slice of record(s) from remote_file_definitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRemoteFileDefinitions(ctx context.Context, page, pagesize int64, order string) (remotefiledefinitions []*model.RemoteFileDefinitions, totalRows int, err error) {

	remotefiledefinitions = []*model.RemoteFileDefinitions{}

	remotefiledefinitionsOrm := DB.Model(&model.RemoteFileDefinitions{})
	remotefiledefinitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		remotefiledefinitionsOrm = remotefiledefinitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		remotefiledefinitionsOrm = remotefiledefinitionsOrm.Limit(pagesize)
	}

	if order != "" {
		remotefiledefinitionsOrm = remotefiledefinitionsOrm.Order(order)
	}

	if err = remotefiledefinitionsOrm.Find(&remotefiledefinitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return remotefiledefinitions, totalRows, nil
}

// GetRemoteFileDefinitions is a function to get a single record from the remote_file_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetRemoteFileDefinitions(ctx context.Context, argId string) (record *model.RemoteFileDefinitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRemoteFileDefinitions is a function to add a single record to remote_file_definitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddRemoteFileDefinitions(ctx context.Context, record *model.RemoteFileDefinitions) (result *model.RemoteFileDefinitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRemoteFileDefinitions is a function to update a single record from remote_file_definitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRemoteFileDefinitions(ctx context.Context, argId string, updated *model.RemoteFileDefinitions) (result *model.RemoteFileDefinitions, RowsAffected int64, err error) {

	result = &model.RemoteFileDefinitions{}
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

// DeleteRemoteFileDefinitions is a function to delete a single record from remote_file_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRemoteFileDefinitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.RemoteFileDefinitions{}
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
