package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/chain-identity/walletGenerator/internal/generator"
	"github.com/chain-identity/walletGenerator/internal/progressBar"
	"github.com/chain-identity/walletGenerator/internal/repository"
	"github.com/chain-identity/walletGenerator/internal/wallet"
)

func init() {
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		if err := os.Mkdir("db", 0o750); err != nil {
			panic(err)
		}
	}
}

func main() {
	// Context with gracefully shutdown signal
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)
	defer stop()

	fmt.Println("===============ETH Wallet Generator===============")
	fmt.Println(" ")

	// Parse flags
	number := flag.Int("n", 100, "set number of generate times (not number of result wallet) (set number to 0 for Infinite loop ∞)")
	dbPath := flag.String("db", "", "set sqlite output name that will be created in /db)")
	csvPath := flag.String("csv", "../result.csv", "csv filename")
	concurrency := flag.Int("c", 1, "set concurrency value")
	flag.Parse()

	// Repository
	var repo repository.Repository
	switch {
	case *csvPath != "":
		repo = repository.NewFileRepository(*csvPath)
	case *dbPath != "":
		db, err := gorm.Open(sqlite.Open("../db/"+*dbPath), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}

		defer func() {
			db, _ := db.DB()
			db.Close()
		}()

		if err := db.AutoMigrate(&wallet.Wallet{}); err != nil {
			panic(err)
		}

		repo = repository.NewGormRepository(db, uint64(*concurrency))
	default:
		repo = repository.NewInMemoryRepository()
	}

	// Wallet generator
	walletGenerator := wallet.PrivateKeyGenerator()
	pb := progressBar.NewCompatibleProgressBar(*number)

	generator := generator.New(
		walletGenerator,
		repo,
		generator.Config{
			ProgressBar: pb,
			Concurrency: *concurrency,
			Number:      *number,
		},
	)

	go func() {
		<-ctx.Done()

		if err := generator.Shutdown(); err != nil {
			log.Printf("Generator Shutdown Error: %+v", err)
		}

		if err := repo.Close(); err != nil {
			log.Printf("walletRepo Close Error: %+v", err)
		}
	}()

	if err := generator.Start(); err != nil {
		log.Printf("Generator Error: %+v", err)
	}
}
