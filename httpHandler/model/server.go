package model

import (
	"database/sql"
	"github.com/olivere/elastic"
	"github.com/streadway/amqp"
)

type Configs struct {
	Db     *sql.DB
	Ec     *elastic.Client
	Rabbit *amqp.Connection
	Env    Env
}
