package ttime

import (
	"time"

	"zqx.com/goutil/base"
)

// func main() {
// 	//一  获取当前时间
// 	// 返回当前时间，注意此时返回的是 time.Time 类型
// 	now := time.Now()
// 	fmt.Println(now)
// 	// 当前时间戳
// 	fmt.Println(now.Unix())
// 	// 纳秒级时间戳
// 	fmt.Println(now.UnixNano())
// 	// 时间戳小数部分 单位：纳秒
// 	fmt.Println(now.Nanosecond())
// 	//二   格式化
// 	fmt.Println(now.Format("2006-01-02 15:03:04"))
// 	fmt.Println(now.Format("2006-01-02"))
// 	fmt.Println(now.Format("15:03:04"))
// 	fmt.Println(now.Format("2006/01/02 15:04"))
// 	fmt.Println(now.Format("15:04 2006/01/02"))
// 	// 三
// 	layout := "2006-01-02 15:04:05"
// 	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
// 	fmt.Println(t.Format(layout))
// 	//参考 https://zhuanlan.zhihu.com/p/362936088

// 	GetPeriod(0, 0)
// }

func GetPeriod(start int64, end int64) []string {
	allDateArray := make([]string, 0)
	// start, end = 1631584515, 1631584516
	timeFormat := "2006-01-02"
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	//After方法 a.After(b) a,b Time类型 如果a时间在b时间之后，则返回true
	for endTime.After(startTime) {
		allDateArray = append(allDateArray, startTime.Format(timeFormat))
		startTime = startTime.AddDate(0, 0, 1)
	}
	allDateArray = append(allDateArray, endTime.Format(timeFormat))
	allDateArray = base.RemoveDuplicateElement(allDateArray)
	//fmt.Println(allDateArray)
	return allDateArray
}
