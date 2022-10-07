package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// Constraints
// - Every clock has center (150, 150)
// Hour hand is 50 long
// Minute hand is 80 long
// Second hand is 90 long

const hrHandLen = 50
const minHandLen = 80
const secHandLen = 90

const (
    secondsInHalfClock = 30
    secondsInClock = 2 * secondsInHalfClock

    minutesInHalfClock = 30
    minutesInClock = 2 * minutesInHalfClock

    hoursInHalfClock = 6
    hoursInClock = 2 * hoursInHalfClock
)


type Point struct {
	X float64
	Y float64
}

var clockCenter = Point{150, 150}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
    hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hrHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p Point, handLen float64) Point {
	p = Point{p.X * handLen, p.Y * handLen}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenter.X, p.Y + clockCenter.Y}
    return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func secondsInRadians(t time.Time) float64 {
	seconds := float64(t.Second())
	return math.Pi / (secondsInHalfClock / seconds)
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
    minutes := float64(t.Minute())
	return (secondsInRadians(t) / secondsInClock) + (math.Pi / (minutesInHalfClock / minutes))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
    hours := float64(t.Hour() % hoursInClock)
    return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / hours))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
    x := math.Sin(angle)
    y := math.Cos(angle)
    return Point{x, y}
}
