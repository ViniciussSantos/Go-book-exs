package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type multiTier struct {
	Tracks    []*Track
	Primary   string
	Secondary string
}

func (mt multiTier) Len() int      { return len(mt.Tracks) }
func (mt multiTier) Swap(i, j int) { mt.Tracks[i], mt.Tracks[j] = mt.Tracks[j], mt.Tracks[i] }

func (mt multiTier) Less(i, j int) bool {

	sortKey := mt.Primary

	for i := 0; i < 3; i++ {

		switch sortKey {

		case "Title":
			if mt.Tracks[i].Title != mt.Tracks[j].Title {
				return mt.Tracks[i].Title < mt.Tracks[j].Title
			}

		case "Artist":
			if mt.Tracks[i].Artist != mt.Tracks[j].Artist {
				return mt.Tracks[i].Artist < mt.Tracks[j].Artist
			}

		case "Album":
			if mt.Tracks[i].Album != mt.Tracks[j].Album {
				return mt.Tracks[i].Album < mt.Tracks[j].Album
			}

		case "Year":
			if mt.Tracks[i].Year != mt.Tracks[j].Year {
				return mt.Tracks[i].Year < mt.Tracks[j].Year
			}

		case "Length":
			if mt.Tracks[i].Length != mt.Tracks[j].Length {
				return mt.Tracks[i].Length < mt.Tracks[j].Length
			}

		}

		if i == 1 {
			sortKey = mt.Secondary
		}

	}

	return false
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {

	mt := multiTier{tracks, "Artist", "Year"}

	printTracks(mt.Tracks)

	fmt.Println("Sorting by Artist and then by Year")

	sort.Sort(mt)

	printTracks(mt.Tracks)

}
