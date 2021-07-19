package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"postgre-dev-go/configs"
	"postgre-dev-go/pkg/attack"
	"time"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfDB()

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	duration := time.Duration(10 * time.Second)
	threads := 1000
	fmt.Println("start attack")
	res := attack.Attack(ctx, duration, threads, dbpool)
	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)

}	