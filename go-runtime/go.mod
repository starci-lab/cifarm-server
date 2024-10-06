module cifarm-server

go 1.22.5

require (
	github.com/go-co-op/gocron/v2 v2.11.0
	github.com/google/uuid v1.6.0
	github.com/hasura/go-graphql-client v0.13.0
	github.com/heroiclabs/nakama-common v1.33.0
)

require (
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
)

replace github.com/gorilla/websocket => github.com/gorilla/websocket v1.5.3
