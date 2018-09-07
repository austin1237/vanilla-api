.PHONEY: test
test:
	docker-compose run api go test ./... -short

.PHONEY: integration
integration:
	docker-compose run api go test ./...
