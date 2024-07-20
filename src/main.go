package main

import "fmt"

const ebulicaoK float64  = 373.2

func main() {
    var tempC float64 = ebulicaoK - 273.15
    fmt.Printf("A temperatuda de ebulicao em °K: %.2f°K\nA temperatura de ebulicao em °C: %.2f°C\n",ebulicaoK , tempC)   
}
