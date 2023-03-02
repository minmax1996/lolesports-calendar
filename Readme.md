# LOLesports calendar

run it
```
docker run -d \
	-p 8080:8080 \
	-e "ESPORTS_TOKEN=yourtoken" \
	minmax1996/lolesports-calendar
```

call it

```
curl --location 'http://localhost:8080/calendar?leagues=lec,lck&teams=fnc,c9&spoiler=true'
```

query params:
- leagues: list of target leagues divided with `,` to fetch calendar for
- teams: list of favorite teams to watch (also divided with `,`) api will filter out matches for any other team, if param teams is precent
- spoiler: boolean flag, to include match results in calendar