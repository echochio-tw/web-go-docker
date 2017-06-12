package main

import (
        "fmt"
        "net/http"
        "github.com/ajstarks/svgo"
        "math/rand"
        "time"
        "log"
)

const arcfmt = "stroke:%s;stroke-opacity:%.2f;stroke-width:%dpx;fill:none"

var colors = []string{"red", "green", "blue", "white", "gray"}

func randarc(canvas *svg.SVG, aw, ah, sw int, f1, f2 bool) {
        begin := rand.Intn(aw)
        arclength := rand.Intn(aw)
        end := begin + arclength
        baseline := ah / 2
        al := arclength / 2
        cl := len(colors)
        canvas.Arc(begin, baseline, al, al, 0, f1, f2, end, baseline,
                fmt.Sprintf(arcfmt, colors[rand.Intn(cl)], rand.Float64(), rand.Intn(sw)))

}

func main() {
        http.Handle("/", http.HandlerFunc(circle))
        err := http.ListenAndServe(":9090", nil)
        if err != nil {
                log.Fatal("ListenAndServe:", err)
        }
}

func circle(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")

        rand.Seed(time.Now().UnixNano())

        width := 800
        height := 800
        aw := width / 2
        maxstroke := height / 10
        literactions := 50
        canvas := svg.New(w)
        canvas.Start(width, height)
        canvas.Rect(0, 0, width, height)
        for i := 0; i < literactions; i++ {
                randarc(canvas, aw, height, maxstroke, false, true)
                randarc(canvas, aw, height, maxstroke, false, false)
        }
        canvas.End()
}

