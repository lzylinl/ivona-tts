package main
import (
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	ivona "github.com/jpadilla/ivona-go"
)
func main() {
	client := ivona.New("GDNAICTDMLSLU5426OAA", "2qUFTF8ZF9wqy7xoGBY+YXLEu+M2Qqalf/pSrd9m")
	text, err := ioutil.ReadFile("/home/vitaly/Desktop/article.txt")
	if err != nil {
		log.Fatal(err)
	}

	arrayOfParagraphs := strings.Split(string(text), "\n\n")
	i := 0
	for _,paragraph := range arrayOfParagraphs {
		paragraph = strings.TrimSpace(paragraph)
		if (len(paragraph) < 1) { // against empty lines
			continue
		}
		log.Printf("%s", paragraph)
		options := ivona.NewSpeechOptions(paragraph)
		options.Voice.Language = "ru-RU"
		options.Voice.Name = "Maxim"
		options.Voice.Gender = "Male"
		options.OutputFormat.Codec = "OGG"
		r, err := client.CreateSpeech(options)
		if err != nil {
			log.Fatal(err)
		}

		i++
		file := fmt.Sprintf("/home/vitaly/Desktop/ivona-article/tts%04d.ogg", i) // files like 0001.ogg
		ioutil.WriteFile(file, r.Audio, 0644)
	}
}
