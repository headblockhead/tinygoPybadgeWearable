run:
	go run generator/main.go
	cd ..
	tinygo flash -target pybadge .