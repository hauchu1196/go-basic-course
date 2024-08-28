// internal/api/handlers/order_handler_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hauchu1196/ecombase/internal/api/dto"
	"github.com/hauchu1196/ecombase/internal/mocks"
	"github.com/hauchu1196/ecombase/internal/models"
	"github.com/hauchu1196/ecombase/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
    mockService := new(mocks.MockOrderService)
    handler := NewOrderHandler(mockService)

    order := &models.Order{
        CustomerName: "John Doe",
        TotalAmount:  100.0,
        Status:       "pending",
        LineItems: []models.LineItem{
            {ProductName: "Item 1", Quantity: 2, UnitPrice: 50.0},
        },
    }

    mockService.On("CreateOrder", mock.AnythingOfType("*models.Order")).Return(nil).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*models.Order)
        arg.ID = 1
        arg.CreatedAt = time.Now()
        arg.UpdatedAt = time.Now()
    })

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    jsonOrder, _ := json.Marshal(order)
    c.Request, _ = http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(jsonOrder))
    c.Request.Header.Set("Content-Type", "application/json")

    handler.CreateOrder(c)

    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response dto.OrderResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, uint(1), response.ID)
    assert.Equal(t, "John Doe", response.CustomerName)
    assert.Equal(t, 100.0, response.TotalAmount)
    assert.Equal(t, "pending", response.Status)
    assert.Len(t, response.LineItems, 1)

    mockService.AssertExpectations(t)
}

func TestGetOrder(t *testing.T) {
    mockService := new(mocks.MockOrderService)
    handler := NewOrderHandler(&service.OrderService{})

    order := &models.Order{
        Model:        gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
        CustomerName: "John Doe",
        TotalAmount:  100.0,
        Status:       "pending",
        LineItems: []models.LineItem{
            {Model: gorm.Model{ID: 1}, ProductName: "Item 1", Quantity: 2, UnitPrice: 50.0, TotalPrice: 100.0},
        },
    }

    mockService.On("GetOrder", uint(1)).Return(order, nil)

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{{Key: "id", Value: "1"}}

    handler.GetOrder(c)

    assert.Equal(t, http.StatusOK, w.Code)

    var response dto.OrderResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, uint(1), response.ID)
    assert.Equal(t, "John Doe", response.CustomerName)
    assert.Equal(t, 100.0, response.TotalAmount)
    assert.Equal(t, "pending", response.Status)
    assert.Len(t, response.LineItems, 1)

    mockService.AssertExpectations(t)
}

func TestListOrders(t *testing.T) {
    mockService := new(mocks.MockOrderService)
    handler := NewOrderHandler(mockService)

    orders := []models.Order{
        {
            Model:        gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
            CustomerName: "John Doe",
            TotalAmount:  100.0,
            Status:       "pending",
            LineItems: []models.LineItem{
                {Model: gorm.Model{ID: 1}, ProductName: "Item 1", Quantity: 2, UnitPrice: 50.0, TotalPrice: 100.0},
            },
        },
        {
            Model:        gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
            CustomerName: "Jane Doe",
            TotalAmount:  150.0,
            Status:       "completed",
            LineItems: []models.LineItem{
                {Model: gorm.Model{ID: 2}, ProductName: "Item 2", Quantity: 1, UnitPrice: 150.0, TotalPrice: 150.0},
            },
        },
    }

    mockService.On("ListOrders", 1, 10).Return(orders, 2, nil)

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodGet, "/orders?page=1&page_size=10", nil)

    handler.ListOrders(c)

    assert.Equal(t, http.StatusOK, w.Code)

    var response dto.OrderListResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Len(t, response.Orders, 2)
    assert.Equal(t, 2, response.Total)
    assert.Equal(t, "John Doe", response.Orders[0].CustomerName)
    assert.Equal(t, "Jane Doe", response.Orders[1].CustomerName)

    mockService.AssertExpectations(t)
}