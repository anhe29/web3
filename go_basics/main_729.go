package main

type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(-1, -1) // 哨兵
	return MyCalendar{t}
}

func (c *MyCalendar) Book(start, end int) bool {
	floor, _ := c.Floor(start)
	if floor.Value.(int) > start { // [start,end) 左侧区间的右端点超过了 start
		return false
	}
	if it := c.IteratorAt(floor); it.Next() && it.Key().(int) < end { // [start,end) 右侧区间的左端点小于 end
		return false
	}
	c.Put(start, end) // 可以插入区间 [start,end)
	return true
}
