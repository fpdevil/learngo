package main

import "fmt"

// package level declaration
// #1b: create a song struct type
type song struct {
	title, artist string
}

// #5: structs can contain other structs as their fields
// creation of the playlists as per the genre
type playlist struct {
	genre string

	// cerating title and artist for each is a cumbersome task
	// instead we can offload that to another struct
	// songTitles []string
	// songArtist []string

	// #6: include a slice of song structs (from 1b)
	songs []song
}

func main() {
	fmt.Println("---- 24-structs/03-compare-assign ----")
	fmt.Println("\n")

	// #1: create two struct values with the same type
	song1 := song{title: "skyfall", artist: "adele adkins"}
	song2 := song{title: "million years ago", artist: "adele adkins"}

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// #2:
	// structs can be compared with one another
	// field by field as below
	if song1.title == song2.title && song1.artist == song2.artist {
		fmt.Println("#2 -> songs are equal.")
	} else {
		fmt.Println("#2 -> songs are not equal.")
	}

	// structs can be copied with assignment
	// song1 = song2

	// #3:
	// instead of the above, if we put comparision of structs
	// directly go can take care of the inndividual comparision
	if song1 == song2 {
		fmt.Println("#3 -> songs are equal.")
	} else {
		fmt.Println("#3 -> songs are not equal.")
	}

	// #8:
	// moving to a separate variable from # 7b
	songs := []song{
		{title: "skyfall", artist: "adele adkins"},
		{title: "million years ago", artist: "adele adkins"},
	}

	// #7: now create a playlist as below
	rock := playlist{
		genre: "indie rock",
		songs: songs,

		// # 7b:
		// element type also need not be removed, so we can essentially
		// specify them as
		// songs: []song{
		// 	song{title: "skyfall", artist: "adele adkins"},
		// 	song{title: "million years ago", artist: "adele adkins"},
		// },

		// # 7a:
		// instead of typing here, we can move them to a variable
		// songs: []song{
		// 	song{title: "skyfall", artist: "adele adkins"},
		// 	song{title: "million years ago", artist: "adele adkins"},
		// },
	}

	// #9:
	// we can clone the playlist as below
	// clone := rock

	// this is not possible as the underlying []song slices cannot be compared
	// if rock == clone {
	// 	fmt.Println("equal song titles...")
	// }
	// even song == song is also not possible

	// #11:
	// changing the title of the first song
	// song is a clone, it cannot change the original struct value
	song := rock.songs[0]
	song.title = "21 forever"

	// #11b
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])

	// #11c: to change the song, directly set the original one
	rock.songs[0].title = "21 forever"
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])

	// #10: printing
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		// s := rock.songs[i]

		// #12b: s is a copy inside because struct values are copied
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	// #10:
	// printing the songs
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("\n%-20s %20s\n", s.title, s.artist)
	}

}
