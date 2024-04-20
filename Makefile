# build: sqlc-gen
# 	docker-compose -f ./docker-compose-app.yaml build

docker:
	docker-compose up --build --force-recreate

rebuild-containers:
	docker-compose build

local-web-ui:
	cd apps/web-ui && npm run dev

local-backend:
	cd apps/backend && go run ./cmd/app

init-dev-docker:
	docker pull postgres:15
	docker pull sqlc/sqlc
	docker pull ghcr.io/amacneil/dbmate

gen-sqlc:
	docker run --rm -v $(shell pwd):/src -w /src \
		sqlc/sqlc generate -f ./apps/backend/internal/infra/sqlc-pg/sqlc.yalm

# migrations
get_cmd = $(word 2,$(subst -, ,$1))

MIGRATION = #

migrate-%:
	$(eval MIGRATE_CMD = $(call get_cmd,$@))
	docker-compose -f ./apps/backend/docker-compose-migrate.yaml \
		run dbmate ${MIGRATE_CMD} ${MIGRATION}

# sqlc-gen-voting: $(eval SQLC_MODULE = voting) sqlc-gen-module
#
# # gen/mocks
# # todo
#
# # migrations
# get_cmd = $(word 2,$(subst -, ,$1))
#
# MIGRATION = #
#
# migrate-%:
# 	$(eval MIGRATE_CMD = $(call get_cmd,$@))
# 	docker-compose -f ./docker-compose-migrate.yaml \
# 		run dbmate ${MIGRATE_CMD} ${MIGRATION}
