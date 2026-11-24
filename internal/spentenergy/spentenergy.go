package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories считает калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories считает калории, потраченные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func validateInput(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	return nil
}
