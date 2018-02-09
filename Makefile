build: 
	docker-compose build

run:
	docker-compose build
	docker-compose up

	# docker build -t tech-test-jc .
	# docker run -it -p 5000:5000 tech-test-jc

test:
	go test -cover
