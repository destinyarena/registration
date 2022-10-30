package main

import (
	"fmt"
	"os"

	"github.com/arturoguerra/go-logging"
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/registration/internal/jwtmanager"
	"github.com/destinyarena/registration/internal/profilestore"
	registration "github.com/destinyarena/registration/internal/registration"
	"github.com/destinyarena/registration/internal/router"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log := logging.New()

	jwtManager, err := jwtmanager.NewDefault(log)
	if err != nil {
		log.Fatal(err)
	}

	profileStore, err := profilestore.NewDefault(log)
	if err != nil {
		log.Fatal(err)
	}

	dsession, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	r, rconfig, err := router.NewDefault(log)
	if err != nil {
		log.Fatal(err)
	}

	registerGroup := r.Group("/api/v2/registration", middleware.Logger())

	registrationHandler, err := registration.NewDefault(log, jwtManager, profileStore, dsession)
	if err != nil {
		log.Fatal(err)
	}

	registrationHandler.Register(registerGroup)

	log.Infof("Running on %s:%s", rconfig.Host, rconfig.Port)
	r.Logger.Fatal(r.Start(fmt.Sprintf("%s:%s", rconfig.Host, rconfig.Port)))
}
