package cli

import (
    "fmt"

    "gophercise/adventure/cli/context"
    "gophercise/adventure/story"
)

func Run() {
    ArcList := story.ParseStory()
    c := context.Context{CurrentArc: "intro", ArcList: ArcList}

    for {
        c.PrintStory()
        if c.PresentOptions() == 0 {
            break
        }
        c.QueryOption()
    }

    fmt.Println("The end")
}
