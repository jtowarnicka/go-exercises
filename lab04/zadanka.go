package main

import (
	"fmt"
	"math/rand"
)

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func generateMatrixSlice(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = randomFloat(0, float64(10))
		}
	}
	return matrix
}

func sumMatrixes(matrix1, matrix2 [][]float64) [][]float64 {
	rows := len(matrix1)
	cols := len(matrix1[0])
	matrix := make([][]float64, rows, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	return matrix
}

func multiplyMatrixes(matrix1, matrix2 [][]float64) [][]float64 {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])
	if cols1 != rows2 {
		panic("Liczebność kolumn pierwszej macierzy nie jest równa liczbie wierszy drugiej macierzy.")
	}

	result := make([][]float64, rows1)

	for i := 0; i < rows1; i++ {
		result[i] = make([]float64, cols2)
		for j := 0; j < cols2; j++ {
			for k := 0; k < rows2; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result

}

func main() {
	var tablica1 [20]float64;
	var tablica2 [20]float64;
	var tablica3 [20]float64;
	var tablica4 [20]float64;

	for i := 0; i < 20; i++ {
		tablica1[i] = 2.0;
		tablica2[i] = 3.0;
		tablica3[i] = tablica1[i] + tablica2[i];
		tablica4[i] = tablica1[i] * tablica2[i];
	}

	fmt.Println("Tablica 1: ", tablica1);
	fmt.Println("Tablica 2: ", tablica2);
	fmt.Println("Tablica 3: ", tablica3);
	fmt.Println("Tablica 4: ", tablica4);

	wycinek := make([][]float64, 3, 3);
	for i := 0; i < 3; i++ {
		wycinek[i] = make([]float64, 3, 3);
		for j := 0; j < 3; j++ {
			wycinek[i][j] = float64(i * 3 + j);
		}
	}
	wycinek2 := [][]float64{{8.0, 7.0, 6.0}, {5.0, 4.0, 3.0}, {2.0, 1.0, 0.0}};
	fmt.Println("Wycinek 1: ", wycinek);
	fmt.Println("Wycinek 2: ", wycinek2);

	wycinek3 := make([][]float64, 3, 3);
	for i := 0; i < 3; i++ {
		wycinek3[i] = make([]float64, 3, 3);
		for j := 0; j < 3; j++ {
			wycinek3[i][j] = wycinek[i][j] + wycinek2[i][j]
		}
	}
	fmt.Println("Wycinek 3: ", wycinek3);

	macierz1 := generateMatrixSlice(4, 4);
	macierz2 := generateMatrixSlice(4, 4);
	fmt.Println("Macierz 1: ", macierz1);
	fmt.Println("Macierz 2: ", macierz2);

	fmt.Println("Dodawanie macierzy: ");
	result := sumMatrixes(macierz1, macierz2);
	fmt.Println(result)

	mtx1 := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	mtx2 := [][]float64{{7.0, 8.0}, {9.0, 10.0}, {11.0, 12.0}}
	fmt.Println("Mnożenie macierzy: ");
	result2 := multiplyMatrixes(mtx2, mtx1);
	fmt.Println(result2)
}
