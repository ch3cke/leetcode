package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	if word == "" {
		return true
	}
	used := [][]bool{}
	for i := 0; i < m; i++ {
		used = append(used, make([]bool, n))
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == word[0] {
				used[r][c] = true
				if dfs(board, used, r, c, word[1:]) {
					return true
				}
				used[r][c] = false
			}
		}
	}
	return false
}

func dfs(board [][]byte, used [][]bool, row, col int, word string) bool {
	if word == "" {
		return true
	}
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var r, c int
	for _, d := range ds {
		r, c = row+d[0], col+d[1]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !used[r][c] && board[r][c] == word[0] {
			used[r][c] = true
			if dfs(board, used, r, c, word[1:]) {
				return true
			}
			used[r][c] = false
		}
	}
	return false
}

//func exist(board [][]byte, word string) bool {
//	begin_1,begin_2 := 0,0
//	for{
//		begin_1, begin_2 = getIndex(board, byte(word[0]), begin_1, begin_2)
//		if begin_1==-1{
//			return false
//		}else {
//			a,b := begin_1,begin_2
//			from := []int{-2,-2}
//			re :=[][]int{}
//			for i:=1;i<len(word);i++{
//				a,b, from = getNextPlace(board,word[i],a,b,from)
//				if a == -1{
//					if len(re)!=0{
//						a=re[len(re)-1][0]
//						b=re[len(re)-1][1]
//					}else {
//						break
//					}
//				}else {
//					re = append(re, []int{a,b})
//				}
//			}
//			begin_2 = begin_2 +1
//			if a!= -1{
//				return true
//			}
//		}
//	}
//}
//
//func getIndex(board [][]byte, word byte, begin_1 int, begin_2 int)(int, int){
//	rows,colunm := len(board),len(board[0])
//	for i:=begin_1;i<rows;i++{
//		if i!=begin_1{
//			begin_2=0
//		}
//		for j:=begin_2;j<colunm;j++{
//			if board[i][j]==word{
//				return i,j
//			}
//		}
//	}
//	return -1,-1
//}
//
//func getNextPlace(board [][]byte, word byte, row, cloun int, from []int)(int, int, []int){
//	a := row+1
//	b := row-1
//	c := cloun+1
//	d := cloun-1
//	tmp:=[]int{-1,-1}
//	if a<len(board){
//		if board[a][cloun] == word{
//			tmp = []int{1,0}
//			if !reflect.DeepEqual(from,[]int{-1,0}){
//				return a,cloun,tmp
//			}
//		}
//	}
//	if c <len(board[0]){
//		if board[row][c]== word{
//			tmp = []int{0,1}
//			if !reflect.DeepEqual(from,[]int{0,-1}){
//				return row,c,tmp
//			}
//		}
//	}
//	if d >=0 {
//		if board[row][d]==word{
//			tmp = []int{0,-1}
//			if !reflect.DeepEqual(from,[]int{0,1}){
//				return row,d,tmp
//			}
//		}
//	}
//	if b>=0{
//		if board[b][cloun]==word{
//			tmp = []int{-1,0}
//			if !reflect.DeepEqual(from,[]int{1,0}){
//				return b,cloun, tmp
//			}
//		}
//	}
//	return -1,-1,tmp
//}
