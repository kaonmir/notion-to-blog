package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Notion  Notion  `yaml:"notion"`
	Tistory Tistory `yaml:"tistory"`
}

type Notion struct {
	TokenV2      string `yaml:"token_v2" env:"TOKEN_V2" env-required:"true"`
	TablePageUrl string `yaml:"table_page_url" env:"TABLE_PAGE_URL" env-required:"true"`
	// DownloadDir  string `yaml:"download_dir" env:"DOWNLOAD_DIR" env-required:"true"`
}

type Tistory struct {
	Id          string `yaml:"id" env:"ID" env-required:"true"`
	Password    string `yaml:"password" env:"PASSWORD" env-required:"true"`
	BlogName    string `yaml:"blog_name" env:"BLOG_NAME" env-required:"true"`
	SecretKey   string `yaml:"secret_key" env:"SECRET_KEY" env-required:"true"`
	ClientId    string `yaml:"client_id" env:"CLIENT_ID" env-required:"true"`
	RedirectUri string `yaml:"redirect_uri" env:"REDIRECT_URI" env-required:"true"`
}

// ---

func NewConfig() *Config {
	return NewConfigWithFilename("config.yml")
}

func NewConfigWithFilename(filename string) *Config {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadConfig(dir+"/"+filename, cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("config error: %w", err))
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
