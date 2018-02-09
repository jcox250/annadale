build: 
	docker-compose build

run:
	docker-compose build
	docker-compose up

test:
	go test -cover
