.PHONY: clean, build_docker, run_docker


all: clean zoli_linuxos build_docker run_docker

clean:
	rm zoli_linuxos

zoli_linuxos:
	env GOOS=linux GOARCH=amd64 go build -o zoli_linuxos arr_example1.go

build_docker:
	docker build -t mygotest .

run_docker:
	docker run mygotest

test_script:
	./test_script.sh