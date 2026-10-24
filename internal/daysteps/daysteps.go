package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	sc "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("не хватает данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("некорректное количество шагов")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, errors.New("некорректное время")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return "Ошибка: " + err.Error()
	}

	distance := float64(steps) * stepLength / float64(mInKm)
	calories, err := sc.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return "Ошибка: " + err.Error()
	}

	return fmt.Sprintf(
		"Количество шагов: %d\nДистанция: %.2f \nСожгли калорий: %.2f",
		steps,
		distance,
		calories,
	)

}
