package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math/rand"
)

const (
	cols = 3
	rows 
)

func main(){
	arr := [rows][cols]string{}
	playerState := playerSelectionMode()
	computerState := computerMode(playerState)
	var e int =0
	var p int = 0
	var q int = 0	
	var mark int =0
	var status bool = true
	var isOver bool = false


	fmt.Printf("Player Mark:%s \t KIzo Mark:%s\n",playerState, computerState)
 	board := drawBoard(arr)
	for !isOver {			
		if playerState == "x" || playerState == "X" && computerState == "o"  || computerState == "O" || playerState == "o" || playerState == "O" && computerState == "x" || computerState == "X"{
			fmt.Println("\nPlayer Turn:",mark)
	   		e = moveExecute()
	   		p,q = cood2D(e)
			if board[p][q] == "" {
	   				board[p][q] = playerState
	   				screen(board)
			}else if board[p][q] == "x" || board[p][q] == "X" || board[p][q] == "o" || board[p][q] == "O" {
			fmt.Println("Condition execute this place contain x or o")
				e = moveExecute()
				fmt.Println("\n Player re-enter your position:",mark)
				p,q = cood2D(e)
				if board[p][q] != "" {
	   				board[p][q] = playerState
	   				screen(board)
					
				}
			}
				
			mark = computerMove()
			fmt.Println("\nKizo thinking:",mark)			
			p,q = cood2D(mark)
			if board [p][q] == ""{
				board[p][q] = computerState
	   				screen(board)
			}else if board[p][q] == "x" || board[p][q] == "X" || board[p][q] == "o" || board[p][q] == "O" {
			fmt.Println("Condition execute kizo re-decide his move")
			   mark = computerMove()
			   fmt.Println("\nKizo thinking:",mark)
			   p,q = cood2D(e)
			   if board[p][q] == "" {
	   				board[p][q] = computerState
	   				screen(board)
					
			   }
			  
		        }
			
			
		}
			status = gameOver(board)
			fmt.Println(status)
			stop := stopRoutine(status)
			if stop != false{
				isOver = true
				releaseBy()
				os.Exit(0)
			 }else if status != false{
				isOver = false
			 }
			 
	}
	
}

func playerSelectionMode()string{
	
	fmt.Printf("Press x or o\n")
	fmt.Printf("Press yes : quit , no : cancel\n")	
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "yes"{
			break
		}
		if text == "x" || text =="X"{
			fmt.Printf("****You SELECT Player1**** KIzo ready to beat you\n")
		}else if text == "o" || text == "O"{
			fmt.Printf("****You SELECT Player2**** KIzo ready to beat you\n")
		}
		return text
	}
	
	return "nothing"
}

func computerMode(action string)string{
	
		
		if action == "x" || action =="X"{
			return "o"
		}else if action == "o" || action == "O"{
			return "x"
		}
	
	return "nothing"
}

func drawBoard(arr [cols][rows]string) *[cols][rows]string{
	for i := 0; i < cols; i++{
		for j := 0; j < rows; j++{
			arr[i][j] = ""
		}
	}
	return &arr
}

func moveExecute()int{
        fmt.Printf("Keys Range [0-8]\n")
	move := bufio.NewScanner(os.Stdin)
	for move.Scan(){
		key := move.Text()
		i,err := strconv.Atoi(key)
		 if err == nil && i >= 0 && i< 9{
			return i
		 }
	}
	return -1
}
		
func cood2D(point int) (int , int){

	var p int = -1
	var q int = -1	
	switch point {		
		case 0:
			p,q = 0,0
		break
		case 1:
			p,q = 0,1
		break
		case 2:
			p,q = 0,2
		break
		
		case 3:
			p,q = 1,0
		break
		case 4:
			p,q = 1,1
		break
		case 5:
			p,q = 1,2
		break


		case 6:
			p,q = 2,0
		break
		case 7: 
			p,q = 2,1
		break
		case 8:
			p,q = 2,2
		break
		default:
			fmt.Println("invalid key")
	}
	return p,q
}

