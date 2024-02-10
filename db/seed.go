package database

import (
	"context"
	"time"

	"github.com/stasundr/gqlgen-pg-todo-example/graph/model"
	"github.com/theckman/yacspin"
	"github.com/uptrace/bun"
)

func Seed(db *bun.DB) error {
	if err := createSchemas(db); err != nil {
		return err
	}

	if err := seedUsers(db); err != nil {
		return err
	}

	if err := seedTodos(db); err != nil {
		return err
	}

	return nil
}

func createSchemas(db *bun.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Schema",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	models := []interface{}{
		(*model.User)(nil),
		(*model.Todo)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			return err
		}
	}
	spinner.Stop()

	return nil
}

func seedUsers(db *bun.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Users ",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	users := []model.User{
		{
			Email:     "oshalygin@gmail.com",
			FirstName: "Oleg",
			LastName:  "Shalygin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "john.snow@gmail.com",
			FirstName: "John",
			LastName:  "Snow",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "baby.yoda@gmail.com",
			FirstName: "Baby",
			LastName:  "Yoda",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		_, err := db.NewSelect().Model(&user).Where("email = ?", user.Email).Limit(1).Exec(context.Background())
		if err != nil {
			_, err := db.NewInsert().Model(&user).Exec(context.Background())
			if err != nil {
				return err
			}
		}
	}
	spinner.Stop()

	return nil
}

func seedTodos(db *bun.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Todos ",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	todos := []model.Todo{
		{
			Name:       "kubectl all the things",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  1,
			UpdatedBy:  1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Name:       "install a k8s cluster inside of another k8s cluster",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  2,
			UpdatedBy:  2,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Name:       "inception",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  3,
			UpdatedBy:  3,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for _, todo := range todos {
		_, err := db.NewSelect().Model(&todo).Where("name = ?", todo.Name).Limit(1).Exec(context.Background())
		if err != nil {
			_, err := db.NewInsert().Model(&todo).Exec(context.Background())
			if err != nil {
				return err
			}
		}
	}
	spinner.Stop()

	return nil
}
