package alghoritms

import (
	"math/rand"
	"time"
)

type Algorithm interface {
	Finished() bool
	Init()
	Step()
}

type BFS struct {
	Space *SearchSpace

	visited  map[Point]bool
	frontier []Point
}

func NewBFS(space *SearchSpace) *BFS {
	return &BFS{space, make(map[Point]bool), make([]Point, 0)}
}

func (alg *BFS) Finished() bool {
	x, y := alg.Space.Board.Labyrinth.Dimensions()
	return alg.visited[Point{x - 1, y - 1}]
}

func (alg *BFS) Init() {
	alg.visited = make(map[Point]bool)
	alg.frontier = []Point{{0, 0}}
	alg.Space.Visit(0, 0)

	rand.Seed(time.Now().Unix())
}

func (alg *BFS) Step() {
	alg.Space.Update()
	alg.Space.MarkFrontier(alg.frontier...)

	curr := alg.frontier[0]
	alg.Space.Visit(curr.X, curr.Y)
	alg.frontier = alg.frontier[1:]

	neigh := alg.getNeighbours(curr)
	alg.frontier = append(alg.frontier, neigh...)
	for _, n := range neigh {
		alg.visited[n] = true
		alg.Space.Visit(n.X, n.Y)
	}

	if alg.Finished() {
		for p := range alg.visited {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.8, 0.0)
		}
		for _, p := range alg.frontier {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.4, 0.0)
		}
	}
}

func (alg *BFS) getNeighbours(curr Point) []Point {
	points := make([]Point, 0)
	maxX, maxY := alg.Space.Board.Labyrinth.Dimensions()
	maxX -= 1
	maxY -= 1
	for _, val := range rand.Perm(4) {
		if val == 0 && curr.Y != maxY {
			if !alg.visited[Point{curr.X, curr.Y + 1}] && !alg.Space.GetField(curr.X, curr.Y+1) {
				points = append(points, Point{curr.X, curr.Y + 1})
			}
		}
		if val == 1 && curr.X != maxX {
			if !alg.visited[Point{curr.X + 1, curr.Y}] && !alg.Space.GetField(curr.X+1, curr.Y) {
				points = append(points, Point{curr.X + 1, curr.Y})
			}
		}
		if val == 2 && curr.X != 0 {
			if !alg.visited[Point{curr.X - 1, curr.Y}] && !alg.Space.GetField(curr.X-1, curr.Y) {
				points = append(points, Point{curr.X - 1, curr.Y})
			}
		}
		if val == 3 && curr.Y != 0 {
			if !alg.visited[Point{curr.X, curr.Y - 1}] && !alg.Space.GetField(curr.X, curr.Y-1) {
				points = append(points, Point{curr.X, curr.Y - 1})
			}
		}
	}
	return points
}

type DFS struct {
	*BFS
}

func NewDFS(space *SearchSpace) *DFS {
	return &DFS{&BFS{space, make(map[Point]bool), make([]Point, 0)}}
}

func (alg *DFS) Step() {
	alg.Space.Update()
	alg.Space.MarkFrontier(alg.frontier...)

	curr := alg.frontier[0]
	alg.Space.Visit(curr.X, curr.Y)
	neighbours := alg.getNeighbours(curr)

	alg.frontier = alg.frontier[1:]
	alg.frontier = append(neighbours, alg.frontier...)

	for _, n := range neighbours {
		alg.Space.Visit(n.X, n.Y)
		alg.visited[n] = true
	}

	if alg.Finished() {
		for p := range alg.visited {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.8, 0.0)
		}
		for _, p := range alg.frontier {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.4, 0.0)
		}
	}
}

type AStar struct {
	*BFS
}

func NewAStar(space *SearchSpace) *AStar {
	return &AStar{&BFS{space, make(map[Point]bool), make([]Point, 0)}}
}


//todo It's not an AStar
func (alg *AStar) Step() {
	alg.Space.Update()
	alg.Space.MarkFrontier(alg.frontier...)

	curr := alg.frontier[0]
	alg.Space.Visit(curr.X, curr.Y)
	alg.frontier = alg.frontier[1:]

	neighbours := alg.getNeighbours(curr)

	for _, n := range neighbours {
		newFrontier := make([]Point, 0)
		for i, el := range alg.frontier {
			if alg.rank(n) <= alg.rank(el) {
				if i == 0 {
					newFrontier = append([]Point{n}, alg.frontier...)
				} else {
					newFrontier = append(alg.frontier[:i], n)
					newFrontier = append(newFrontier, alg.frontier[i:]...)
				}
				alg.visited[n] = true
				alg.Space.Visit(n.X, n.Y)
				break
			}
		}
		if !alg.visited[n] {
			alg.visited[n] = true
			alg.Space.Visit(n.X, n.Y)
			newFrontier = append(alg.frontier, n)
		}
		alg.frontier = newFrontier
	}

	if alg.Finished() {
		for p := range alg.visited {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.8, 0.0)
		}
		for _, p := range alg.frontier {
			alg.Space.SetColor(p.X, p.Y, 0.0, 0.4, 0.0)
		}
	}

}

func (alg *AStar) rank(p Point) uint {
	x, y := alg.Space.Dimensions()
	return (x - p.X) + (y - p.Y)
}
