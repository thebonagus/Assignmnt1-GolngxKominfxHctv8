package database

import (
	"assignment-2/config"
	"assignment-2/model"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	dbInfo := config.GetDatabaseEnv()

	var config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("Error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(&model.Order{}, &model.Item{})
	if err != nil {
		fmt.Println("Error on migration")
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}

func (d Database) CreateOrder(order model.Order) (model.Order, error) {
	dbg := d.db.Create(&order)
	if err := dbg.Error; err != nil {
		return model.Order{}, err
	}

	newOrder := model.Order{
		OrderID:      order.OrderID,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}

	return newOrder, nil
}

func (d Database) GetOrders() ([]model.Order, error) {
	var orders []model.Order

	err := d.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// UpdateOrder return Order result, error, and boolean isFound
func (d Database) UpdateOrder(id int, data model.Order) (model.Order, error, bool) {
	res := d.db.Where("order_id", id).Omit("Items").Updates(&data)

	if res.RowsAffected == 0 && res.Error == nil {
		return model.Order{}, errors.New("order not found"), false
	}

	if res.Error != nil {
		return model.Order{}, res.Error, true
	}

	order := model.Order{OrderID: id}
	for _, v := range data.Items {
		if v.ItemID != 0 {
			d.db.Where("item_id", v.ItemID).Updates(&v)
		}
	}
	err := d.db.Model(&order).Association("Items").Replace(data.Items)

	if err != nil {
		return model.Order{}, err, true
	}

	updatedOrder := model.Order{
		OrderID:      id,
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		Items:        data.Items,
	}

	return updatedOrder, nil, true
}

// DeleteOrder return Order error, and boolean isFound
func (d Database) DeleteOrder(id int) (error, bool) {
	res := d.db.Delete(&model.Order{}, "order_id", id)

	if res.RowsAffected == 0 && res.Error == nil {
		return errors.New("order not found"), false
	}

	return res.Error, true
}
