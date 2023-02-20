PWD = $(pwd)

docker-build:
	docker build -t minmax1996/lolesports-calendar:latest .

docker-run:
	docker run -d \
	-p 8080:8080 \
	-e "ESPORTS_TOKEN=yourtoken" \
	minmax1996/lolesports-calendar