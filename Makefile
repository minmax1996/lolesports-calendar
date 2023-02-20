PWD = $(pwd)

docker-build:
	docker build -t minmax1996/lolesports-calendar:latest .

docker-run:
	docker run -d -p 8080:8080  minmax1996/lolesports-calendar

docker-run-withdiscord:
	docker run -d \
	-p 8080:8080 \
	minmax1996/lolesports-calendar