package fare

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"math"
	"strconv"
)

var (
	base      = 18
	incr      = 5
	fareCache = map[int]int{}
)

var CmdCalc = &cobra.Command{
	Use:   "calc",
	Short: "calc fare",
	Long:  "calc fare",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("need uid")
		}
		uidStr := args[0]
		uid, err := strconv.Atoi(uidStr)
		if err != nil || uid > 10000 {
			log.Fatal("invalid uid")
		}
		cost := calcUserCost(uid)
		fmt.Printf("total: %d\n", cost)
	},
}

// 用户总运费计算
func calcUserCost(uid int) int {
	db := getDb()
	var orders []*Order
	err := db.Where("uid=?", uid).Find(&orders).Error
	if err != nil {
		log.Fatal("query order failed: " + err.Error())
	}
	cost := 0
	for _, order := range orders {
		cost += calc(order.Weight)
	}
	return cost
}

// 运费计算
func calc(weight float64) int {
	w := int(math.Ceil(weight))
	if w <= 1 {
		return base
	}
	if fare, ok := fareCache[w]; ok {
		return fare
	}
	fare := int(math.Round(float64(calc(weight-1))*0.01 + float64(base+incr*(w-1))))
	fareCache[w] = fare
	return fare
}
