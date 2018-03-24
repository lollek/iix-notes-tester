
all:
	test_target=http://localhost:9002 test_jwt=debug go test

live-ro:
	test_target=https://iix.se test_jwt=debug go test
