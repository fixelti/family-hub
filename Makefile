# Миграции для базы данных
migration_up:
	goose -dir internal/repository/migration postgres "postgresql://root:root@localhost:5432/test?sslmode=disable" up

migration_down:
	goose -dir internal/repository/migration postgres "postgresql://root:root@localhost:5432/test?sslmode=disable" down