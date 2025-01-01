package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

var (
	mariadbContainer testcontainers.Container
	mariadbHost      string
	mariadbPort      nat.Port
)

func SetupContainer(ctx context.Context) (string, error) {
	var err error
	mariadbContainer, err = mariadb.Run(ctx,
		"mariadb:latest",
		mariadb.WithDatabase("test_db"),
		mariadb.WithUsername("root"),
		mariadb.WithPassword("password"),
	)
	if err != nil {
		return "", err
	}

	mariadbHost, err = mariadbContainer.Host(ctx)
	if err != nil {
		return "", err
	}

	mariadbPort, err = mariadbContainer.MappedPort(ctx, "3306")
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("root:password@tcp(%s:%s)/test_db?charset=utf8mb4&parseTime=True&loc=Local", mariadbHost, mariadbPort.Port())
	return dsn, nil
}

func TerminateMariaDBContainer(ctx context.Context) {
	if mariadbContainer != nil {
		if err := mariadbContainer.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate container: %s", err)
		}
	}
}
