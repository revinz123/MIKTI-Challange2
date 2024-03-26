package main

import (
    "fmt"
)

func main() {
    // Data 1
    markWeight1 := 78.0  // kg
    markHeight1 := 1.69  // m

    johnWeight1 := 92.0  // kg
    johnHeight1 := 1.95  // m

    // Menghitung BMI (versi 1)
    markBMI1 := markWeight1 / (markHeight1 * markHeight1)
    johnBMI1 := johnWeight1 / (johnHeight1 * johnHeight1)

    // Menghitung BMI (versi 2)
    markBMI2 := markWeight1 / (markHeight1 * markHeight1)
    johnBMI2 := johnWeight1 / (johnHeight1 * johnHeight1)

    // Membuat variabel Boolean 'markHigherBMI'
    markHigherBMI1 := markBMI1 > johnBMI1
    markHigherBMI2 := markBMI2 > johnBMI2

    // Menampilkan hasil
    fmt.Println("Data 1:")
    fmt.Printf("BMI Mark (versi 1): %.2f\n", markBMI1)
    fmt.Printf("BMI John (versi 1): %.2f\n", johnBMI1)
    fmt.Println("Mark memiliki BMI lebih tinggi dari John (versi 1):", markHigherBMI1)

    fmt.Printf("\nBMI Mark (versi 2): %.2f\n", markBMI2)
    fmt.Printf("BMI John (versi 2): %.2f\n", johnBMI2)
    fmt.Println("Mark memiliki BMI lebih tinggi dari John (versi 2):", markHigherBMI2)

    // Data 2
    markWeight2 := 95.0  // kg
    markHeight2 := 1.88  // m

    johnWeight2 := 85.0  // kg
    johnHeight2 := 1.76  // m

    // Menghitung BMI (versi 1)
    markBMI1 = markWeight2 / (markHeight2 * markHeight2)
    johnBMI1 = johnWeight2 / (johnHeight2 * johnHeight2)

    // Menghitung BMI (versi 2)
    markBMI2 = markWeight2 / (markHeight2 * markHeight2)
    johnBMI2 = johnWeight2 / (johnHeight2 * johnHeight2)

    // Membuat variabel Boolean 'markHigherBMI'
    markHigherBMI1 = markBMI1 > johnBMI1
    markHigherBMI2 = markBMI2 > johnBMI2

    // Menampilkan hasil
    fmt.Println("\nData 2:")
    fmt.Printf("BMI Mark (versi 1): %.2f\n", markBMI1)
    fmt.Printf("BMI John (versi 1): %.2f\n", johnBMI1)
    fmt.Println("Mark memiliki BMI lebih tinggi dari John (versi 1):", markHigherBMI1)

    fmt.Printf("\nBMI Mark (versi 2): %.2f\n", markBMI2)
    fmt.Printf("BMI John (versi 2): %.2f\n", johnBMI2)
    fmt.Println("Mark memiliki BMI lebih tinggi dari John (versi 2):", markHigherBMI2)
}
