package main

// ACHTUNG: Unvollständig: Entscheidungsproblem beim Decode:  a+b= c oder y+d= c
import (
	"fmt"
	"strings"
)

func main(args []string) {
	cleartext := `ZA`

	/*
		key := `bevor die erste schlacht mich rief oder kam die schlacht zu mir stand
		ich am rand der kaempfe nur ein schwaechling war ich hier sah dort
		krieger so viel staerker als ich jemals koennte sein mein herz es sehnte
		sich nach herd und heim doch ich blieb und sah die streiter und ein
		kaempfer stach heraus die ruestung schon zerschunden und doch hielt sie
		tapfer aus ein sturm aus zorn und feuer stets dort wo die reihe bricht
		wo sie auch war da hin da zog es mich hin ref und es war nur ein
		augenblick ein kurzer wimpernschlag trag ihr bild nun stets bei mir
		wohin ich auch gehen mag bis an das ende meiner tage und wenn auch die
		welt zerbricht ist sie für mich in tiefster nacht ein licht ihr herz ist
		stark stolz und mutig hab sie nie verzagen sehn sie war eine von uns und
		wollte nie darueber stehn ich stell mir vor wie sie beschloss einst
		eines tages gross zu sein nd fuehlte mich von da an nicht mehr klein ref
		und so sammelte ich ruestung stueck fuer stueck von hier und da und
		machte mir gedanken wer ich bin und wer ich war ich weiss auch nicht wer
		mich sieht und ob das schicksal mich wohl streift und ob mein feuer auch
		nach andern greift`
	*/

	key := `AZF`

	prepped_clear := prep(cleartext)
	prepped_key := prep(key)

	fmt.Printf("prepped key: %s\n\n", string(prepped_key))

	if args[0] == "encode" {
		encode(prepped_clear, prepped_key)
	}
}

func encode(msg []rune, key []rune) []rune {
	chiffre := []rune{}
	for i, in := range msg {

		ch := (((in - 64) + (key[i] - 64)) % 26) + 64
		chiffre = append(chiffre, ch)
	}

	fmt.Printf("chiffre: %s\n\n", string(chiffre))

	return chiffre
}

func decode(chiffre []rune, key []rune) []rune {
	clear := []rune{}
	for i, in := range chiffre {

		ch := (((in - 64) - (key[i] - 64)) + 26) + 64
		clear = append(clear, ch)
	}

	fmt.Printf("chiffre: %s\n\n", string(clear))

	return clear
}

func prep(t string) []rune {
	// convert to upper
	t = strings.ToUpper(t)

	runes := []rune(t)
	out := []rune{}

	for _, r := range runes {
		if r > 64 && r < 91 {
			out = append(out, r)
		}
	}

	return out
}
