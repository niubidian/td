package middleware_test

import (
	"context"
	"time"

	"golang.org/x/time/rate"

	"github.com/gotd/td/middleware"
	"github.com/gotd/td/middleware/floodwait"
	"github.com/gotd/td/middleware/ratelimit"
	"github.com/gotd/td/telegram"
)

func Example() {
	// Create a new telegram.Client instance that handles FLOOD_WAIT errors
	// and limits request rate to 10 Hz with max burst size of 5 request.
	//
	// Note that you must not use test app credentials in production.
	// See https://core.telegram.org/api/obtaining_api_id
	//
	client := telegram.NewClient(
		telegram.TestAppID,
		telegram.TestAppHash,
		telegram.Options{
			Middleware: middleware.Chain(
				floodwait.Middleware(),
				ratelimit.Middleware(
					rate.NewLimiter(rate.Every(100*time.Millisecond), 5),
				),
			),
		},
	)

	ctx := context.TODO()
	err := client.Run(ctx, func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		panic(err)
	}
}