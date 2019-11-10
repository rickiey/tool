package tool

import (
	"container/list"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	idsGenerator *idpool
	initOnce     sync.Once
)

type idpool struct {
	ids           *list.List
	lock          sync.Mutex
	snowFlakeNode *snowflake.Node
	/*
		The ID as a whole is a 63 bit integer stored in an int64
		41 bits are used to store a timestamp with millisecond precision, using a custom epoch.
		10 bits are used to store a node id - a range from 0 through 1023.
		12 bits are used to store a sequence number - a range from 0 through 4095.
	*/
}

const (
	// IDpool 大小
	IDpoolCount = 1000
	// 添加门槛
	AddThreshold = 300
)

func init() {
	initOnce.Do(func() {

		idsGenerator = new(idpool)
		idsGenerator.ids = list.New()

		node, err := snowflake.NewNode(201)
		if err != nil {
			panic(err.Error())
		}
		idsGenerator.snowFlakeNode = node
		addID()
	})

}

func GetID() snowflake.ID {

	if idsGenerator.ids.Len() < AddThreshold {
		addID()
	}
	id := idsGenerator.ids.Front()
	idsGenerator.ids.Remove(id)
	return id.Value.(snowflake.ID)
}

// 添加 ID
func addID() {
	idsGenerator.lock.Lock()
	defer idsGenerator.lock.Unlock()
	for i := 0; i < IDpoolCount; i++ {
		if idsGenerator.ids.Len() > IDpoolCount {
			return
		}
		idsGenerator.ids.PushBack(idsGenerator.snowFlakeNode.Generate())
	}
}

/////////////////////////////////////////////////////
//
//var (
//	SnowFlakeID *sonyflake.Sonyflake
//	initOnce    sync.Once
//	IDpool      *list.List
//)
//
//func init() {
//	//err := errors.New("already initialized")
//	IDpool = list.New()
//	initOnce.Do(func() {
//		SnowFlakeID = InitSnowFlakeIDgenerator(16)
//	})
//	for i := 0; i < 5; i++ {
//		AddID()
//	}
//}
//
//func InitSnowFlakeIDgenerator(machineID uint16) *sonyflake.Sonyflake {
//
//	t := time.Now()
//	t.Add(- (time.Minute * 3))
//	st := sonyflake.Settings{
//		StartTime:      t,
//		MachineID:      func() (uint16, error) { return machineID, nil },
//		CheckMachineID: func(u uint16) bool { return true },
//	}
//	sf := sonyflake.NewSonyflake(st)
//	if sf == nil {
//		panic("Snowflake algorithm ID generator error")
//		return nil
//	}
//	return sf
//}
//
//func GetID() uint64 {
//	if IDpool.Len() > 0 {
//		id :=  IDpool.Front()
//		IDpool.Remove(id)
//		if IDpool.Len() < 600 {
//			go AddID()
//		}
//		v := id.Value.(uint64)
//		return v
//	}
//	if SnowFlakeID == nil {
//		panic("SnowFlakeID is nil")
//		return 0
//	}
//	id, err := SnowFlakeID.NextID()
//	if err != nil {
//		panic(err.Error())
//		return 0
//	}
//	return id
//}
//
//
//// add 200 id
//func AddID() {
//	if SnowFlakeID == nil {
//		panic("SnowFlakeID is nil")
//	}
//	for i:=0; i< 200; i++ {
//		id, err := SnowFlakeID.NextID()
//		if err != nil {
//			panic(err.Error())
//		}
//		IDpool.PushBack(id)
//	}
//}
