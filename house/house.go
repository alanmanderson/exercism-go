package house

//import "fmt"

type verse [2]string

var verses = []verse{
	verse{"lay in", "the house that Jack built"},
	verse{"ate", "the malt"},
	verse{"killed", "the rat"},
	verse{"worried", "the cat"},
	verse{"tossed", "the dog"},
	verse{"milked", "the cow with the crumpled horn"},
	verse{"kissed", "the maiden all forlorn"},
	verse{"married", "the man all tattered and torn"},
	verse{"woke", "the priest all shaven and shorn"},
	verse{"kept", "the rooster that crowed in the morn"},
	verse{"belonged to", "the farmer sowing his corn"},
	verse{"", "the horse and the hound and the horn"},
}

//Verse stuff
func Verse(v int) string {
	out := "This is " + verses[v-1][1] + "\n"
	for i := v - 2; i >= 0; i-- {
		out += "that " + verses[i][0] + " " + verses[i][1] + "\n"
	}
	return out[:len(out)-1] + "."
}

//Song stuff
func Song() string {
	out := ""
	for i := 1; i <= len(verses); i++ {
		out += Verse(i) + "\n\n"
	}
	return out[:len(out)-2]
	//return out
}
