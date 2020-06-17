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

// GetAllRoles is a function to get a slice of record(s) from roles table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRoles(ctx context.Context, page, pagesize int64, order string) (roles []*model.Roles, totalRows int, err error) {

	roles = []*model.Roles{}

	rolesOrm := DB.Model(&model.Roles{})
	rolesOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		rolesOrm = rolesOrm.Offset(offset).Limit(pagesize)
	} else {
		rolesOrm = rolesOrm.Limit(pagesize)
	}

	if order != "" {
		rolesOrm = rolesOrm.Order(order)
	}

	if err = rolesOrm.Find(&roles).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return roles, totalRows, nil
}

// GetRoles is a function to get a single record from the roles table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetRoles(ctx context.Context, argId string) (record *model.Roles, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRoles is a function to add a single record to roles table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddRoles(ctx context.Context, record *model.Roles) (result *model.Roles, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRoles is a function to update a single record from roles table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRoles(ctx context.Context, argId string, updated *model.Roles) (result *model.Roles, RowsAffected int64, err error) {

	result = &model.Roles{}
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

// DeleteRoles is a function to delete a single record from roles table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRoles(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.Roles{}
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
