package controller

import (
	"assignment-2/database"
	"assignment-2/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error bind json request",
		})
		return
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error create order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, orderResult)
}

func (c Controller) GetOrders(ctx *gin.Context) {
	orders, err := c.db.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error get data",
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) UpdateOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	var newOrder model.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error bind json request",
		})
		return
	}

	orderResult, err, isFound := c.db.UpdateOrder(orderId, newOrder)
	if err != nil {
		if !isFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    "404",
				"message": err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error update order",
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func (c Controller) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	err, isFound := c.db.DeleteOrder(orderId)
	if err != nil {
		if !isFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    "404",
				"message": err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error delete order",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "Successfully delete order",
	})
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}
