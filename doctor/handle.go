package doctor

import (
	"choice/dal"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
)

var mark int = 0

func SelectFromItemSheet(file *excelize.File, n int) {
	orthodontics := make([]dal.Doctor, 0)
	implantedTeeth := make([]dal.Doctor, 0)
	implantedHair := make([]dal.Doctor, 0)
	rows, _ := file.GetRows("医生列表")
	var suggestedIndex int
	for _, row := range rows {
		if mark == 0 {
			mark = 1
			continue
		}
		if row[3] != "" {
			suggestedIndex, _ = strconv.Atoi(row[4])
		}
		doctor := dal.Doctor{
			Local:           row[3],
			Suggested_index: suggestedIndex,
			Doc_id:          row[1],
		}
		switch row[0] {
		case "种植牙":
			implantedTeeth = append(implantedTeeth, doctor)
		case "牙齿矫正":
			orthodontics = append(orthodontics, doctor)
		case "植发":
			implantedHair = append(implantedHair, doctor)
		}
	}
	itemTypes := []string{"ImplantedTeeth", "Orthodontics", "ImplantedHair"}
	itemLists := [][]dal.Doctor{implantedTeeth, orthodontics, implantedHair}
	for i := 0; i < 3; i++ {
		citys := selectFromCity(itemLists[i])
		for index, city := range citys {
			sort.Sort(SortList(city))
			citys[index] = city[0:min(len(city), n)] // 最多展示n个
		}
		output(citys, itemTypes[i])
	}
}

func output(lists [][]dal.Doctor, itemType string) {
	fmt.Printf("%s = [\n", itemType)
	for _, list := range lists {
		for _, item := range list {
			fmt.Printf("	{%s},\n", item.Output())
		}
		fmt.Println("")
	}
	fmt.Println("]")
}

type SortList []dal.Doctor

func (s SortList) Len() int {
	return len(s)
}

func (s SortList) Less(i, j int) bool {
	if s[i].Local != s[j].Local {
		return s[i].Local > s[j].Local
	}
	return s[i].Suggested_index < s[j].Suggested_index
}

func (s SortList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func selectFromCity(lists []dal.Doctor) [][]dal.Doctor {
	sort.Sort(SortList(lists))
	res := make([][]dal.Doctor, 0)
	i := 0
	var j int

	for {
		if i > len(lists) {
			return res
		}
		for j = i + 1; j < len(lists) && lists[j].Local == lists[i].Local; j++ {

		}
		if j > len(lists) {
			return res
		}
		for _, list := range lists[i:j] {
			if list.Local == "" {
				fmt.Println("err")
			}
		}
		res = append(res, lists[i:j])
		i = j
	}
}
