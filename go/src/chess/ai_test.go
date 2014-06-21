package chess

import (
	"fmt"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := createBoard(initial_board_json)
	fmt.Println(board)
}

func TestCreateChessNode(t *testing.T) {
	node := createChessNode(initial_board_json)
	fmt.Println("node = ", node)
}

func TestPrintNode(t *testing.T) {
	node := createChessNode(initial_board_json)
	printNode(node)
}

func TestParseFen(t *testing.T) {
	//fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	//TODO split on space above, for now just use placement directly
	fenPlacement := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	fmt.Println("board = ", board)
}

//func TestGetNextMoveUsingPointValue(t *testing.T) {
//	var dat []string
//	dat = append(dat, "[[\"wr\",\"wn\",\"wb\",\"wq\",\"wk\",\"wb\",\"wn\",\"wr\"],[\"wp\",\"wp\",\"wp\",\"wp\",0,\"wp\",\"wp\",\"wp\"],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[\"bp\",0,0,0,\"wp\",0,0,0],[0,\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\"],[\"br\",\"bn\",\"bb\",\"bq\",\"bk\",\"bb\",\"bn\",\"br\"]]")
//	fmt.Println("dat=", dat)
//	move := GetNextMoveUsingPointValue(dat)
//	fmt.Println("move=", move)
//}

func TestMakeMove(t *testing.T) {
	fenPlacement := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	node := ChessNode{board: board}
	node.legal_moves = getMoves("b", node.board)

	board2 := makeMove(node.board, node.legal_moves[0])
	fmt.Println("board2=", board2)
}

func TestMiniMax(t *testing.T) {
	fenPlacement := "7k/8/8/7p/6Q1/8/8/K8"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)

	//the white queen moved into the black pawn's attack square
	node := createState(board, "b", "g3-g4")

	move := miniMaxDecision(node)
	fmt.Println("move = ", move)

	//the black pawn should attack this square!
	if move != "\"next-move\":\"h5-g4\"" {
		t.Error("WRONG MOVE!!!")
	}
}

func TestBlackKillQueen(t *testing.T) {
	fenPlacement := "r1bqk2r/1pp1b3/p1PpQn1p/6p1/8/2PBBN1P/PP3PP1/R3K2R"
	board := createBoardUsingFen(fenPlacement)

	//black should kill the white queen
	node := createState(board, "b", "e4-e6")

	move := miniMaxDecision(node)
	fmt.Println("move = ", move)

	//the black pawn should attack this square!
	if move != "\"next-move\":\"c8-e6\"" {
		t.Error("WRONG MOVE!!!")
	}
}
