package main

import (
	"datenote/cms/handler"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	cpb "datenote/gunk/v1/category"
	epb "datenote/gunk/v1/event"
)

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("cms/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("Error loading configuration: %v", err)
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	store := sessions.NewCookieStore([]byte(config.GetString("session.secret")))
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", config.GetString("datenote.host"), config.GetString("datenote.port")),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("Connection failed", err)
	}

	ec := epb.NewEventServiceClient(conn)
	cc := cpb.NewCategoryServiceClient(conn)
	r := handler.GetHandler(decoder, store, ec, cc)

	host, port := config.GetString("server.host"), config.GetString("server.port")
	log.Println("Server  starting...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r); err != nil {
		log.Fatal(err)
	}
}
