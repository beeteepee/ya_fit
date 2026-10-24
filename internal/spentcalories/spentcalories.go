package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataParts := strings.Split(data, ",")
	if len(dataParts) != 3 {
		return 0, "", 0, errors.New("invalid input format")
	}
	steps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return 0, "", 0, errors.New("Некорректное количество шагов")
	}
	trainingType := dataParts[1]
	duration, err := time.ParseDuration(dataParts[2])
	if err != nil {
		return 0, "", 0, errors.New("Некорректное время")
	}
	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient / mInKm
	return float64(steps) * lenStep
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration.Minutes() == 0 {
		return 0
	}
	averageSpeed := distance(steps, height) / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	switch trainingType {
	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: Ходьба\n Длительность: %v\n Дистанция: %.2f км Шаги: %d\nСкорость: %.2f км/ч\n Сожгли калорий: %.2f",
			duration, distance(steps, height), steps, meanSpeed(steps, height, duration), calories), nil
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: Бег\n Длительность: %v\n Дистанция: %.2f км Шаги: %d\nСкорость: %.2f км/ч\n Сожгли калорий: %.2f",
			duration, distance(steps, height), steps, meanSpeed(steps, height, duration), calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Некорректные входные данные")
	}
	minutes := duration.Minutes()
	calories := (weight * meanSpeed(steps, height, duration) * minutes) / minInH
	return calories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Некорректные входные данные")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := ((weight * averageSpeed * minutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
