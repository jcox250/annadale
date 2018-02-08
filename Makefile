build: 
	docker build -t dale .

run:
	docker build -t dale .
	docker run -it dale

test:
	go test -cover
