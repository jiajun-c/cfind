package dal

import "fmt"

type Item struct {
	Local           string // 所在地区
	Spu_id          string
	Suggested_index int    // 建议序列
	Cover_url       string // 封面url
	Sub_name        string // 分店信息
	ShopID          string // 店铺id，同一店铺不同分店的id也是相同的
}

func (item Item) Output() string {
	return fmt.Sprintf("City = %s ,SpuID = %s, SortLevel = %d,CoverImage = %s", item.Local, item.Spu_id, item.Suggested_index, item.Cover_url)
}

type Doctor struct {
	Local           string
	Suggested_index int // 建议序列
	Doc_id          string
}

func (item Doctor) Output() string {
	return fmt.Sprintf("City = %s ,DocID = %s, SortLevel = %d, IsFemtosecond = false", item.Local, item.Doc_id, item.Suggested_index)
}

type Organization struct {
	Local           string
	StoreId         string
	Suggested_index int    // 建议序列
	Sups_id         string // 商品信息
}

func (org Organization) Output() string {
	return fmt.Sprintf("City = %s ,StoreID = %s, SortLevel = %d, SpuID = %s", org.Local, org.StoreId, org.Suggested_index, org.Sups_id)

}
