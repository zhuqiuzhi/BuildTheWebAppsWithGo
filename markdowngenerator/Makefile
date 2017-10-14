all: get build run

run:
	./markdowngenerator

build:
	go build -o markdowngenerator main.go

get:
	go get github.com/microcosm-cc/bluemonday
	go get gopkg.in/russross/blackfriday.v2
