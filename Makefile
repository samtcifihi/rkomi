compile:
	go build -o "./bin/rkomi" ./src

run:
	./bin/rkomi

test: compile
	./bin/rkomi
clean:
	rm ./bin/rkomi
