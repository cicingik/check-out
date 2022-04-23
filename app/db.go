package app

import (
	"context"
	"net/http"

	"github.com/cicingik/check-out/repository/postgre"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jinzhu/gorm"
)

func ContextualizeDb(db *postgre.DbEngine) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if db == nil || db.G == nil {
				next.ServeHTTP(w, r) // noop
				return
			}

			ctx := r.Context()
			dbConnectionWithContext := gormtrace.WithContext(ctx, db.G)
			newContext := context.WithValue(ctx, "db", dbConnectionWithContext)
			next.ServeHTTP(w, r.WithContext(newContext))
		})
	}
}
