package main

import (
	"log"

	"github.com/dnsoftware/socialposter/internal/config"
	"github.com/dnsoftware/socialposter/internal/poster"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	vk := poster.NewPlatformVk(cfg.VkToken)

	poster := poster.NewPoster(vk)
	_ = poster
}
