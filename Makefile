compile:
	go build -o "./bin/rkomi" ./src

run:
	./bin/rkomi

test: compile
	echo "ogs glicko test (Samraku, LittlePebble):"
	./bin/rkomi 1745 1062
	echo "ogs glicko test (Humbaba, gennan):"
	./bin/rkomi 1937 2122
	echo "kyu test (Samraku, LittlePebble):"
	./bin/rkomi 3k 14k
	echo "dan test (Humbaba, gennan):"
	./bin/rkomi 1d 3d
clean:
	rm ./bin/rkomi
