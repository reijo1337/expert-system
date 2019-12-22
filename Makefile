build.production:
	go build -mod=vendor -o .bin/production cmd/production/*.go

run.production:build.production
	.bin/production; rm .bin/production

run.logic:
	python3 cmd/logical/pylog.py

run.fuzzy:
	python3 cmd/fuzzy/main.py

build.fcmeans:
	g++ -o .bin/fcmeans -I cmd/fcmeans/eigen/ cmd/fcmeans/fcm.cpp cmd/fcmeans/fcm.h cmd/fcmeans/main.cpp

run.fcmeans:build.fcmeans
	.bin/fcmeans | python3 cmd/fcmeans/main.py; rm .bin/fcmeans