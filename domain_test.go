package main

import (
	"testing"
	"time"
)

var feedParsingTests = []struct {
	name       string
	body       []byte
	parsedFeed *parsedFeed
	errMsg     string
}{
	{"RSS - Minimal",
		[]byte(`<?xml version='1.0' encoding='UTF-8'?>
<rss>
  <channel>
    <title>News</title>
    <item>
      <title>Snow Storm</title>
      <link>http://example.org/snow-storm</link>
      <pubDate>Fri, 03 Jan 2014 22:45:00 GMT</pubDate>
    </item>
    <item>
      <title>Blizzard</title>
      <link>http://example.org/blizzard</link>
      <pubDate>Sat, 04 Jan 2014 08:15:00 GMT</pubDate>
    </item>
  </channel>
</rss>
</xml>`),
		&parsedFeed{
			name: "News",
			items: []parsedItem{
				{
					title:           "Snow Storm",
					url:             "http://example.org/snow-storm",
					publicationTime: time.Date(2014, 1, 3, 22, 45, 0, 0, time.UTC),
				},
				{
					title:           "Blizzard",
					url:             "http://example.org/blizzard",
					publicationTime: time.Date(2014, 1, 4, 8, 15, 0, 0, time.UTC),
				},
			}},
		"",
	},
	{"Atom - Minimal",
		[]byte(`<?xml version='1.0' encoding='UTF-8'?>
<feed>
  <title>News</title>
  <entry>
    <title>Snow Storm</title>
    <link href="http://example.org/snow-storm" />
    <published>2014-01-03T22:45:00Z</published>
  </entry>
  <entry>
    <title>Blizzard</title>
    <link href="http://example.org/blizzard" />
    <published>2014-01-04T08:15:00Z</published>
  </entry>
</feed>
</xml>`),
		&parsedFeed{
			name: "News",
			items: []parsedItem{
				{
					title:           "Snow Storm",
					url:             "http://example.org/snow-storm",
					publicationTime: time.Date(2014, 1, 3, 22, 45, 0, 0, time.UTC),
				},
				{
					title:           "Blizzard",
					url:             "http://example.org/blizzard",
					publicationTime: time.Date(2014, 1, 4, 8, 15, 0, 0, time.UTC),
				},
			}},
		"",
	},
}

func TestParseFeed(t *testing.T) {
	for i, tt := range feedParsingTests {
		actual, err := parseFeed(tt.body)
		if err != nil && err.Error() != tt.errMsg {
			t.Errorf("%d. %s: Unexpected error: %v", i, tt.name, err)
		}
		if actual == nil {
			if tt.parsedFeed != nil {
				t.Errorf("%d. %s: Actual parsed feed should not have been nil, but it was", i, tt.name)
			}
			continue
		}
		if tt.parsedFeed == nil {
			t.Errorf("%d. %s: Actual parsed feed should have been nil, but it was not", i, tt.name)
			continue
		}
		if actual.name != tt.parsedFeed.name {
			t.Errorf("%d. %s: Expected name to be %#v, but it was %#v", i, tt.name, tt.parsedFeed.name, actual.name)
		}
		if len(actual.items) != len(tt.parsedFeed.items) {
			t.Errorf("%d. %s: Expected %d items, but instead found %d items", i, tt.name, len(tt.parsedFeed.items), len(actual.items))
			continue
		}
		for j, actualItem := range actual.items {
			expectedItem := tt.parsedFeed.items[j]
			if actualItem.title != expectedItem.title {
				t.Errorf("%d. %s Item %d: Expected title %#v, but is was %#v", i, tt.name, j, expectedItem.title, actualItem.title)
			}
			if actualItem.url != expectedItem.url {
				t.Errorf("%d. %s Item %d: Expected url %#v, but is was %#v", i, tt.name, j, expectedItem.url, actualItem.url)
			}
			if !actualItem.publicationTime.Equal(expectedItem.publicationTime) {
				t.Errorf("%d. %s Item %d: Expected publicationTime %s, but is was %s", i, tt.name, j, expectedItem.publicationTime, actualItem.publicationTime)
			}
		}
	}
}