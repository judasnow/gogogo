package foopkg

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
