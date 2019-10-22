package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

    There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
    The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
    In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    The host allows no more than 2 philosophers to eat concurrently.
    Each philosopher is numbered, 1 through 5.
    When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
    When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.


*/
const (
	EAT_MIN_MS   = 1e9
	EAT_MAX_MS   = 2e9
	THINK_MIN_MS = 2e9
	THINK_MAX_MS = 5e9
	CHOPS_QTY    = 5
)

type Philosopher struct {
	Id          int
	Left_index  int
	Right_index int
	counter     int
}

type Monitor interface {
	AllowDiner(p *Philosopher)
}

type Host struct {
	name string
}

var ChopsMutex sync.Cond
var Chopsticks [CHOPS_QTY]chan bool
var rgen *rand.Rand
var randMutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	var p [CHOPS_QTY]Philosopher
	initChops()
	host := Host{"Jenkins"}

	wg.Add(CHOPS_QTY)
	for i := 1; i < CHOPS_QTY; i++ {
		p[i] = Philosopher{i + 1, i, (i + 1) % CHOPS_QTY, 0}
		go Run(&p[i], &host, &wg)
	}
	p[0] = Philosopher{1, 0, (0 + 1) % CHOPS_QTY, 0}
	Run(&p[0], &host, &wg)

	wg.Wait()
	for _, philo := range p {
		fmt.Printf("philosopher %d ate %d times\n", philo.Id, philo.counter)
	}
}

func initChops() {
	ChopsMutex = sync.Cond{L: &sync.Mutex{}}
	for i := 0; i < CHOPS_QTY; i++ {
		Chopsticks[i] = make(chan bool, 1)
		Chopsticks[i] <- true
	}
}

func initRandGen() {
	rgen = rand.New(rand.NewSource(time.Now().Unix()))
	randMutex = sync.Mutex{}
}

func (p *Philosopher) Eat() {
	randMutex.Lock()
	r := (EAT_MIN_MS) + EAT_MAX_MS - EAT_MIN_MS
	randMutex.Unlock()

	fmt.Println("starting to eat", p.Id)
	p.incrementCounter()
	time.Sleep(time.Duration(r) * time.Nanosecond)
	fmt.Println("finishing eating", p.Id)
}

func (p *Philosopher) Think() {
	randMutex.Lock()
	r := (THINK_MIN_MS) + THINK_MAX_MS - THINK_MIN_MS
	randMutex.Unlock()

	time.Sleep(time.Duration(r) * time.Nanosecond)
}

func (self *Philosopher) incrementCounter() {
	self.counter++
}

func (p *Philosopher) ReplaceChopsticks() {
	Chopsticks[p.Left_index] <- true
	Chopsticks[p.Right_index] <- true
	ChopsMutex.Signal()
}

func (h *Host) AllowsDiner(p *Philosopher, wg *sync.WaitGroup) {
	ChopsMutex.L.Lock()
	for {
		select {
		case _ = <-Chopsticks[p.Left_index]:
			_ = <-Chopsticks[p.Right_index]
			ChopsMutex.L.Unlock()
			wg.Done()
			return
		default:
			ChopsMutex.Wait()
		}
	}
}

func Run(p *Philosopher, host *Host, wg *sync.WaitGroup) {
	for eatTime := 3; eatTime > 0; eatTime-- {
		var wgroup sync.WaitGroup
		p.Think()
		wgroup.Add(1)
		go host.AllowsDiner(p, &wgroup)
		wgroup.Wait()
		p.Eat()
		p.ReplaceChopsticks()
	}
	wg.Done()
}
