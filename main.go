package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-prisma-bug/db"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	// Set Timezone
	location, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		panic(err)
	}
	time.Local = location

	fmt.Printf("check timezone: %s\n", time.Now())

	// create a user
	firstUser, err := client.User.CreateOne(
		db.User.Name.Set("FirstUser"),
		db.User.LastLogin.Set(time.Now()),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(firstUser, "", "  ")
	fmt.Printf("created first user: %s\n", result)

	// second user
	secondUser, err := client.User.CreateOne(
		db.User.Name.Set("SecondUser"),
		db.User.LastLogin.Set(time.Now().Add(1*time.Hour)),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ = json.MarshalIndent(secondUser, "", "  ")
	fmt.Printf("created second user: %s\n", result)

	return nil
}
