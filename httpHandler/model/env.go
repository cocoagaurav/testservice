package model

type Env struct {
	SqlUrl            string `envconfig:"SQL_URL" default:"root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local"`
	RabbitUrl         string `envconfig:"RABBIT_URL" default:"amqp://guest:guest@rabbitmq-server:5672/"`
	FirebaseServiceId string `envconfig:"FIREBASE_SERVICE_ID" default:"firebase-adminsdk-6b9tl@testproject-fa267.iam.gserviceaccount.com"`
}
