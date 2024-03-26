package main

import (
    "fmt"
)

func averageScore(scores []int) float64 {
    total := 0
    for _, score := range scores {
        total += score
    }
    return float64(total) / float64(len(scores))
}

func main() {
    // Data 1
    dolphinScores := []int{96, 108, 89}
    koalaScores := []int{88, 91, 110}

    dolphinAverage := averageScore(dolphinScores)
    koalaAverage := averageScore(koalaScores)

    // Skor minimum
    const minimumScore = 100

    // Bonus 1: Skor minimum juga berlaku untuk seri
    if dolphinAverage >= minimumScore && koalaAverage >= minimumScore {
        if dolphinAverage > koalaAverage {
            fmt.Println("Pemenang: Lumba-lumba")
        } else if koalaAverage > dolphinAverage {
            fmt.Println("Pemenang: Koala")
        } else {
            fmt.Println("Seri")
        }
    } else {
        fmt.Println("Tidak ada pemenang karena skor rata-rata kurang dari", minimumScore)
    }

    // Data Bonus 1
    dolphinScoresBonus1 := []int{97, 112, 101}
    koalaScoresBonus1 := []int{109, 95, 123}

    dolphinAverageBonus1 := averageScore(dolphinScoresBonus1)
    koalaAverageBonus1 := averageScore(koalaScoresBonus1)

    // Bonus 1: Skor minimum juga berlaku untuk seri
    if dolphinAverageBonus1 >= minimumScore && koalaAverageBonus1 >= minimumScore {
        if dolphinAverageBonus1 > koalaAverageBonus1 {
            fmt.Println("Pemenang Bonus 1: Lumba-lumba")
        } else if koalaAverageBonus1 > dolphinAverageBonus1 {
            fmt.Println("Pemenang Bonus 1: Koala")
        } else {
            fmt.Println("Seri Bonus 1")
        }
    } else {
        fmt.Println("Tidak ada pemenang karena skor rata-rata kurang dari", minimumScore)
    }

    // Data Bonus 2
    dolphinScoresBonus2 := []int{97, 112, 101}
    koalaScoresBonus2 := []int{109, 95, 106}

    dolphinAverageBonus2 := averageScore(dolphinScoresBonus2)
    koalaAverageBonus2 := averageScore(koalaScoresBonus2)

    // Bonus 2: Skor minimum juga berlaku untuk seri
    if dolphinAverageBonus2 >= minimumScore && koalaAverageBonus2 >= minimumScore {
        if dolphinAverageBonus2 > koalaAverageBonus2 {
            fmt.Println("Pemenang Bonus 2: Lumba-lumba")
        } else if koalaAverageBonus2 > dolphinAverageBonus2 {
            fmt.Println("Pemenang Bonus 2: Koala")
        } else {
            fmt.Println("Seri Bonus 2")
        }
    } else {
        fmt.Println("Tidak ada pemenang karena skor rata-rata kurang dari", minimumScore)
    }
}
