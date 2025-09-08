package handlers

import (
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
//
//	@Param       body body models.CreateOrderRequest true "Order Request" example({
//		  "order": {
//		    "user_id": 1,
//		    "schedule_id": 10,
//		    "payment_id": 2,
//		    "fullname": "John Doe",
//		    "email": "johndoe@mail.com",
//		    "phone": "08123456789"
//		  },
//		  "seat_ids": [1, 2, 3]
//		})
//
// @Router      /orders [post]
func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	order, err := oh.orderRepo.CreateOrder(ctx, &req.Order, req.SeatIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": order})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid order id"})
		return
	}

	order, err := oh.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": order})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	orders, err := oh.orderRepo.GetOrdersByUserID(ctx, userID)
	if err != nil || len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no orders found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}
