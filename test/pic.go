package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var pic [][]uint8 = make([][]uint8, 0)
    for y := 0; y < dy; y++ {
        var line = make([]uint8, 0)
        for x := 0; x < dx; x++ {
            line = append(line, x * y)
        }
        pic = append(pic, line)
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
