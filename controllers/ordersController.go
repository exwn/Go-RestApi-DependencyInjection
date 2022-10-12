package controllers

import (
	"Assignment_02/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *Controllers) GetOrders(ctx *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	c.masterDB.Preload("Items").Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"status": "Order not found",
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": "success",
			"result": orders,
			"count":  len(orders),
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *Controllers) CreateOrder(ctx *gin.Context) {

	var (
		order  models.Order
		result gin.H
	)

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.masterDB.Create(&order)
	result = gin.H{
		"status": "success",
		"result": order,
		// "customerName": order.Customer_name,
		// "orderedAt":    order.Ordered_at,
		// "items": [
		// 	"itemCode": item.Item_code,
		// ],
	}
	ctx.JSON(http.StatusCreated, result)

}

func (c *Controllers) UpdateOrder(ctx *gin.Context) {

	var (
		order  models.Order
		result gin.H
	)

	id := ctx.Param("id")
	c.masterDB.First(&order, id)

	if order.Order_id == 0 {
		result = gin.H{
			"status": "Order not found",
			"result": nil,
		}
		return
	}
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.masterDB.Save(&order)
	c.masterDB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	result = gin.H{
		"status": "success update",
		"result": order,
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *Controllers) DeleteOrder(ctx *gin.Context) {

	var (
		order  models.Order
		result gin.H
	)

	id := ctx.Param("id")
	c.masterDB.First(&order, id)

	if order.Order_id == 0 {
		result = gin.H{
			"status": "Order not found",
			"result": nil,
		}
		return
	}
	// if err := ctx.ShouldBindJSON(&order); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	c.masterDB.Delete(&order)
	c.masterDB.Model(&order).Association("Items").Clear()

	result = gin.H{
		"status": "success delete",
		"result": nil,
	}
	ctx.JSON(http.StatusOK, result)
}
