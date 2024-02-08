package leetcode
//not working 
func orangesRotting (grid [][]int ) int {
	var rottenList[][2]int;
	freshMap := make(map[[2]int]bool);
	for i:=0;i<len(grid); i++ {
		for j:=0;j<len(grid[i]); j++ {
			if grid[i][j] == 2 {
				var temp [2]int;
				temp[0] = i;
				temp[1] = j;
				rottenList = append(rottenList, temp);
			}else if grid[i][j] == 1 {
				var temp [2]int;
				temp[0] = i;
				temp[1] = j;
				freshMap[temp] = true;
			}
		}
	}
	var flag bool = true;
	var count int = 0;
	for flag {
		for i:=0; i< len(rottenList) ; i++ {
			curr := rottenList[i];
			up := getUp(curr);
			down := getDown(curr);
			left := getLeft(curr);
			right := getRight(curr);
			flag = false;
			if freshMap[up] {
				grid[up[0]][up[1]] = 2;
				flag = true;
			}
			if freshMap[down] {
				grid[down[0]][down[1]] = 2;
				flag = true;
			}
			if freshMap[left] {
				grid[left[0]][left[1]] = 2;
				flag = true;
			}
			if freshMap[right] {
				grid[right[0]][right[1]] = 2;
				flag = true;
			}
			if flag {
				count++;
			}
		}
	}
	return 0;
}


func getUp(cords [2]int) [2]int{
	//if cords[0] - 1 >= 0 {
	cords[0] = cords[0] - 1 ;
	return cords;
	//}else {
	//	temp := new([2]int);
	//	return *temp;
	//}
}

func getDown(cords [2]int) [2]int{
	cords[0] = cords[0] + 1 ;
	return cords;
}

func getRight(cords [2]int) [2]int{
	cords[1] = cords[1] + 1 ;
	return cords;
}

func getLeft(cords [2]int) [2]int{
	cords[1] = cords[1]-1 ;
	return cords;
}
