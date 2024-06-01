package config

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"sync"
)

var onceConfig sync.Once

type Room struct {
	Id        uint
	Title     string
	InfoIndex []string
}
type Goods struct {
	Id    uint
	Title string
}
type Config struct {
	Params    map[string]string
	Rooms     []*Room
	GoodsList []*Goods
}

// type global
var (
	payload *Config
)

func GetConfig(projectPath string) *Config {
	if payload != nil {
		return payload
	}

	onceConfig.Do(func() {
		payload = new(Config)

		// 生成配置
		envFile, err2 := os.Getwd()
		if err2 != nil {
			return
		}
		if os.Getenv("CODENATION_ENV") == "prod" {
			envFile = envFile + projectPath + "/env-prod.json"
		} else if os.Getenv("CODENATION_ENV") == "test" {
			envFile = envFile + projectPath + "/env-test.json"
		} else {
			envFile = envFile + projectPath + "/env-dev.json"
		}
		content, err := os.ReadFile(envFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(content, &payload.Params)
		if err != nil {
			panic(err)
		}

		err = payload.GetExcel()
		if err != nil {
			panic(err)
		}
	})

	return payload
}

func (c *Config) GetExcel() error {
	file, err := excelize.OpenFile("config.xlsx")
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := file.GetRows("goods")
	if err != nil {
		return err
	}
	header := make([]string, 0)
	for i, row := range rows {
		if i == 0 {
			for _, title := range row {
				header = append(header, title)
			}
		} else {
			for i2, s := range row {
				rowConfig := Goods{}
				switch header[i2] {
				case "id":
					id, err := strconv.Atoi(s)
					if err != nil {
						panic(err)
					}
					rowConfig.Id = uint(id)
				case "title":
					rowConfig.Title = s
				}
			}
		}
	}

	return nil
}
