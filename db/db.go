package db

import (
	"context"
	"fmt"
	"goravel/app/facades"
)

func DelayedQuery(ctx context.Context, framework string) {
	facades.Orm().WithContext(ctx).Query().Exec("WAITFOR DELAY '0:00:10'")
	fmt.Printf("\n[%s] Wait DB call done\n\n", framework)
}
