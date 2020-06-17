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

// GetAllIdentityDefinitions is a function to get a slice of record(s) from identity_definitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIdentityDefinitions(ctx context.Context, page, pagesize int64, order string) (identitydefinitions []*model.IdentityDefinitions, totalRows int, err error) {

	identitydefinitions = []*model.IdentityDefinitions{}

	identitydefinitionsOrm := DB.Model(&model.IdentityDefinitions{})
	identitydefinitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		identitydefinitionsOrm = identitydefinitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		identitydefinitionsOrm = identitydefinitionsOrm.Limit(pagesize)
	}

	if order != "" {
		identitydefinitionsOrm = identitydefinitionsOrm.Order(order)
	}

	if err = identitydefinitionsOrm.Find(&identitydefinitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return identitydefinitions, totalRows, nil
}

// GetIdentityDefinitions is a function to get a single record from the identity_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetIdentityDefinitions(ctx context.Context, argId string) (record *model.IdentityDefinitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIdentityDefinitions is a function to add a single record to identity_definitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddIdentityDefinitions(ctx context.Context, record *model.IdentityDefinitions) (result *model.IdentityDefinitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIdentityDefinitions is a function to update a single record from identity_definitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIdentityDefinitions(ctx context.Context, argId string, updated *model.IdentityDefinitions) (result *model.IdentityDefinitions, RowsAffected int64, err error) {

	result = &model.IdentityDefinitions{}
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

// DeleteIdentityDefinitions is a function to delete a single record from identity_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIdentityDefinitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.IdentityDefinitions{}
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
