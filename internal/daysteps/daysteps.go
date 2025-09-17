// Пакет daysteps формирует сводку после дневной прогулки
package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse заполняет данные о прошедшей тренировке.
// При недействительных значениях выводит ошибку.
func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("data is invalid")
	}

	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		log.Println(err)
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("steps must be greater than 0")
	}

	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		log.Println(err)
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("duration is not positive")
	}

	return nil
}

// ActionInfo выодит результаты о прошедшей тренировку на основе полученных данных.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
