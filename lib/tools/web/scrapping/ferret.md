[ferret](https://github.com/MontFerret/ferret)

*Description*: ferret is a web scraping system. It aims to simplify data extraction from the web for UI testing, machine learning, analytics and more.

*Labels*: #Go #HTML #WebScraping

*Docs*: https://www.montferret.dev/

*Examples*:

```bash
$ cat > fetch.fql << EOF

LET doc = DOCUMENT('https://soundcloud.com/charts/top', {
    driver: 'cdp'
})

WAIT_ELEMENT(doc, '.chartTrack__details', 5000)

LET tracks = ELEMENTS(doc, '.chartTrack__details')

FOR track IN tracks
    RETURN {
        artist: TRIM(INNER_TEXT(track, '.chartTrack__username')),
        track: TRIM(INNER_TEXT(track, '.chartTrack__title'))
    }
EOF
$ ferret exec fetch.fql
```
