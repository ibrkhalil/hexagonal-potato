watch-logger:
	tail -f app.log
watch:
	nodemon --exec go run main.go --signal SIGTERM
