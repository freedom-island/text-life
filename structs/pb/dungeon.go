package pb

type DungeonEnv int

const (
	AbandonHome        DungeonEnv = 1 // 废弃住所
	AbandonResearchLab DungeonEnv = 2 // 废弃研究所
)

type DungeonConfig struct {
	LevelMonster int // 难度
	LevelGoods   int // 设施物资的稀有度
	RoomCount    int
	Deep         int // 最终深度
	BossIndex    int // 0 表示没有 Boss
	Environment  DungeonEnv
}

func (d *NowDungeon) Init(playerId uint, config DungeonConfig) {
	d.PlayerId = uint64(playerId)
	//d.RoomFirst = new(Room)

	d.CreateRooms(config)
	d.CreateCorridor(config)
	d.CreateBox(config)
}

// CreateRooms 创建房间
// 房子位置为 Y 轴 + 1
func (d *NowDungeon) CreateRooms(config DungeonConfig) {
	for once, _ := range make([]struct{}, config.RoomCount) {
		d.Rooms = append(d.Rooms, &Room{
			Id:        1,
			Pos:       &Vec2D{X: 0, Y: uint64(once)},
			Title:     "入口",
			Box:       []*Box{},
			LinkRooms: []uint64{},
			Monsters:  []*Monster{},
		})
	}
}

// CreateCorridor 创建走廊
func (d *NowDungeon) CreateCorridor(config DungeonConfig) {

}

// CreateBox 创建容器
func (d *NowDungeon) CreateBox(config DungeonConfig) {

}
