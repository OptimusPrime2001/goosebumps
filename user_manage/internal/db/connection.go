package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"user-manage-backend/internal/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var DB *sql.DB
var DB *gorm.DB

func InitDb() error {
	return Database_GORM()
}
func Database_GORM() error {
	connStr := configs.NewConfig().GetDSN()
	fmt.Println("connStr:", connStr)
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	var errs error
	DB, errs = gorm.Open(postgres.New(postgres.Config{
		DSN: connStr,
	}), config)
	if errs != nil {
		return fmt.Errorf("failed to connect to database: %w", errs)
	}

	sqlDB, errs := DB.DB()
	if errs != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", errs)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if errs := sqlDB.PingContext(ctx); errs != nil {
		sqlDB.Close()
		return fmt.Errorf("failed to ping database: %w", errs)
	}
	log.Println("Connect postgres database using gorm success✅✅✅")
	return nil
}
func Database_SQL() error {
	var DB *sql.DB
	connStr := configs.NewConfig().GetDSN()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	DB.SetMaxIdleConns(3)                   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.SetMaxOpenConns(30)                  // SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.SetConnMaxLifetime(30 * time.Minute) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	DB.SetConnMaxIdleTime(5 * time.Minute)  // SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := DB.PingContext(ctx); err != nil {
		DB.Close()
		log.Fatalf("Failed to ping database: %v", err)
	}
	return nil
}
