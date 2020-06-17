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

// GetAllProvisionedTeams is a function to get a slice of record(s) from provisioned_teams table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProvisionedTeams(ctx context.Context, page, pagesize int64, order string) (provisionedteams []*model.ProvisionedTeams, totalRows int, err error) {

	provisionedteams = []*model.ProvisionedTeams{}

	provisionedteamsOrm := DB.Model(&model.ProvisionedTeams{})
	provisionedteamsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		provisionedteamsOrm = provisionedteamsOrm.Offset(offset).Limit(pagesize)
	} else {
		provisionedteamsOrm = provisionedteamsOrm.Limit(pagesize)
	}

	if order != "" {
		provisionedteamsOrm = provisionedteamsOrm.Order(order)
	}

	if err = provisionedteamsOrm.Find(&provisionedteams).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return provisionedteams, totalRows, nil
}

// GetProvisionedTeams is a function to get a single record from the provisioned_teams table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetProvisionedTeams(ctx context.Context, argId string) (record *model.ProvisionedTeams, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProvisionedTeams is a function to add a single record to provisioned_teams table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddProvisionedTeams(ctx context.Context, record *model.ProvisionedTeams) (result *model.ProvisionedTeams, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProvisionedTeams is a function to update a single record from provisioned_teams table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProvisionedTeams(ctx context.Context, argId string, updated *model.ProvisionedTeams) (result *model.ProvisionedTeams, RowsAffected int64, err error) {

	result = &model.ProvisionedTeams{}
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

// DeleteProvisionedTeams is a function to delete a single record from provisioned_teams table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProvisionedTeams(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.ProvisionedTeams{}
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
