run:
	go run ./httpWaitServer/main.go &
	go test -count 1000 -v ./httpClients -run $(test)

runAll:
	make run test=".*"