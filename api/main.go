package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skill-api-kafka/config"
	"skill-api-kafka/database"
	"skill-api-kafka/kafka"
	"skill-api-kafka/skill"
	"syscall"
	"time"
)

func main() {
	c := config.Configuration()

	db := database.Postgres(c.PostgresURI)
	storage := skill.NewSkillStrage(db)
	producer := kafka.Producer(c.Kafka)
	queue := skill.NewSkillQueue(producer, c.Kafka)

	r := Router(storage, queue)

	defer db.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closedChannel := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("Received signal. Gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Error:", err)
			}
		}
		close(closedChannel)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println("Error Serve:", err)
	}

	<-closedChannel
	fmt.Println("Server shutdown successfully")

}

func Router(storage skill.SkillStorage, producer skill.SkillQueue) *gin.Engine {
	r := gin.Default()
	h := skill.NewSkillHandler(storage, producer)

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/skills/:key", h.GetSkill)
		v1Group.GET("/skills", h.GetSkills)
		v1Group.POST("/skills")
		v1Group.PUT("/skills/:key")
		v1Group.PATCH("/skills/:key/actions/name")
		v1Group.PATCH("/skills/:key/actions/description")
		v1Group.PATCH("/skills/:key/actions/logo")
		v1Group.PATCH("/skills/:key/actions/tags")
		v1Group.DELETE("/skills/:key")
	}

	return r
}
