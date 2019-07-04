package story

import (
    "encoding/json"
    "os"

    "gophercise/adventure/utils"
)

type Arc struct {
    Title   string   `json:"title"`
    Story   []string `json:"story"`
    Options []struct {
        Text    string `json:"text"`
        DestArc string `json:"arc"`
    } `json:"options"`
}

type Arcs map[string]Arc

func ParseStory() (ArcList *Arcs) {
    file, err := os.Open("./story.json")
    utils.CheckError(err)
    decoder := json.NewDecoder(file)
    decoder.Decode(&ArcList)
    return
}
