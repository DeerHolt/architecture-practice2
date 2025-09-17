// Пакет spentcalories формирует сводку о потраченных калориях после активности на основе физиологических данных.
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

// WalkingSpentCalories возвращает информацию о потраченных калориях после ходьбы.
// При неположительных или нулевых физиологических параметрах, выводим нулевое значение
// и записываем ошибку.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration is not positive")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	resultCalories := (averageSpeed * weight * durationInMinutes) / minInH * walkingCaloriesCoefficient
	return resultCalories, nil
}

// RunningSpentCalories возвращает информацию о потраченных калориях после бега.
// При неположительных или нулевых физиологических параметрах, выводим нулевое значение
// и записываем ошибку.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration is not positive")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	resultCalories := (averageSpeed * weight * durationInMinutes) / minInH
	return resultCalories, nil
}

// MeanSpeed возвращает информацию о средней скорости после тренировки.
// При неположительных или нулевых физиологических параметрах, выводим нулевое значение.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	totalDistance := Distance(steps, height)
	if totalDistance <= 0 {
		return 0.0
	}

	return totalDistance / duration.Hours()
}

// Distance возвращает общую дистанцию после тренировки.
// При неположительных или нулевых физиологических параметрах, выводим нулевое значение.
func Distance(steps int, height float64) float64 {
	if height <= 0 {
		return 0.0
	}
	stepLength := height * stepLengthCoefficient
	if steps <= 0 {
		return 0.0
	}

	totalDistance := float64(steps) * stepLength
	return totalDistance / mInKm
}
