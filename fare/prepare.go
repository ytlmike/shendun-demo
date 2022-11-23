package fare

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

var CmdInit = &cobra.Command{
	Use:   "init",
	Short: "init data",
	Long:  "init data",
	Run: func(cmd *cobra.Command, args []string) {
		PrepareData()
		fmt.Println("Done.")
	},
}

type Order struct {
	Id        int
	Uid       int
	Weight    float64
	CreatedAt time.Time
}

// PrepareData 生成数据
func PrepareData() {
	db := getDb()
	rand.Seed(time.Now().Unix())
	db.AutoMigrate(&Order{})
	chooser := Chooser{}
	for w := 1; w <= 100; w++ {
		chooser.AddChoice(w, float64(1)/(float64(w)))
	}
	batchSize := 1000
	var batch []Order
	for id := 1; id <= 100000; id++ {
		batch = append(batch, Order{
			Id:        id,
			Uid:       rand.Intn(10000) + 1,
			Weight:    float64(chooser.Pick()),
			CreatedAt: time.Now(),
		})
		if len(batch) >= batchSize {
			err := db.Create(batch).Error
			if err != nil {
				log.Fatal("save orders failed: " + err.Error())
			}
			batch = []Order{}
		}
	}
	return
}

// ------------- 加权随机选择 -------------

type Choice struct {
	Item   int
	Weight float64
}

type Chooser struct {
	choices []Choice
}

func (c *Chooser) AddChoice(item int, weight float64) {
	c.choices = append(c.choices, Choice{
		Item:   item,
		Weight: weight,
	})
}

func (c *Chooser) Pick() int {
	if len(c.choices) == 1 {
		return 0
	}
	var sum, t float64
	for _, w := range c.choices {
		sum += w.Weight
	}
	r := rand.Float64() * sum
	for _, w := range c.choices {
		t += w.Weight
		if t > r {
			return w.Item
		}
	}
	return len(c.choices) - 1
}
