package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println(os.Args[1])
	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("CSF111_202425_01_GradeBook")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows[1][4])
	numberOfStudents := len(rows) - 1 // No of students in A group
	// Will store the student's total marks of every component of the course
	studentMarks := make([]float64, numberOfStudents) // 0th index is irrelevant
	quizavg := 0.0
	midsemavg := 0.0
	labtestavg := 0.0
	weeklylabavg := 0.0
	finalavg := 0.0
	precompreevalavg := 0.0
	compreavg := 0.0
	branchwisemarks := make(map[string][]float64)
	var quizrank [3][11]string
	for i := 0; i < 11; i++ {
		quizrank[0][i] = rows[1][i]
		quizrank[1][i] = rows[1][i]
		quizrank[2][i] = rows[1][i]
	}
	var midsemrank [3][11]string
	for i := 0; i < 11; i++ {
		midsemrank[0][i] = rows[1][i]
		midsemrank[1][i] = rows[1][i]
		midsemrank[2][i] = rows[1][i]
	}
	var labtestrank [3][11]string
	for i := 0; i < 11; i++ {
		labtestrank[0][i] = rows[1][i]
		labtestrank[1][i] = rows[1][i]
		labtestrank[2][i] = rows[1][i]
	}
	var weeklylabrank [3][11]string
	for i := 0; i < 11; i++ {
		weeklylabrank[0][i] = rows[1][i]
		weeklylabrank[1][i] = rows[1][i]
		weeklylabrank[2][i] = rows[1][i]
	}
	var comprerank [3][11]string
	for i := 0; i < 11; i++ {
		comprerank[0][i] = rows[1][i]
		comprerank[1][i] = rows[1][i]
		comprerank[2][i] = rows[1][i]
	}
	var finalrank [3][11]string
	for i := 0; i < 11; i++ {
		finalrank[0][i] = rows[1][i]
		finalrank[1][i] = rows[1][i]
		finalrank[2][i] = rows[1][i]
	}
	var precomprerank [3][11]string
	for i := 0; i < 11; i++ {
		precomprerank[0][i] = rows[1][i]
		precomprerank[1][i] = rows[1][i]
		precomprerank[2][i] = rows[1][i]
	}

	for i := 1; i < numberOfStudents; i++ {
		studentMarks[i] = 0
		for j := 4; j < 10; j++ {
			if j == 8 {
				continue
			} // To avoid adding precompre total two times in the total marks
			a, err := strconv.ParseFloat(rows[i][j], 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			studentMarks[i] += a
		}
		studentMarks[i] = math.Round(studentMarks[i]*100) / 100 // As some values have a lot of decimal places , rounding off to 2 decimal places
		b, err := strconv.ParseFloat(rows[i][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if studentMarks[i] != math.Round(b*100)/100 {
			fmt.Println("Error spotted in the calculation of Total in row: #", i+1)
		}
		quizMarks, err := strconv.ParseFloat(rows[i][4], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		midsemMarks, err := strconv.ParseFloat(rows[i][5], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		labtestMarks, err := strconv.ParseFloat(rows[i][6], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		weeklylabMarks, err := strconv.ParseFloat(rows[i][7], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		compreevalMarks, err := strconv.ParseFloat(rows[i][9], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		precompreMarks, err := strconv.ParseFloat(rows[i][8], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		quizr1, err := strconv.ParseFloat(quizrank[0][4], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		quizr2, err := strconv.ParseFloat(quizrank[1][4], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		quizr3, err := strconv.ParseFloat(quizrank[2][4], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		midsemr1, err := strconv.ParseFloat(midsemrank[0][5], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		midsemr2, err := strconv.ParseFloat(midsemrank[1][5], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		midsemr3, err := strconv.ParseFloat(midsemrank[2][5], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		labtestr1, err := strconv.ParseFloat(labtestrank[0][6], 64)
		if err != nil {
			fmt.Println(err)
			return

		}
		labtestr2, err := strconv.ParseFloat(labtestrank[1][6], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		labtestr3, err := strconv.ParseFloat(labtestrank[2][6], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		weeklylabr1, err := strconv.ParseFloat(weeklylabrank[0][7], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		weeklylabr2, err := strconv.ParseFloat(weeklylabrank[1][7], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		weeklylabr3, err := strconv.ParseFloat(weeklylabrank[2][7], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		comprer1, err := strconv.ParseFloat(comprerank[0][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		comprer2, err := strconv.ParseFloat(comprerank[1][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		comprer3, err := strconv.ParseFloat(comprerank[2][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		finalr1, err := strconv.ParseFloat(finalrank[0][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		finalr2, err := strconv.ParseFloat(finalrank[1][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		finalr3, err := strconv.ParseFloat(finalrank[1][10], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		precomprerank1, err := strconv.ParseFloat(precomprerank[0][8], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		precomprerank2, err := strconv.ParseFloat(precomprerank[1][8], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		precomprerank3, err := strconv.ParseFloat(precomprerank[2][8], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if midsemMarks > midsemr1 {
			for j := 0; j < 11; j++ {
				midsemrank[2][j] = midsemrank[1][j]
				midsemrank[1][j] = midsemrank[0][j]
				midsemrank[0][j] = rows[i][j]
			}
		} else if midsemMarks > midsemr2 {
			for j := 0; j < 11; j++ {
				midsemrank[2][j] = midsemrank[1][j]
				midsemrank[1][j] = rows[i][j]
			}
		} else if midsemMarks > midsemr3 {
			for j := 0; j < 11; j++ {
				midsemrank[2][j] = rows[i][j]
			}
		}
		if labtestMarks > labtestr1 {
			for j := 0; j < 11; j++ {
				labtestrank[2][j] = labtestrank[1][j]
				labtestrank[1][j] = labtestrank[0][j]
				labtestrank[0][j] = rows[i][j]
			}
		} else if labtestMarks > labtestr2 {
			for j := 0; j < 11; j++ {
				labtestrank[2][j] = labtestrank[1][j]
				labtestrank[1][j] = rows[i][j]
			}
		} else if labtestMarks > labtestr3 {
			for j := 0; j < 11; j++ {
				labtestrank[2][j] = rows[i][j]
			}
		}
		if weeklylabMarks > weeklylabr1 {
			for j := 0; j < 11; j++ {
				weeklylabrank[2][j] = weeklylabrank[1][j]
				weeklylabrank[1][j] = weeklylabrank[0][j]
				weeklylabrank[0][j] = rows[i][j]
			}
		} else if weeklylabMarks > weeklylabr2 {
			for j := 0; j < 11; j++ {
				weeklylabrank[2][j] = weeklylabrank[1][j]
				weeklylabrank[1][j] = rows[i][j]
			}
		} else if weeklylabMarks > weeklylabr3 {
			for j := 0; j < 11; j++ {
				weeklylabrank[2][j] = rows[i][j]
			}
		}
		if compreevalMarks > comprer1 {
			for j := 0; j < 11; j++ {
				comprerank[2][j] = comprerank[1][j]
				comprerank[1][j] = comprerank[0][j]
				comprerank[0][j] = rows[i][j]
			}
		} else if compreevalMarks > comprer2 {
			for j := 0; j < 11; j++ {
				comprerank[2][j] = comprerank[1][j]
				comprerank[1][j] = rows[i][j]
			}
		} else if compreevalMarks > comprer3 {
			for j := 0; j < 11; j++ {
				comprerank[2][j] = rows[i][j]
			}
		}
		if b > finalr1 {
			for j := 0; j < 11; j++ {
				finalrank[2][j] = finalrank[1][j]
				finalrank[1][j] = finalrank[0][j]
				finalrank[0][j] = rows[i][j]
			}
		} else if b > finalr2 {
			for j := 0; j < 11; j++ {
				finalrank[2][j] = finalrank[1][j]
				finalrank[1][j] = rows[i][j]
			}
		} else if b > finalr3 {
			for j := 0; j < 11; j++ {
				finalrank[2][j] = rows[i][j]
			}
		}
		if precompreMarks > precomprerank1 {
			for j := 0; j < 11; j++ {
				precomprerank[2][j] = precomprerank[1][j]
				precomprerank[1][j] = precomprerank[0][j]
				precomprerank[0][j] = rows[i][j]
			}
		} else if precompreMarks > precomprerank2 {
			for j := 0; j < 11; j++ {
				precomprerank[2][j] = precomprerank[1][j]
				precomprerank[1][j] = rows[i][j]
			}
		} else if precompreMarks > precomprerank3 {
			for j := 0; j < 11; j++ {
				precomprerank[2][j] = rows[i][j]
			}
		}
		if quizMarks > quizr1 {
			for j := 0; j < 11; j++ {
				quizrank[2][j] = quizrank[1][j]
				quizrank[1][j] = quizrank[0][j]
				quizrank[0][j] = rows[i][j]

			}
		} else if quizMarks > quizr2 {
			for j := 0; j < 11; j++ {
				quizrank[2][j] = quizrank[1][j]
				quizrank[1][j] = rows[i][j]
			}
		} else if quizMarks > quizr3 {
			for j := 0; j < 11; j++ {
				quizrank[2][j] = rows[i][j]
			}
		}

		if rows[i][3][0:4] == "2024" {
			branchwisemarks[rows[i][3][4:6]] = append(branchwisemarks[rows[i][3][4:6]], studentMarks[i])
		}

		quizavg += math.Round(quizMarks*100) / 100
		precompreevalavg += math.Round(precompreMarks*100) / 100
		midsemavg += math.Round(midsemMarks*100) / 100
		weeklylabavg += math.Round(weeklylabMarks*100) / 100
		labtestavg += math.Round(labtestMarks*100) / 100
		compreavg += math.Round(compreevalMarks*100) / 100
		finalavg += b
	}
	fmt.Println("Branch wise average marks: ")
	for key, value := range branchwisemarks {
		var marks float64 = 0
		for _, val := range value {
			marks += val
		}
		fmt.Println("Branch: ", key, " Average: ", marks/float64(len(value)))
	}
	quizavg /= float64(numberOfStudents)
	midsemavg /= float64(numberOfStudents)
	labtestavg /= float64(numberOfStudents)
	weeklylabavg /= float64(numberOfStudents)
	finalavg /= float64(numberOfStudents)
	precompreevalavg /= float64(numberOfStudents)
	compreavg /= float64(numberOfStudents)
	var expectedtotalavg float64 = quizavg + midsemavg + labtestavg + weeklylabavg + compreavg
	fmt.Println("Expected Total Average: ", expectedtotalavg)
	fmt.Println("Quiz Average: ", quizavg)
	fmt.Println("Midsem Average: ", midsemavg)
	fmt.Println("Lab Test Average: ", labtestavg)
	fmt.Println("Weekly Lab Average: ", weeklylabavg)
	fmt.Println("Final Average: ", finalavg)
	fmt.Println("Precompre Evaluation Average: ", precompreevalavg)
	fmt.Println("Compre Average: ", compreavg)
	fmt.Println("Quiz Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", quizrank[i][4], "ID: ", quizrank[i][3], "Emplid: ", quizrank[i][2])
	}
	fmt.Println("Midsem Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", midsemrank[i][5], "ID: ", midsemrank[i][3], "Emplid: ", midsemrank[i][2])
	}
	fmt.Println("Lab Test Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", labtestrank[i][6], "ID: ", labtestrank[i][3], "Emplid: ", labtestrank[i][2])
	}
	fmt.Println("Weekly Lab Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", weeklylabrank[i][7], "ID: ", weeklylabrank[i][3], "Emplid: ", weeklylabrank[i][2])
	}
	fmt.Println("Compre Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", comprerank[i][10], "ID: ", comprerank[i][3], "Emplid: ", comprerank[i][2])
	}
	fmt.Println("Final Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", finalrank[i][10], "ID: ", finalrank[i][3], "Emplid: ", finalrank[i][2])
	}
	fmt.Println("Precompre Rank: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Rank #", i+1, "Marks: ", precomprerank[i][8], "ID: ", precomprerank[i][3], "Emplid: ", precomprerank[i][2])
	}
}
