package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// If you remember, we discuss that task should be done in the most possible minimalistic way
// So here is no code architecture, code cleanness, but just base calculating algorithm
// To learn more about my architecture skills welcome to my github https://github.com/entertainment-veks

type Vertex struct {
	ID      string
	Address string
}

type CycleCount struct {
	Count int
	mutex sync.Mutex
}

type Graph struct {
	transactions map[string][]Vertex
	cycleCount   CycleCount
}

type Server struct {
	graph Graph
}

func NewServer() *Server {
	return &Server{
		graph: Graph{
			transactions: make(map[string][]Vertex),
			cycleCount:   CycleCount{},
		},
	}
}

func (s *Server) CountCycles(w http.ResponseWriter, r *http.Request) {
	fromBlock, err := strconv.Atoi(r.FormValue("fromBlock"))
	if err != nil {
		http.Error(w, "Invalid fromBlock value", http.StatusBadRequest)
		return
	}

	toBlock, err := strconv.Atoi(r.FormValue("toBlock"))
	if err != nil {
		http.Error(w, "Invalid toBlock value", http.StatusBadRequest)
		return
	}

	maxCycleLength, err := strconv.Atoi(r.FormValue("maxCycleLength"))
	if err != nil {
		http.Error(w, "Invalid maxCycleLength value", http.StatusBadRequest)
		return
	}

	cycleCount := s.graph.countCyclesInRange(fromBlock, toBlock, maxCycleLength)

	fmt.Fprintf(w, "NumberOfCycles: %d", cycleCount)
}

func (s *Server) MineBlock(w http.ResponseWriter, r *http.Request) {
	transactions := strings.Split(r.FormValue("transactions"), ",")

	s.graph.addTransactions(transactions)

	fmt.Fprint(w, "-1")
}

func (g *Graph) countCyclesInRange(fromBlock, toBlock, maxCycleLength int) int {
	g.cycleCount.mutex.Lock()
	defer g.cycleCount.mutex.Unlock()

	g.cycleCount.Count = 0

	for block := fromBlock; block <= toBlock; block++ {
		transactions, found := g.transactions[strconv.Itoa(block)]
		if !found {
			continue
		}

		visited := make(map[string]bool)

		for _, transaction := range transactions {
			g.dfs(transaction.ID, transaction.ID, transaction.Address, 0, maxCycleLength, visited)
		}
	}

	return g.cycleCount.Count
}

func (g *Graph) dfs(startID, currentID, address string, depth, maxCycleLength int, visited map[string]bool) {
	if depth > maxCycleLength {
		return
	}

	if depth > 0 && currentID == startID && g.transactions[currentID][0].Address == address {
		g.cycleCount.Count++
	}

	visited[currentID] = true

	for _, nextTransaction := range g.transactions[currentID] {
		nextID := nextTransaction.ID

		if !visited[nextID] && nextTransaction.Address == address {
			g.dfs(startID, nextID, address, depth+1, maxCycleLength, visited)
		}
	}

	visited[currentID] = false
}

func (g *Graph) addTransactions(transactions []string) {
	blockID := strconv.Itoa(len(g.transactions))

	for _, transaction := range transactions {
		inputOutput := strings.Split(transaction, ":")
		if len(inputOutput) != 2 {
			log.Printf("Invalid transaction format: %s", transaction)
			continue
		}

		txID := inputOutput[0]
		outputIndex, err := strconv.Atoi(inputOutput[1])
		if err != nil {
			log.Printf("Invalid transaction output index: %s", inputOutput[1])
			continue
		}

		address := g.transactions[txID][outputIndex].Address

		vertex := Vertex{
			ID:      fmt.Sprintf("%s:%d", blockID, len(g.transactions[blockID])),
			Address: address,
		}

		g.transactions[blockID] = append(g.transactions[blockID], vertex)
	}
}

// You can run this server, and it will listen on port 8080 for incoming HTTP requests.
// You can send the queries using HTTP GET or POST methods to the appropriate endpoints.
//
// For example, to execute the CountCycles query, you can make a GET request
// to http://localhost:8080/count_cycles?fromBlock=0&toBlock=2&maxCycleLength=1.
// The response will contain the number of cycles.
// To execute the MineBlock query, you can make a POST request to http://localhost:8080/mine_block
// with the transaction data in the request body as a comma-separated string, e.g.,
// "transactions_from_block_2,transactions_from_block_3".

func main() {
	server := NewServer()

	http.HandleFunc("/count_cycles", server.CountCycles)
	http.HandleFunc("/mine_block", server.MineBlock)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
