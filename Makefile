run: build
	@./bin/app

build:
	@go build -o bin/app .

templ:
	@templ generate --watch --proxy=http://localhost:3000
