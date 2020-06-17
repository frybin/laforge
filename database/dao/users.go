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

// GetAllUsers is a function to get a slice of record(s) from users table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUsers(ctx context.Context, page, pagesize int64, order string) (users []*model.Users, totalRows int, err error) {

	users = []*model.Users{}

	usersOrm := DB.Model(&model.Users{})
	usersOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		usersOrm = usersOrm.Offset(offset).Limit(pagesize)
	} else {
		usersOrm = usersOrm.Limit(pagesize)
	}

	if order != "" {
		usersOrm = usersOrm.Order(order)
	}

	if err = usersOrm.Find(&users).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return users, totalRows, nil
}

// GetUsers is a function to get a single record from the users table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetUsers(ctx context.Context, argId string) (record *model.Users, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUsers is a function to add a single record to users table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddUsers(ctx context.Context, record *model.Users) (result *model.Users, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUsers is a function to update a single record from users table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUsers(ctx context.Context, argId string, updated *model.Users) (result *model.Users, RowsAffected int64, err error) {

	result = &model.Users{}
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

// DeleteUsers is a function to delete a single record from users table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUsers(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.Users{}
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
