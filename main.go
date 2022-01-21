package main

import (
	"ChooseYourOwnAdventure/story"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "story.json", "The JSON file with the story")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := story.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
