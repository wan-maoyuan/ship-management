package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"ship-management/pkg/config"
	"ship-management/pkg/database"
	"ship-management/pkg/model"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	FirstCategoryPath  = "./resources/data/pms_equipment_first_category.csv"
	SecondCategoryPath = "./resources/data/pms_equipment_second_category.csv"
	ThirdCategoryPath  = "./resources/data/pms_equipment_third_category.csv"
	FourthCategoryPath = "./resources/data/pms_equipment_fourth_category.csv"
)

func init() {
	c, err := config.New()
	if err != nil {
		logrus.Fatalf("配置信息加载失败: %v", err)
	}

	c.Show()
}

func main() {
	if err := database.InitDatabase(config.Get().Server.SqlserverDsn); err != nil {
		logrus.Fatalf("数据库初始化失败: %v", err)
	}

	if err := importFirstCategory(); err != nil {
		logrus.Errorf("importFirstCategory: %v", err)
	}

	if err := importSecondCategory(); err != nil {
		logrus.Errorf("importSecondCategory: %v", err)
	}

	if err := importThirdCategory(); err != nil {
		logrus.Errorf("importThirdCategory: %v", err)
	}

	if err := importFourthCategory(); err != nil {
		logrus.Errorf("importFourthCategory: %v", err)
	}
}

func importFirstCategory() error {
	file, err := os.Open(FirstCategoryPath)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	for _, line := range lines[1:] {
		err = database.InsertPMSEquipmentFirstCategory(&model.PMSEquipmentFirstCategory{
			FirstCode:   line[0],
			Description: line[1],
			Editor:      "system",
			Source:      "system",
			ModifyTime:  time.Now().UTC(),
			CreateTime:  time.Now().UTC(),
		})

		if err != nil {
			logrus.Errorf("data: %v insert database error: %v", line, err)
		}
	}

	return nil
}

func importSecondCategory() error {
	file, err := os.Open(SecondCategoryPath)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	for _, line := range lines[1:] {
		err = database.InsertPMSEquipmentSecondCategory(&model.PMSEquipmentSecondCategory{
			FirstCode:   line[2],
			SecondCode:  line[0],
			Description: line[1],
			Editor:      "system",
			Source:      "system",
			ModifyTime:  time.Now().UTC(),
			CreateTime:  time.Now().UTC(),
		})

		if err != nil {
			logrus.Errorf("data: %v insert database error: %v", line, err)
		}
	}

	return nil
}

func importThirdCategory() error {
	file, err := os.Open(ThirdCategoryPath)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	for _, line := range lines[1:] {
		code := strings.Split(line[0], ".")

		err = database.InsertPMSEquipmentThirdCategory(&model.PMSEquipmentThirdCategory{
			FirstCode:   line[3],
			SecondCode:  line[2],
			ThirdCode:   code[1],
			Description: line[1],
			Editor:      "system",
			Source:      "system",
			ModifyTime:  time.Now().UTC(),
			CreateTime:  time.Now().UTC(),
		})

		if err != nil {
			logrus.Errorf("data: %v insert database error: %v", line, err)
		}
	}

	return nil
}

func importFourthCategory() error {
	file, err := os.Open(FourthCategoryPath)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	for _, line := range lines[1:] {
		fourthList := strings.Split(line[0], ".")

		if len(fourthList) != 3 {
			continue
		}

		fourth := fourthList[2]
		third := strings.Split(line[0], ".")[1]

		hour, err := strconv.Atoi(line[2])
		if err != nil {
			logrus.Errorf("convert string to int error: %v", err)
			continue
		}

		err = database.InsertPMSEquipmentFourthCategory(&model.PMSEquipmentFourthCategory{
			FirstCode:         line[5],
			SecondCode:        line[6],
			ThirdCode:         third,
			FourthCode:        fourth,
			Description:       line[1],
			DailyRunningHours: hour,
			Location:          "",
			Editor:            "system",
			Source:            "system",
			ModifyTime:        time.Now().UTC(),
			CreateTime:        time.Now().UTC(),
		})

		if err != nil {
			logrus.Errorf("data: %v insert database error: %v", line, err)
		}
	}

	return nil
}
