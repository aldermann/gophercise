package context

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "gophercise/adventure/story"
    "gophercise/adventure/utils"
)

var reader = bufio.NewReader(os.Stdin)

type Context struct {
    CurrentArc string
    ArcList    *story.Arcs
}

func (c *Context) PrintStory() {
    utils.ClearScreen()
    arc := (*c.ArcList)[c.CurrentArc]
    fmt.Println(arc.Title)
    for _, s := range arc.Story {
        fmt.Println(s)
    }
}

func (c *Context) PresentOptions() int {
    arc := (*c.ArcList)[c.CurrentArc]
    if len(arc.Options) != 0 {
        fmt.Println("Here are the options:")
    }
    for i, v := range arc.Options {
        fmt.Printf("%d: %s\n", i+1, v.Text)
    }
    return len(arc.Options)
}

func (c *Context) QueryOption() {
    arc := (*c.ArcList)[c.CurrentArc]
    cnt := len(arc.Options)
    for {
        fmt.Print("Please choose your path: ")
        input, err := reader.ReadString('\n')
        utils.CheckError(err)
        input = strings.TrimSpace(input)
        id, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input")
        } else if id <= 0 || id > cnt {
            fmt.Println("Out of range")
        } else {
            c.moveToArc(id - 1)
            break
        }
    }
}

func (c *Context) moveToArc(id int) {
    arc := (*c.ArcList)[c.CurrentArc]
    c.CurrentArc = arc.Options[id].DestArc
}
