package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Album struct {
	title		string
	artist		string
	year		string
}

type Track struct {
	title		string
	album		*Album
	artist		string
	duration	time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func (a Album) String() string {
	return fmt.Sprintf("%s is an album by %s released on %s", a.title, a.artist, a.year)
}

func (t Track) String() string {
	return fmt.Sprintf("%s is a track in %s by %s with a duration of %s", t.title, t.album.title, t.artist, t.duration)
}

// The interface for sorting
type byDuration []Track

func (tracks byDuration) Len() int { return len(tracks) }

func (tracks byDuration) Less(i, j int) bool {
	return tracks[i].duration < tracks[j].duration
}

func (tracks byDuration) Swap(i, j int) {
	tracks[i], tracks[j] = tracks[j], tracks[i]
}
// ------

func printTracks(tracks *[]Track) {
	fmt.Println(strings.Repeat("-", 30))
	for _, t := range *tracks {
		fmt.Printf("%s\t|%s\n", t.title, t.duration)
	}
	fmt.Println(strings.Repeat("-", 30))
}
func main() {
	casaBabylon := Album{"Casa Babylon", "Mano Negra", "1994"}
	tracks := []Track {
		{"Viva Zapata", &casaBabylon, "Mano Negra", length("2m04s")},
		{"Casa Babylon", &casaBabylon, "Mano Negra", length("2m01s")},
		{"Love & Hate", &casaBabylon, "Mano Negra", length("2m28s")},
		{"Machine Gun", &casaBabylon, "Mano Negra", length("4m25s")},
		{"Bala Perdida", &casaBabylon, "Mano Negra", length("2m13s")},
	}

	fmt.Println(casaBabylon)
	fmt.Println(tracks)

	sort.Sort(byDuration(tracks))

	printTracks(&tracks)
}