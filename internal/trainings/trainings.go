// Пакет trainings формирует сводку после тренировки.
package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse заполняет данные о прошедшей тренировке.
// При недействительных значениях выводит ошибку.
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("data is invalid")
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		log.Println(err)
		return err
	}
	if t.Steps <= 0 {
		return errors.New("steps must be greater than 0")
	}

	t.TrainingType = data[1]
	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		log.Println(err)
		return err
	}
	if t.Duration <= 0 {
		return errors.New("duration is not positive")
	}

	return nil
}

// ActionInfo выодит результаты о прошедшей тренировку на основе полученных данных.
func (t Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)
	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, averageSpeed, calories)
	return result, nil
}
