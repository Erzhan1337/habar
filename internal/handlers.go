package internal

import (
	"beverage-classifier/auto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandlers(r *gin.Engine, database *gorm.DB) {
	authHandler := &auto.Handler{DB: database}
	r.POST("/auth/signup", authHandler.SignUp)
	r.POST("/auth/signin", authHandler.SignIn)

	api := r.Group("/api")
	api.Use(auto.AuthMiddleware())

	api.GET("/beverages", func(c *gin.Context) {
		var beverages []Beverage
		if err := database.Preload("Ingredients").Preload("Nutrition").Preload("Flags").Find(&beverages).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve data"})
			return
		}
		c.JSON(http.StatusOK, beverages)
	})

	catalogWrite := api.Group("")
	catalogWrite.Use(auto.AuthorizeRole(auto.ADMIN, auto.MODERATOR))
	catalogWrite.POST("/beverage", func(c *gin.Context) {
		var b Beverage
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data not correct: " + err.Error()})
			return
		}
		if err := b.NormalizeAndValidate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		b.Flags = AnalyzeBeverage(&b)
		if err := database.Create(&b).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save beverage"})
			return
		}
		c.JSON(http.StatusCreated, b)
	})

	catalogWrite.PATCH("/beverage/:id", func(c *gin.Context) {
		id := c.Param("id")
		var b Beverage
		if err := database.Preload("Nutrition").Preload("Ingredients").First(&b, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "beverage not found"})
			return
		}

		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := b.NormalizeAndValidate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&b).Error; err != nil {
				return err
			}

			if err := tx.Where("product_id = ?", b.ProductID).Delete(&Flag{}).Error; err != nil {
				return err
			}

			flags := AnalyzeBeverage(&b)
			for i := range flags {
				flags[i].ProductID = b.ProductID
			}
			if len(flags) > 0 {
				if err := tx.Create(&flags).Error; err != nil {
					return err
				}
			}
			b.Flags = flags
			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update beverage: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, b)
	})

	adminOnly := api.Group("")
	adminOnly.Use(auto.AuthorizeRole(auto.ADMIN))
	adminOnly.DELETE("/beverage/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := database.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("product_id = ?", id).Delete(&Flag{}).Error; err != nil {
				return err
			}
			if err := tx.Where("product_id = ?", id).Delete(&Ingredient{}).Error; err != nil {
				return err
			}
			if err := tx.Where("product_id = ?", id).Delete(&Nutrition{}).Error; err != nil {
				return err
			}
			if err := tx.Delete(&Beverage{}, id).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete: " + err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	analytics := api.Group("/analytics")
	analytics.Use(auto.AuthorizeRole(auto.ADMIN, auto.MODERATOR, auto.ANALYST))
	analytics.GET("/market-stats", func(c *gin.Context) {
		var stats struct {
			SugarPct         float64 `json:"sugar_pct"`
			SweetenerPct     float64 `json:"sweetener_pct"`
			CaffeinePct      float64 `json:"caffeine_pct"`
			HeavyAdditivePct float64 `json:"heavy_additive_pct"`
		}
		if err := database.Raw(QueryMarketStats).Scan(&stats).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, stats)
	})

	analytics.GET("/category-distribution", func(c *gin.Context) {
		var rows []struct {
			CategoryLvl1  string `json:"category_lvl_1"`
			BeverageCount int    `json:"beverage_count"`
		}
		if err := database.Raw(QueryCategoryDistribution).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/additives-heatmap", func(c *gin.Context) {
		var rows []struct {
			CategoryLvl1     string `json:"category_lvl_1"`
			AdditiveCategory string `json:"additive_category"`
			UsageCount       int    `json:"usage_count"`
		}
		if err := database.Raw(QueryAdditivesHeatmap).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/top-e-additives", func(c *gin.Context) {
		var rows []struct {
			ENumber        string `json:"e_number"`
			IngredientName string `json:"ingredient_name"`
			UsageCount     int    `json:"usage_count"`
		}
		if err := database.Raw(QueryETopChart).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/sweetener-matrix", func(c *gin.Context) {
		var rows []struct {
			CategoryLvl1  string `json:"category_lvl_1"`
			Sweetener     string `json:"sweetener"`
			BeverageCount int    `json:"beverage_count"`
		}
		if err := database.Raw(QuerySweetenerMatrix).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/sugar-distribution", func(c *gin.Context) {
		var rows []struct {
			SugarSegment  string `json:"sugar_segment"`
			BeverageCount int    `json:"beverage_count"`
		}
		if err := database.Raw(QuerySugarDistribution).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/clean-label-share", func(c *gin.Context) {
		var row struct {
			CleanLabelPct float64 `json:"clean_label_pct"`
		}
		if err := database.Raw(QueryCleanLabelShare).Scan(&row).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, row)
	})

	analytics.GET("/composition-classification", func(c *gin.Context) {
		var rows []struct {
			ProductID          uint   `json:"product_id"`
			Name               string `json:"name"`
			CategoryLvl1       string `json:"category_lvl_1"`
			BaseType           string `json:"base_type"`
			CaffeineType       string `json:"caffeine_type"`
			Carbonation        string `json:"carbonation"`
			Sweeteners         string `json:"sweeteners"`
			AdditiveCategories string `json:"additive_categories"`
			Allergens          string `json:"allergens"`
		}
		if err := database.Raw(QueryCompositionClassification).Scan(&rows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		c.JSON(http.StatusOK, rows)
	})

	analytics.GET("/report", func(c *gin.Context) {
		var marketStats struct {
			SugarPct         float64 `json:"sugar_pct"`
			SweetenerPct     float64 `json:"sweetener_pct"`
			CaffeinePct      float64 `json:"caffeine_pct"`
			HeavyAdditivePct float64 `json:"heavy_additive_pct"`
		}
		var categoryDistribution []struct {
			CategoryLvl1  string `json:"category_lvl_1"`
			BeverageCount int    `json:"beverage_count"`
		}
		var additivesHeatmap []struct {
			CategoryLvl1     string `json:"category_lvl_1"`
			AdditiveCategory string `json:"additive_category"`
			UsageCount       int    `json:"usage_count"`
		}
		var topEAdditives []struct {
			ENumber        string `json:"e_number"`
			IngredientName string `json:"ingredient_name"`
			UsageCount     int    `json:"usage_count"`
		}
		var sweetenerMatrix []struct {
			CategoryLvl1  string `json:"category_lvl_1"`
			Sweetener     string `json:"sweetener"`
			BeverageCount int    `json:"beverage_count"`
		}
		var sugarDistribution []struct {
			SugarSegment  string `json:"sugar_segment"`
			BeverageCount int    `json:"beverage_count"`
		}
		var cleanLabelShare struct {
			CleanLabelPct float64 `json:"clean_label_pct"`
		}

		if err := database.Raw(QueryMarketStats).Scan(&marketStats).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QueryCategoryDistribution).Scan(&categoryDistribution).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QueryAdditivesHeatmap).Scan(&additivesHeatmap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QueryETopChart).Scan(&topEAdditives).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QuerySweetenerMatrix).Scan(&sweetenerMatrix).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QuerySugarDistribution).Scan(&sugarDistribution).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}
		if err := database.Raw(QueryCleanLabelShare).Scan(&cleanLabelShare).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "analytics query failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"market_stats":          marketStats,
			"category_distribution": categoryDistribution,
			"additives_heatmap":     additivesHeatmap,
			"top_e_additives":       topEAdditives,
			"sweetener_matrix":      sweetenerMatrix,
			"sugar_distribution":    sugarDistribution,
			"clean_label_share":     cleanLabelShare,
		})
	})
}
