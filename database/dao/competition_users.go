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

// GetAllCompetitionUsers is a function to get a slice of record(s) from competition_users table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCompetitionUsers(ctx context.Context, page, pagesize int64, order string) (competitionusers []*model.CompetitionUsers, totalRows int, err error) {

	competitionusers = []*model.CompetitionUsers{}

	competitionusersOrm := DB.Model(&model.CompetitionUsers{})
	competitionusersOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		competitionusersOrm = competitionusersOrm.Offset(offset).Limit(pagesize)
	} else {
		competitionusersOrm = competitionusersOrm.Limit(pagesize)
	}

	if order != "" {
		competitionusersOrm = competitionusersOrm.Order(order)
	}

	if err = competitionusersOrm.Find(&competitionusers).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return competitionusers, totalRows, nil
}

// GetCompetitionUsers is a function to get a single record from the competition_users table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetCompetitionUsers(ctx context.Context, argId string) (record *model.CompetitionUsers, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCompetitionUsers is a function to add a single record to competition_users table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddCompetitionUsers(ctx context.Context, record *model.CompetitionUsers) (result *model.CompetitionUsers, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCompetitionUsers is a function to update a single record from competition_users table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCompetitionUsers(ctx context.Context, argId string, updated *model.CompetitionUsers) (result *model.CompetitionUsers, RowsAffected int64, err error) {

	result = &model.CompetitionUsers{}
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

// DeleteCompetitionUsers is a function to delete a single record from competition_users table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCompetitionUsers(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.CompetitionUsers{}
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
