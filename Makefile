build: 
	docker build -t tech-test-jc .

run:
	docker build -t tech-test-jc .
	docker run -it -p 5000:5000 tech-test-jc

test:
	go test -cover
