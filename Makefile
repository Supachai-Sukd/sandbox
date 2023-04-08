install:
	go mod tidy

dev:
	go run main.go

test: test-unit test-integration

test-unit:
	go test -tags=unit -v ./...

test-coverage:
	go test -cover -tags=unit ./...

test-integration:
	docker-compose -f docker-compose.it-test.yaml down && \
	docker-compose -f docker-compose.it-test.yaml up --build --force-recreate --abort-on-container-exit --exit-code-from it_tests


