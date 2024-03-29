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

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")}, {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
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

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	sort.Sort(sort.Reverse(byArtist(tracks)))
	var ani animal
	ani = &dog{}
	ani.Reverse(1, 2)
	ani = &cat{ani}
	ani.Reverse(1, 2)
}

type animal interface {
	Voice()
	Reverse(i, j int)
}

type dog struct{}

func (*dog) Voice() {
	fmt.Println("Woo")
}

func (*dog) Reverse(i, j int) {
	fmt.Printf("%d, %d\n", i, j)
}

type cat struct {
	animal
}

func (*cat) Voice() {
	fmt.Println("Meow")
}

func (c *cat) Reverse(i, j int) {
	c.animal.Reverse(j, i)
}
