migrate:
	docker-compose run api /bin/bash -c "migrate -database postgres://postgres:postgres@postgres:5432/flashcards?sslmode=disable -path migrations up"
