package main

import (
	"skill-api-kafka-consumer/config"
	"skill-api-kafka-consumer/database"
)

func main() {
	c := config.Configuration()

	_ = database.Postgres(c.PostgresURI)
}
