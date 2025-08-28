package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/korshunvladislav/testTaskEM/initializers"
	"github.com/korshunvladislav/testTaskEM/models"
	"github.com/korshunvladislav/testTaskEM/utils"
)

func SubscriptionsCreate(c *gin.Context) {
	var body struct {
		ServiceName string          `json:"service_name"`
		Price       int             `json:"price"`
		UserID      uuid.UUID       `json:"user_id"`
		StartDate   utils.MonthYear `json:"start_date"`
		EndDate     utils.MonthYear `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	subscription := models.Subscription{
		ServiceName: body.ServiceName,
		Price:       body.Price,
		UserID:      body.UserID,
		StartDate:   body.StartDate.Time,
		EndDate:     body.EndDate.Time,
	}

	result := initializers.DB.Create(&subscription)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"subscription": subscription,
	})
}

func SubscriptionsIndex(c *gin.Context) {
	var subscriptions []models.Subscription
	initializers.DB.Find(&subscriptions)

	c.JSON(200, gin.H{
		"subscriptions": subscriptions,
	})
}

func SubscriptionsShow(c *gin.Context) {
	id := c.Param("id")

	var subscription models.Subscription
	initializers.DB.First(&subscription, id)

	c.JSON(200, gin.H{
		"subscription": subscription,
	})
}

func SubscriptionsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		ServiceName string          `json:"service_name"`
		Price       int             `json:"price"`
		UserID      uuid.UUID       `json:"user_id"`
		StartDate   utils.MonthYear `json:"start_date"`
		EndDate     utils.MonthYear `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var subscription models.Subscription
	initializers.DB.First(&subscription, id)

	initializers.DB.Model(&subscription).Updates(models.Subscription{
		ServiceName: body.ServiceName,
		Price:       body.Price,
		UserID:      body.UserID,
		StartDate:   body.StartDate.Time,
		EndDate:     body.EndDate.Time,
	})

	c.JSON(200, gin.H{
		"subscription": subscription,
	})
}

func SubscriptionsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Subscription{}, id)

	c.Status(200)
}

func SubscriptionsSummary(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	userIdStr := c.Query("user_id")
	serviceName := c.Query("service_name")

	query := initializers.DB.Model(&models.Subscription{})

	if userIdStr != "" {
		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
		query = query.Where("user_id = ?", userId)
	}

	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	if startDateStr != "" {
		if t, err := time.Parse(utils.DateLayout, startDateStr); err == nil {
			query = query.Where("start_date >= ?", t)
		}
	}

	if endDateStr != "" {
		if t, err := time.Parse(utils.DateLayout, endDateStr); err == nil {
			query = query.Where("end_date <= ?", t)
		}
	}

	var total int64
	query.Select("SUM(price)").Scan(&total)

	c.JSON(200, gin.H{
		"total_price": total,
	})
}