func screen(board *[cols][rows]string){
		for i := 0; i < cols; i++{
			for j := 0; j < rows; j++{
				fmt.Printf("%s",board[i][j])
			}
			fmt.Println()
		}
}

func computerMove()int{
	return rand.Intn(8)+rand.Intn(1)
}
				
func gameOver(board *[cols][rows]string)bool{
	var stats bool =false
	if board[0][0] == "x" && board[0][1] == "x" && board[0][2] == "x"{
         stats =true
 }  else if board[0][0] == "X" && board[0][1]== "X" && board[0][2]== "X"{
         stats =true
 }  else if board[0][0] == "x" && board[1][0]== "x" && board[2][0]== "x"{
         stats =true
 }  else if board[0][0] == "X" && board[1][0]== "X" && board[2][0]== "X"{
         stats =true
 }  else if board[0][0] == "x" && board[1][1]== "x" && board[2][2]== "x"{
         stats =true
 }  else if board[0][0] == "X" && board[1][1]== "X" && board[2][2]== "X"{
         stats =true
 }  else if board[0][2] == "X" && board[1][2]== "X" && board[2][2]== "X"{
         stats =true
 }  else if board[0][2] == "x" && board[1][2]== "x" && board[2][2]== "x"{
         stats =true
 }  else if board[0][2] == "x" && board[1][1]== "x" && board[2][0]== "x"{
         stats =true
 }  else if board[0][2] == "X" && board[1][1]== "X" && board[2][0]== "X"{
         stats =true
 }  else if board[2][0] == "x" && board[2][1]== "x" && board[2][2]== "x"{
         stats =true
 }  else if board[2][0] == "X" && board[2][1]== "X" && board[2][2]== "X"{
         stats =true
 }  else if board[1][0] == "X" && board[1][1]== "X" && board[1][2]== "X"{
         stats =true
 } else if board[1][0] == "x" && board[1][1]== "x" && board[1][2]== "X"{
         stats =true
 } else if board[0][0] == "o" && board[0][1] == "o" && board[0][2] == "o"{
         stats =false
 }  else if board[0][0] == "O" && board[0][1]== "O" && board[0][2]== "O"{
        stats =false
 }  else if board[0][0] == "o" && board[1][0]== "o" && board[2][0]== "o"{
         stats =false
 }  else if board[0][0] == "O" && board[1][0]== "O" && board[2][0]== "O"{
        stats =false
 }  else if board[0][0] == "o" && board[1][1]== "o" && board[2][2]== "o"{
         stats =false
 }  else if board[0][0] == "O" && board[1][1]== "O" && board[2][2]== "O"{
         stats =false
 }  else if board[0][2] == "O" && board[1][2]== "O" && board[2][2]== "O"{
         stats =false
 }  else if board[0][2] == "o" && board[1][2]== "o" && board[2][2]== "o"{
         stats =false
 }  else if board[0][2] == "o" && board[1][1]== "o" && board[2][0]== "o"{
        stats =false
 }  else if board[0][2] == "O" && board[1][1]== "O" && board[2][0]== "O"{
        stats =false
 }  else if board[2][0] == "o" && board[2][1]== "o" && board[2][2]== "o"{
        stats =false
 }  else if board[2][0] == "O" && board[2][1]== "O" && board[2][2]== "O"{
        stats =false
 }  else if board[1][0] == "O" && board[1][1]== "O" && board[1][2]== "O"{
       	stats =false
 } else if board[1][0] == "o" && board[1][1]== "o" && board[1][2]== "o"{
         stats =false
 }
	return  stats			 			
}


func stopRoutine(s bool) bool{
	
	if s == true { 
		fmt.Println("\t ******Player Ko****")
		return s
	}
	return false
}	

func releaseBy(){
	fmt.Printf("Ali Hassan")
	fmt.Println("\t Alideveloper95@gmail.com")			 
}
