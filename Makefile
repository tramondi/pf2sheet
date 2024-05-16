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
	docker pull ghcr.io/amacneil/dbmate

# migrations
get_cmd = $(word 2,$(subst -, ,$1))

MIGRATION = #

dbmate-%:
	$(eval MIGRATE_CMD = $(call get_cmd,$@))
	docker-compose -f ./apps/backend/docker-compose-migrate.yaml \
		run dbmate ${MIGRATE_CMD} ${MIGRATION}
