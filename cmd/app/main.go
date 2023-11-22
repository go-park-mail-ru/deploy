package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/errgroup"
)

type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
}

var dbPath = flag.String("db", "./tasks.db", "path to database file")

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sqlx.Connect("sqlite3", *dbPath)
	if err != nil {
		return err
	}

	m, err := migrate.New("file://migrations", "sqlite3://"+*dbPath)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/tasks/", func(c echo.Context) error {
		ctx := c.Request().Context()

		var tasks []Task
		err := db.SelectContext(ctx, &tasks, "SELECT * FROM tasks")
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/tasks/", func(c echo.Context) error {
		ctx := c.Request().Context()

		var task Task
		err := c.Bind(&task)
		if err != nil {
			return err
		}

		if task.ID == uuid.Nil {
			task.ID = uuid.New()
		}

		_, err = db.NamedExecContext(ctx, "INSERT INTO tasks (id, title, description) VALUES (:id, :title, :description)", task)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	})

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()
		return e.Shutdown(ctx)
	})

	g.Go(func() error {
		fmt.Println("Starting service on http://127.0.0.1:8080")

		err := e.Start(":8080")
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	return g.Wait()
}
