package defaults

import (
	"github.com/theNullP0inter/kenobi/cmd"
	"github.com/theNullP0inter/kenobi/config"
	"time"
)

var Config = config.Config{
	"HttpPort":        8000,
	"AllowedOrigins":  []string{"*"},
	"AllowedMethods":  []string{"GET", "HEAD", "POST", "OPTIONS", "PUT", "PATCH"},
	"AllowedHeaders":  []string{"accept", "accept-encoding", "authorization", "content-type", "dnt", "origin", "user-agent", "x-csrftoken", "x-requested-with"},
	"ReadTimeout":     1 * time.Minute,
	"CommandRegistry": []cmd.ExecCommand{},
}
