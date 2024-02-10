package dataloaders

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/stasundr/gqlgen-pg-todo-example/graph/model"
	"github.com/uptrace/bun"
	"github.com/vikstrous/dataloadgen"
)

func User(db *bun.DB, w http.ResponseWriter, r *http.Request, next http.Handler) {
	fetchFn := func(ctx context.Context, keys []int) ([]*model.User, []error) {
		var dbUsers []*model.User
		// Using Bun, we adjust the query method to fit Bun's API.
		// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
		// query method calls for Bun.
		err := db.NewSelect().Model(&dbUsers).Where("id IN (?)", bun.In(keys)).Scan(ctx)

		if err != nil {
			return []*model.User{}, []error{err}
		}

		log.Println("DBUSERS !!!!", dbUsers, len(keys), keys)
		// Mapping user IDs to users for quick lookup.
		userMap := make(map[int]*model.User)
		for _, user := range dbUsers {
			userMap[user.ID] = user
		}

		// Reassembling the results in the order of keys.
		users := make([]*model.User, len(keys))
		var errs []error
		for i, key := range keys {
			if user, ok := userMap[key]; ok {
				users[i] = user
			} else {
				// Handle the case where a key does not have a corresponding user.
				errs = append(errs, fmt.Errorf("no user found for key: %d", key))
				users[i] = nil // Keep place with nil if user is not found.
			}
		}

		return users, errs
	}

	loader := dataloadgen.NewLoader(fetchFn)

	ctx := context.WithValue(r.Context(), UserLoader, loader)
	next.ServeHTTP(w, r.WithContext(ctx))
}
