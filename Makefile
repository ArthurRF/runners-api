include .env
.PHONY: m-generate

m-generate:
		atlas migrate diff --env gorm

m-hash:
		atlas migrate hash

seed:
		atlas migrate new seed_$(name)

m-apply:
		atlas migrate hash; atlas migrate apply --url ${DB_URL}