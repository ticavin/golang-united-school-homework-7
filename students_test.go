package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestLen(t *testing.T) {
	var persons People
	persons = append(persons, Person{"Monk", "Vlodkar", time.Now()})
	persons = append(persons, Person{"Ivan", "Petrov", time.Now()})
	var expected = 2

	length := persons.Len()

	if expected != length {
		t.Errorf("Expected %d got %d", expected, length)
	}
}

func TestLessBirthday(t *testing.T) {
	var persons People
	persons = append(persons, Person{"same", "lastsame", time.Now()})
	persons = append(persons, Person{"same", "lastsame", time.Now().AddDate(2, 5, 21)})

	less := persons.Less(0, 1)

	if less != false {
		t.Errorf("Expected %t, but got %t", false, less)
	}
}
func TestLessLastName(t *testing.T) {
	var persons People
	persons = append(persons, Person{"same", "lastsame", time.Now()})
	persons = append(persons, Person{"same", "notlastsame", time.Now()})

	less := persons.Less(0, 1)

	if less != true {
		t.Errorf("Expected %t, but got %t", false, less)
	}
}
func TestLessFirstName(t *testing.T) {
	var persons People
	persons = append(persons, Person{"same", "lastsame", time.Now()})
	persons = append(persons, Person{"notsame", "lastsame", time.Now()})

	less := persons.Less(0, 1)

	if less != false {
		t.Errorf("Expected %t, but got %t", true, less)
	}
}

func TestSwap(t *testing.T) {
	var person1 = Person{"Monk", "Vlodkar", time.Now()}
	var person2 = Person{"Ivan", "Petrov", time.Now()}
	persons := People{person1, person2}

	persons.Swap(0, 1)

	if persons[0].firstName != person2.firstName {
		t.Errorf("Expected %+v, but got %+v", person2, persons[0])
	}
}

func TestNewMatrix(t *testing.T) {
	matrixStr := "1 2 3\n5 6 7\n9 10 11"
	matrix, err := New(matrixStr)

	if err != nil {
		t.Errorf("Unexpected error")
	}

	if matrix.rows != 3 {
		t.Error("Rows must be 3")
	}
	if matrix.cols != 3 {
		t.Error("Cols must be 3")
	}
	if matrix.cols != matrix.rows {
		t.Error("Cols must be 3")
	}
	if matrix.data[3] != 5 {
		t.Error("wrong data")
	}
}
func TestMatrixWithStringError(t *testing.T) {
	_, err := New("DOG 2 3\n 4 5 6\n 7 8 9")

	if err == nil {
		t.Errorf("matrix with word")
	}
}
func TestMatrixIfEmpty(t *testing.T) {
	_, err := New("")

	if err == nil {
		t.Errorf("mpty param ")
	}
}

func TestRows(t *testing.T) {
	matrixStr := "1 2 3 4\n5 6 7 8\n9 10 11 12"
	matrix, _ := New(matrixStr)

	rowMatrix := matrix.Rows()

	if rowMatrix[1][1] != 6 {
		t.Errorf("Expected to be 6, but got %d", rowMatrix[1][1])
	}
}

func TestCols(t *testing.T) {
	matrixStr := "1 2 3 4\n5 6 7 8\n9 10 11 12"
	matrix, _ := New(matrixStr)

	colMatrix := matrix.Cols()

	if colMatrix[1][2] != 10 {
		t.Errorf("Expected to be 10, but got %d", colMatrix[1][2])
	}
}

func TestSet(t *testing.T) {
	matrixStr := "1 2 3 4\n5 6 7 8\n9 10 11 12"
	matrix, _ := New(matrixStr)

	set := matrix.Set(1, 0, 69)

	if set && matrix.data[4] != 69 {
		t.Errorf("Expected value 99, but got %d", matrix.data[4])
	}
}
func TestSetWithNegativeRow(t *testing.T) {
	matrixStr := "1 2 3 4\n5 6 7 8\n9 10 11 12"
	matrix, _ := New(matrixStr)

	var set bool = matrix.Set(-1, 0, -11)

	if set {
		t.Errorf("Negative row")
	}
}
func TestSetLargeColAndRow(t *testing.T) {
	matrixStr := "1 2 3 4\n5 6 7 8\n9 10 11 12"
	matrix, _ := New(matrixStr)

	wrongRow := matrix.Set(9, 1, 55)
	wrongCol := matrix.Set(1, 9, 88)

	if wrongRow {
		t.Error("Row cant be greater then matrix")
	}

	if wrongCol {
		t.Error("Col cant be greeter then matrix cols")
	}
}
