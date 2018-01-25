// This sample program demonstrates how to use the pool package
// to share a simulated set of database connections.
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elitecodegroovy/gmessage/apps/basic/pool"
)

//Change the value if you want to test with big value.
const (
	maxGoroutines   = 10 // the number of routines to use.
	pooledResources = 5 // number of resources in the pool
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the io.Closer interface so dbConnection
// can be managed by the pool. Close performs any resource
// release management.
func (dbConn *dbConnection) Close() error {
	log.Println("关闭连接 ID：", dbConn.ID)
	return nil
}

// idCounter provides support for giving each connection a unique id.
var idCounter int32

// createConnection is a factory method that will be called by
// the pool when a new connection is needed.
func newConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("创建新连接：", id)

	return &dbConnection{id}, nil
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *pool.Pool) {
	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if conn, ok := conn.(*dbConnection); ok {
		log.Printf("&&&查询: QID[%d] CID[%d]\n", query, conn.ID)
	}else {
		log.Fatalf("断言conn 不是*dbConnection")
	}
}


func startup(){
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// Create the pool to manage our connections.
	p, err := pool.New(newConn, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// Perform queries using connections from the pool.
	for query := 0; query < maxGoroutines; query++ {
		// Each goroutine needs its own copy of the query
		// value else they will all be sharing the same query
		// variable.
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	log.Println("关闭程序.")
	p.Close()
}

// main is the entry point for all Go programs.
func main() {
	startup()
}
