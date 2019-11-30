package workshop

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Add(numbers string) (result string) {
	if numbers == "" {
		return "0"
	}

	if numbers == "1" {
		return "1"
	}

	subs, err := split(numbers)
	if err != nil {
		return err.Error()
	}
	parsed := parse(subs...)
	sumRes, err := sum(parsed...)
	if err != nil {
		return err.Error()
	}
	result = toString(sumRes)

	return
}

func split(numbers string) (subs []string, err error) {
	if strings.HasSuffix(numbers, ",") {
		err = fmt.Errorf("Number expected but EOF found")
		return
	}
	if strings.HasPrefix(numbers, "//") {
		return splitWithDelimiter(numbers)
	}
	return splitDefault(numbers)
}

func splitDefault(numbers string) (subs []string, err error) {
	var previous rune
	var position int
	subs = strings.FieldsFunc(numbers, func(c rune) bool {
		if err != nil {
			return false
		}
		defer func() {
			previous = c
			position++
		}()
		if c == ',' || c == '\n' {
			if previous == ',' && c == '\n' {
				err = fmt.Errorf("Number expected but '\n' found at position %d", position)
			}
			return true
		}
		return false
	})
	return
}

func splitWithDelimiter(numbers string) (subs []string, err error) {
	trimedNumbers := strings.TrimPrefix(numbers, "//")
	parts := strings.SplitN(trimedNumbers, "\n", 2)
	delimiter := parts[0]
	rest := parts[1]
	subs = strings.Split(rest, delimiter)
	if delimiter != "," {
		for _, s := range subs {
			if !strings.Contains(s, ",") {
				continue
			}
			index := strings.Index(rest, ",")
			err = fmt.Errorf("'%s' expected but ',' found at position %d", delimiter, index)
			return
		}
	}
	return
}

func parse(subs ...string) (result []float64) {
	for _, s := range subs {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, f)
	}
	return
}

func sum(numbers ...float64) (result float64, err error) {
	var negatives []float64
	for _, n := range numbers {
		if n < 0 {
			negatives = append(negatives, n)
			continue
		}
		result += n
	}
	err = checkNegatives(negatives...)
	return
}

func checkNegatives(negatives ...float64) error {
	if len(negatives) == 0 {
		return nil
	}
	errorMessage := fmt.Sprintf("Negative not allowed : %v", negatives[0])
	for _, n := range negatives[1:] {
		errorMessage = fmt.Sprintf(errorMessage+", %v", n)
	}
	return fmt.Errorf(errorMessage)
}

func toString(n float64) string {
	if n == math.Floor(n) {
		return fmt.Sprint(int64(n))
	}
	return fmt.Sprintf("%.1f", n)
}
