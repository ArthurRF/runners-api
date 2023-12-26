include .env
.PHONY: m-generate

dev:
		go run cmd/server/main.go

m-generate:
		atlas migrate diff --env gorm

m-hash:
		atlas migrate hash

seed:
		atlas migrate new seed_$(name)

m-apply:
		atlas migrate hash; atlas migrate apply --url ${DB_URL}