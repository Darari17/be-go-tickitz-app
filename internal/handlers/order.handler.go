package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepo *repositories.OrderRepo
}

func NewOrderHandler(orderRepo *repositories.OrderRepo) *OrderHandler {
	return &OrderHandler{orderRepo: orderRepo}
}

// CreateOrder godoc
// @Summary     Create a new Order
// @Description Buat pesanan baru
// @Tags        Orders
// @Security    BearerToken
// @Accept      json
// @Produce     json
// @Param       body body models.CreateOrderExample true "Order Request"
// @Router      /orders [post]
func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	if len(req.SeatIDs) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "at least one seat must be selected",
		})
		return
	}

	newOrder, err := oh.orderRepo.CreateOrder(ctx.Request.Context(), &req.Order, req.SeatIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	orderWithSeats, err := oh.orderRepo.GetOrderByID(ctx.Request.Context(), newOrder.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "failed to fetch created order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   orderWithSeats,
	})
}

// GetOrderByID godoc
// @Summary     Get Order Detail
// @Description Detail pesanan berdasarkan Order ID
// @Tags        Orders
// @Security    BearerToken
// @Produce     json
// @Param       id path int true "Order ID"
// @Router      /orders/{id} [get]
func (oh *OrderHandler) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid order id",
		})
		return
	}

	order, err := oh.orderRepo.GetOrderByID(ctx.Request.Context(), id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   order,
	})
}

// GetOrdersByUser godoc
// @Summary     Get Order History by User
// @Description Semua order milik user berdasarkan User ID
// @Tags        Orders
// @Security    BearerToken
// @Produce     json
// @Param       user_id path int true "User ID"
// @Router      /orders/user/{user_id} [get]
func (oh *OrderHandler) GetOrdersByUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid user id",
		})
		return
	}

	orders, err := oh.orderRepo.GetOrdersByUserID(ctx.Request.Context(), userID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "failed to fetch orders",
		})
		return
	}

	if len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "no orders found for this user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}
