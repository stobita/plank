DEV_COMPOSE=deployments/development/docker-compose.yml
TEST_COMPOSE=deployments/test/docker-compose.yml

dev-up:
	docker-compose -f $(DEV_COMPOSE) up -d
dev-down:
	docker-compose -f $(DEV_COMPOSE) down
dev-logs:
	docker-compose -f $(DEV_COMPOSE) logs -f

sqlboiler:
	docker-compose -f $(DEV_COMPOSE) exec api sqlboiler mysql --wipe -o ./internal/rdb -c ./configs/sqlboiler.toml -p rdb --no-auto-timestamps --no-tests
