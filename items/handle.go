package items

import (
	"choice/dal"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
)

var mark int = 0

func SelectFromItemSheet(file *excelize.File, n int, k int) {
	orthodontics := make([]dal.Item, 0)
	implantedTeeth := make([]dal.Item, 0)
	implantedHair := make([]dal.Item, 0)
	rows, _ := file.GetRows("商品列表")
	var suggestedIndex int
	for _, row := range rows {
		if mark == 0 {
			mark = 1
			continue
		}
		if row[3] != "" {
			suggestedIndex, _ = strconv.Atoi(row[3])
		} else {
		}

		item := dal.Item{
			Local:           row[1],
			Spu_id:          row[2],
			Suggested_index: suggestedIndex,
			Cover_url:       row[6],
			Sub_name:        row[8],
			ShopID:          row[9],
		}
		switch row[0] {
		case "种植牙":
			implantedTeeth = append(implantedTeeth, item)
		case "牙齿矫正":
			orthodontics = append(orthodontics, item)
		case "植发":
			implantedHair = append(implantedHair, item)
		}
	}
	itemTypes := []string{"ImplantedTeeth", "Orthodontics", "ImplantedHair"}
	itemLists := [][]dal.Item{implantedTeeth, orthodontics, implantedHair}
	for i := 0; i < 3; i++ {
		citys := selectFromCity(itemLists[i], k)
		for index, city := range citys {
			sort.Sort(SortList(city))
			citys[index] = city[0:min(len(city), n)] // 最多展示n个
		}
		output(citys, itemTypes[i])
	}
}

func output(lists [][]dal.Item, itemType string) {
	fmt.Printf("%s = [\n", itemType)
	for _, list := range lists {
		for _, item := range list {
			fmt.Printf("	{%s},\n", item.Output())
		}
		fmt.Println("")
	}
	fmt.Println("]")
}

type SortList []dal.Item

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

// 从城市中进行筛选，允许一个店铺有两个展示的
func selectFromCity(lists []dal.Item, k int) [][]dal.Item {
	sort.Sort(SortList(lists))
	res := make([][]dal.Item, 0)
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
		temp := filterSameShop(lists[i:j], k)
		for index, _ := range temp {
			temp[index].Suggested_index = index + 1
		}
		res = append(res, temp)

		i = j
	}
}

func filterSameShop(lists []dal.Item, k int) []dal.Item {
	ans := make([]dal.Item, 0)
	i := 0
	var j int
	for {
		if i > len(lists) {
			return ans
		}
		for j = i + 1; j < len(lists) && lists[j].ShopID == lists[i].ShopID; j++ {
		}
		if j > len(lists) {
			return ans
		}
		ans = append(ans, lists[i:min(j, i+k)]...)
		i = j
	}
}
