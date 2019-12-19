build.production:
	go build -mod=vendor -o .bin/production cmd/production/*.go

run.production:build.production
	.bin/production; rm .bin/production

run.logic:
	python3 cmd/logical/pylog.py