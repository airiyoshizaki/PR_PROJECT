include .env

.PHONY: setup
setup:
	docker network create hoge-network-1 && \
	cp .env.example .env && \
	docker-compose build

.PHONY: build
build:
	docker-compose build

.PHONY: up-d
up-d:
	docker-compose up -d

.PHONY: up
up:
	docker-compose up 

.PHONY: down
down:
	docker-compose down

.PHONY: exec-app
exec-app:
	docker-compose exec hoge-app-1 sh

.PHONY: exec-app-db
exec-app-db:
	docker-compose exec hoge-app-1_db sh


# APIドキュメント生成コマンド
# # GoのエントリーポイントファイルからSwaggerドキュメントを生成
# # Swagger (OpenAPI 2.0) 仕様を OpenAPI 3.x 仕様に変換し、YAML形式で保存
# # Swagger (OpenAPI 2.0) 仕様を JSON形式に変換し、保存
.PHONY: gen-api-doc
gen-api-doc:
	docker-compose exec hoge-app-1 sh -c "\
		swag init ./main.go && \
		swagger2openapi --outfile ./docs/openapi.yaml ./docs/swagger.yaml && \
		swagger2openapi --outfile ./docs/swagger.json ./docs/swagger.yaml \
    "