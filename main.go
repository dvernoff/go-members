package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dg *discordgo.Session
var botReady = false
var botToken string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	botToken = os.Getenv("DISCORD_TOKEN")
	if botToken == "" {
		log.Fatal("DISCORD_TOKEN is not set")
	}

	dg, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	dg.AddHandler(ready)
	dg.Identify.Intents = discordgo.IntentsGuildMembers

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	for !botReady {
		time.Sleep(500 * time.Millisecond)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/check/:serverid/:userid", checkMemberDirect)

	go r.Run(":5372")

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	fmt.Println("API: GET http://localhost:5372/check/:serverid/:userid")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func checkMemberDirect(c *gin.Context) {
	serverID := c.Param("serverid")
	userID := c.Param("userid")

	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s", serverID, userID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bot "+botToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"is_member": false})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var member map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&member)

		username := "unknown"
		if user, ok := member["user"].(map[string]interface{}); ok {
			if name, ok := user["username"].(string); ok {
				username = name
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"is_member": true,
			"username":  username,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"is_member": false})
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Bot %s is ready!\n", event.User.Username)
	botReady = true
}
