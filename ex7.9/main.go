package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
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

var trackTable = template.Must(template.New("trackTable").Parse(`
<table>

<tr style='text-align: left'>
<th onclick="submitform('Title')">Title
<form action="" name="Title" method="post">
		<input type="hidden" name="orderby" value="Title"/>
</form>
</th>
<th onclick="submitform('Artist')">Artist
<form action="" name="Artist" method="post">
		<input type="hidden" name="orderby" value="Artist"/>
</form>
</th>
<th onclick="submitform('Album')">Album
<form action="" name="Album" method="post">
		<input type="hidden" name="orderby" value="Album"/>
</form>
</th>
<th onclick="submitform('Year')">Year
<form action="" name="Year" method="post">
		<input type="hidden" name="orderby" value="Year"/>
</form>
</th>
<th onclick="submitform('Length')">Length
<form action="" name="Length" method="post">
		<input type="hidden" name="orderby" value="Length"/>
</form>
</th>
</tr>

{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))

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

func printTracks(w io.Writer, s sort.Interface) {

	trackTable.Execute(w, s)
}

func main() {

	mt := multiTier{tracks, "Album", ""}
	sort.Sort(mt)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			r.ParseForm()
			mt.Primary = r.FormValue("orderby")
			sort.Sort(mt)
			printTracks(w, mt)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
