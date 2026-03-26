package hashtable

import "fmt"

type NodeBfs struct {
	row, col, steps int
}

var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}

func BFS2D(grid []string, start, end NodeBfs) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := []NodeBfs{{start.row, start.col, 0}}

	fmt.Printf("🚀 Bắt đầu: (%d,%d)\n", start.row, start.col)
	fmt.Printf("🎯 Đích:    (%d,%d)\n", end.row, end.col)
	fmt.Println("─────────────────────────────────────────────")

	for len(queue) > 0 {
		fmt.Print("Queue: [")
		for _, n := range queue {
			fmt.Printf("(%d,%d):%d ", n.row, n.col, n.steps)
		}
		fmt.Println("]")
		curr := queue[0]
		queue = queue[1:]

		fmt.Printf("👉 Đang xét: (%d,%d) bước %d\n", curr.row, curr.col, curr.steps)
		if curr.row == end.row && curr.col == end.col {
			fmt.Println("─────────────────────────────────────────────")
			fmt.Printf("✅ Tìm thấy đích sau %d bước!\n", curr.steps)
			return curr.steps
		}
		for i := 0; i < 4; i++ {
			nr := curr.row + dr[i]
			nc := curr.col + dc[i]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] != '#' && !visited[nr][nc] {
				visited[nr][nc] = true
				queue = append(queue, NodeBfs{nr, nc, curr.steps + 1})
				fmt.Printf("Thêm (%d,%d) bước %d\n", nr, nc, curr.steps+1)
			}
			fmt.Println()
		}
	}
	return -1
}
