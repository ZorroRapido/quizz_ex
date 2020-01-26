package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}

type Question struct {
	question string
	answer   string
}

func readCsv(filename string) []Question {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	//fmt.Println(records[0])

	questions := []Question{}

	for _, i := range records {
		questions = append(questions, Question{i[0], i[1]})
	}

	return questions
}

func main() {
	total := 0

	questions := readCsv("problems.csv")
	//[]Question{{"1 + 1", "2"}, {"2 + 2", "4"}}

	var answer string

	for _, question := range questions {
		fmt.Println(question.question)

		fmt.Scan(&answer)

		if answer == question.answer {
			total++
		}
	}

	fmt.Printf("You got %d/%d", total, len(questions))
	// Пройтись циклом. Вывести вопрос, предложить пользователю ввести ответ.
	// Если ответ правильный, увеличить total.
	// for
}
