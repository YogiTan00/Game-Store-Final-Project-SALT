package item

import (
	"context"
	"database/sql"
	"fmt"
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type ItemRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewItemMysqlInteractor(conndb *sql.DB) repository.ItemRepository {
	return &ItemRepositoryMysqlInteractor{dbConn: conndb}
}

func (i *ItemRepositoryMysqlInteractor) GetAllItem(ctx context.Context) ([]*item.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s `, model.GetTableItem())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.ItemModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, i.dbConn, stmt, opts)
	if result == nil {
		return nil, nil
	}

	listItem, errMap := mapper.ListModelToDomainItem(result.([]*model.ItemModel))
	if errMap != nil {
		return nil, errMap
	}

	return listItem, nil
}

func (i *ItemRepositoryMysqlInteractor) GetItemByID(ctx context.Context, id string) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id =?`, model.GetTableItem())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.ItemModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, i.dbConn, stmt, opts, id)
	if result == nil {
		return nil, nil
	}

	dataItem, errMap := mapper.ModelToDomainItem(result.(*model.ItemModel))
	if errMap != nil {
		return nil, errMap
	}

	return dataItem, nil
}
