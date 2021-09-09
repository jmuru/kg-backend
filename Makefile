run: bin/server
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/server: main.go
	go build -o bin/server main.go

clean:
	rm -rf bin