include .env
export

run:
	@tailwindcss -i assets/input.css -o assets/output.css
	@templ generate
	@go run main.go
