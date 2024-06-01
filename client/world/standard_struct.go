package world

type Industry struct {
	Id         int
	Name       string
	GoodsId    string
	Size       int
	Employee   int
	EmployPay  int
	GoodsPrice int
	Cash       int
}

type Country struct {
	Id        int
	IsPlayer  bool
	Human     [25]float64 // 75 岁寿命、每 3 年一层
	Worker    [4]int64
	Industry  []Industry
	birthRate float64
}

type GoodsType int8

var (
	GoodsTypeFood     = GoodsType(1)
	GoodsTypeClothing = GoodsType(2)
)

type Goods struct {
	Id              int
	Type            GoodsType
	ConsumePreHuman int
}

type World struct {
	Countries []*Country
	Goods     []*Goods
}

func (w *World) Init() {
	w.Countries = []*Country{
		{Id: 0, Human: [25]float64{}, Worker: [4]int64{}, IsPlayer: true, Industry: make([]Industry, 3)},
	}
}
