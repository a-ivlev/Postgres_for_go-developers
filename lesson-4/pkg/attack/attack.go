package attack

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"postgre-dev-go/internal/postgresDB"
	"sync"
	"sync/atomic"
	"time"
)

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func Attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64
	attacker := func(stopAt time.Time) {
		for {
			//_, err := search.Search(ctx, dbpool, "alex", 5)
			//_, err := postgresDB.SearchClientByPhone(ctx, dbpool, "+7 411 923 8377")
			//_, err := postgresDB.SearchRentItemsByPhone(ctx, dbpool, "+7 411 923 8377")
			_, err := postgresDB.SearchRentItemsByPhone(ctx, dbpool, "")
			if err != nil {
				log.Fatal(err)
			}
			atomic.AddUint64(&queries, 1)
			if time.Now().After(stopAt) {
				return
			}
		}
	}
	var wg sync.WaitGroup
	wg.Add(threads)
	startAt := time.Now()
	stopAt := startAt.Add(duration)
	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}
	wg.Wait()
	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}
