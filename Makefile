dev:
	go build cmd/main.go && mv main hangman
	./hangman

build:
	go build cmd/main.go && mv main hangman