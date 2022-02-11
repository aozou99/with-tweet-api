package repository

import (
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/spiegel-im-spiegel/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

var repository = &Repository{}

func NewRepository() *Repository {
	repository.init()
	return repository
}

func (r *Repository) TranslatedTweet() *TranslatedTweetRepository {
	return &TranslatedTweetRepository{*repository}
}

func (r *Repository) init() {
	if r.db != nil {
		return
	}
	r.db = newDB()
}

func newDB() *gorm.DB {
	cfg, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Interface("error", errs.Wrap(err)).Send()
		panic("failed to connect db.")
	}
	cfg.Logger = zerologadapter.NewContextLogger()
	cfg.LogLevel = pgx.LogLevelDebug
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*cfg),
	}), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Error().Interface("error", errs.Wrap(err)).Send()
		panic("failed to connect db.")
	}
	return db
}
