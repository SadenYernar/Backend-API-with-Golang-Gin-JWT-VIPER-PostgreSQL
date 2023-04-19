docker:
	docker pull postgres
	docker run --name mypg -p 5432:5432 -e POSTGRES_PASSWORD=qwerty -d postgres	
	docker exec mypg psql -h localhost -U postgres
	

go: 
	go run ./cmd/main.go

.PHONY: docker go