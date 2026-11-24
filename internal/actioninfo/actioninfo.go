package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает типы, которые умеют разбирать строку
// с данными об активности и формировать отчёт по ней.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит информацию обо всех активностях из dataset.
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		if err := dp.Parse(datastring); err != nil {
			log.Printf("ошибка парсинга %q: %v", datastring, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования отчёта: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
