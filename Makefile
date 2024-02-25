# build: sqlc-gen
# 	docker-compose -f ./docker-compose-app.yaml build

docker:
	docker-compose up

local-webapp:
	cd apps/webapp && npm run dev

local-backend:
	cd apps/backend && go run ./cmd/app

init-dev-docker:
	docker pull postgres:15
	docker pull sqlc/sqlc

# gen/sql
# sqlc-gen:
# 	@for mod in ${MODULES}; do \
# 		$(MAKE) sqlc-gen-$$mod; \
# 	done
#
# sqlc-gen-module:
# 	docker run --rm -v $(shell pwd):/src -w /src \
# 		sqlc/sqlc generate -f $(call get_sqlc_config_path,${SQLC_MODULE})
#
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
