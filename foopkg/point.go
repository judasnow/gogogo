package foopkg

// 导出的内容需要精心设计
// 和 api 接口一样 一旦暴露就容易再修改

type Pointfoo struct {
	x int
	y int
	// 这里的名字也要是大写的才能暴露到包外
	Name string
}

func (p Pointfoo) distance(q Pointfoo) int {
	deltaX := p.x - q.x
	deltaY := p.y - q.y
	return (deltaX * deltaX) - (deltaY * deltaY)
}
