package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type ServerConfig struct {
	URL      string `json:"url"`
	Schedule string `json:"schedule"` // Cron expression (e.g., "*/5 * * * *" for every 5 minutes)
	Name     string `json:"name"`
}

type Config struct {
	Servers []ServerConfig `json:"servers"`
}

func loadConfig() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &config, nil
}

func pingRedis(url string, name string) error {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return fmt.Errorf("error parsing Redis URL: %v", err)
	}

	client := redis.NewClient(opt)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the ping result
	result, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("error pinging Redis: %v", err)
	}

	// Print more detailed success information
	log.Printf("Redis server ping successful - Name: %s, Response: %s, Time: %s",
		name,
		result,                          // Usually returns "PONG"
		time.Now().Format(time.RFC3339)) // Add timestamp
	return nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	c := cron.New()

	for _, server := range config.Servers {
		// 1. Create a new variable to avoid closure issues
		server := server // This is important for goroutines in the loop

		// 2. Validate the cron expression
		if _, err := cron.ParseStandard(server.Schedule); err != nil {
			log.Printf("Invalid cron expression for server %s: %v", server.URL, err)
			continue // Skip this server if schedule is invalid
		}
		//before adding the cron job, print the server name and do ping test
		if err := pingRedis(server.URL, server.Name); err != nil {
			log.Printf("Error pinging Redis server %s: %v", server.Name, err)
			log.Printf("Skipping cron job for %s", server.Name)
			continue
		}
		// 3. Add the cron job
		_, err := c.AddFunc(server.Schedule, func() {
			if err := pingRedis(server.URL, server.Name); err != nil {
				log.Printf("Error pinging Redis server %s: %v", server.Name, err)
			}
		})

		// 4. Handle any errors in adding the cron job
		if err != nil {
			log.Printf("Error adding cron job for %s: %v", server.URL, err)
			continue
		}

		// 5. Log success
		log.Printf("Added cron job for Redis server %s with schedule: %s", server.Name, server.Schedule)
	}

	c.Start()
	log.Println("Started Redis ping service")

	// Keep the application running
	select {}
}
