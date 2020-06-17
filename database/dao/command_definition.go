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

// GetAllCommandDefinition is a function to get a slice of record(s) from command_definition table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCommandDefinition(ctx context.Context, page, pagesize int64, order string) (commanddefinition []*model.CommandDefinition, totalRows int, err error) {

	commanddefinition = []*model.CommandDefinition{}

	commanddefinitionOrm := DB.Model(&model.CommandDefinition{})
	commanddefinitionOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		commanddefinitionOrm = commanddefinitionOrm.Offset(offset).Limit(pagesize)
	} else {
		commanddefinitionOrm = commanddefinitionOrm.Limit(pagesize)
	}

	if order != "" {
		commanddefinitionOrm = commanddefinitionOrm.Order(order)
	}

	if err = commanddefinitionOrm.Find(&commanddefinition).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return commanddefinition, totalRows, nil
}

// GetCommandDefinition is a function to get a single record from the command_definition table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetCommandDefinition(ctx context.Context, argId string) (record *model.CommandDefinition, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCommandDefinition is a function to add a single record to command_definition table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddCommandDefinition(ctx context.Context, record *model.CommandDefinition) (result *model.CommandDefinition, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCommandDefinition is a function to update a single record from command_definition table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCommandDefinition(ctx context.Context, argId string, updated *model.CommandDefinition) (result *model.CommandDefinition, RowsAffected int64, err error) {

	result = &model.CommandDefinition{}
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

// DeleteCommandDefinition is a function to delete a single record from command_definition table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCommandDefinition(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.CommandDefinition{}
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
