package database

import (
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewDatabase(config *config.Config) *Database {
	return &Database{
		Config: config,
	}
}

func (db *Database) Connect() {
	dsn := "host=" + db.Config.PostgresqlHost + " port=" + db.Config.PostgresqlPort + " user=" + db.Config.PostgresqlUser + " password=" + db.Config.PostgresqlPassword + " dbname=" + db.Config.PostgresqlDatabase + " sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Настройки логирования (соответствует console.log в Sequelize)
		Logger: logger.Default.LogMode(logger.Info),

		// Отключение использования транзакций по умолчанию (включено в Sequelize)
		SkipDefaultTransaction: true,

		// Стратегия именования (camelCase to snake_case, table names in plural form)
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // Таблицы во множественном числе
			NoLowerCase:   false, // Использование snake_case
		},

		// Внешние ключи создаются при миграции (как в Sequelize)
		DisableForeignKeyConstraintWhenMigrating: false,

		// Отключение подготовленных выражений (аналогично поведению по умолчанию в Sequelize)
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}

	// Настройки пула соединений
	sqlDB.SetMaxOpenConns(5)                   // Устанавливаем максимальное количество соединений в пуле
	sqlDB.SetMaxIdleConns(5)                   // Устанавливаем минимальное количество соединений в пуле
	sqlDB.SetConnMaxIdleTime(10 * time.Second) // Устанавливаем максимальное время простоя соединения
	sqlDB.SetConnMaxLifetime(60 * time.Second) // Устанавливаем максимальное время жизни соединения

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Connected to the database")

	db.DB = conn

	db.AutoMigrate()
}

func (db *Database) AutoMigrate() {
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during migration: ", err)
	}
	fmt.Println("Database migration completed")
}
