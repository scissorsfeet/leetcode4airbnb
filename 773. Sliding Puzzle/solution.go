package _73__Sliding_Puzzle

type pair struct {
	board [][]int
	idx   []int
}

func slidingPuzzle(board [][]int) int {
	correct := [][]int{{1, 2, 3}, {4, 5, 0}}
	directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	q := NewQueue()
	rowLen := len(board)
	colLen := len(board[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if board[i][j] == 0 {
				q.enQueue(&pair{board, []int{i, j}})
			}
		}
	}

	visited := make(map[string]bool)
	var count int
	for q.size() > 0 {
		size := q.size()
		for i := 0; i < size; i++ {
			p := q.deQueue()
			if isSame(correct, p.board) {
				return count
			}
			key := trans2str(p.board)
			visited[key] = true
			for _, dir := range directions {
				x, y := p.idx[0]+dir[0], p.idx[1]+dir[1]
				if x < 0 || x > rowLen-1 || y < 0 || y > colLen-1 {
					continue
				}
				newBoard := make([][]int, rowLen)
				for i := range newBoard {
					newBoard[i] = make([]int, colLen)
					for j := 0; j < colLen; j++ {
						newBoard[i][j] = p.board[i][j]
					}
				}
				newBoard[p.idx[0]][p.idx[1]] = newBoard[x][y]
				newBoard[x][y] = 0
				key := trans2str(newBoard)
				if v, ok := visited[key]; ok && v {
					continue
				}
				q.enQueue(&pair{newBoard, []int{x, y}})
			}
		}
		count++
	}

	return -1
}

func trans2str(board [][]int) string {
	res := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			res += string(board[i][j] + 'a')
		}
	}
	return res
}

func isSame(a, b [][]int) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return false
	}
	for i := 0; i < aLen; i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

type queue struct {
	data []*pair
}

func NewQueue() *queue {
	return &queue{data: make([]*pair, 0)}
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) enQueue(x *pair) {
	q.data = append(q.data, x)
}

func (q *queue) deQueue() *pair {
	x := q.data[0]
	q.data = q.data[1:]
	return x
}
