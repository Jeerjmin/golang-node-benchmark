package api

import (
	"go-auth/api/user"
	database "go-auth/internal/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Создаем гистограмму для измерения времени выполнения запросов
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "url"}, // Метки для группировки по методу и маршруту
	)
)

func init() {
	// Регистрируем метрики в Prometheus
	prometheus.MustRegister(requestDuration)
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		// Записываем данные в метрику
		requestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
	}
}

func RegisterRoutes(db *database.Database) *gin.Engine {
	router := gin.Default()

	router.Use(prometheusMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	user.RegisterRoutes(router, db)

	return router
}
