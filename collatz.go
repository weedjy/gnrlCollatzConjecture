package main

import (
			"fmt"
			"math"
			"math/big"
			"math/rand"
			"time"
			"sync"
)

var wg sync.WaitGroup

var primes []big.Int
var ZERO *(big.Int) = big.NewInt(0)
var ONE *(big.Int) = big.NewInt(1)
var TWO *(big.Int) = big.NewInt(2)

func next_step(ntt *big.Int, ps *big.Int) *big.Int {
	for _, q := range primes {
		if q.Cmp(ps) == -1 {
			r := new(big.Int)
			r = r.Mod(ntt, &q)
			if r.Cmp(ZERO) == 0 {
				d := new(big.Int)
				d = d.Div(ntt, &q)
				return d
			}
		} else {
			break
		}
	}
	ans := new(big.Int)
	ans = ans.Mul(ntt, ps)
	ans.Add(ans, ONE)
	return ans
}

func check_p(pp *big.Int) {
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
	// fmt.Println("exploring ", pp) // useless when parallel run
	n_limit := big.NewInt(23000000)
	n_start := big.NewInt(22000001)
	for n := n_start; n.Cmp(n_limit) == -1; n.Add(n, TWO) {
		n_started := n
		nt := n
		n_set := map[string]bool{}
		n_set[nt.String()] = true
		sl := len(n_set)
		for {
			nt = next_step(nt, pp)
			n_set[nt.String()] = true
			_, n_started_in_n_set := n_set[n_started.String()]
			if nt.Cmp(ONE) == 0 {
				break
			} else if nt.Cmp(n_started) == 0 {
				fmt.Printf("=n with  %v %v\n", pp, n)
				break
			} else if nt.Cmp(n_started) == -1 {
				break
			} else if (len(n_set) == sl) && (n_started_in_n_set == false) {
				fmt.Printf(">n with  %v %v\n", pp, n)
				break
			} else if (len(n_set) == sl) && (n_started_in_n_set == true) {
				fmt.Printf("=n with  %v %v\n", pp, n)
				break
			} else {
				sl = len(n_set)
				/* if sl > 10000 {
					fmt.Printf("warn: sl=%d for n_started=%v\n", sl, n_started)
					// fmt.Println(pp, nt)
				}*/
				// useless in parallel run
			}
		}
	}
	wg.Done()
}

func main() {
	fmt.Printf("started ")
	fmt.Println(time.Now().Format(time.RFC850))
	primes = make([]big.Int, 0)
	primes = append(primes, *TWO)
	primes = append(primes, *(big.NewInt(3)))
	count := 1
	i := 1
	for count < 1000 {
		i += 2
		k_lim := 1 + int(math.Sqrt(float64(i+1)))
		breaked := false
		if i%2 == 1 {
			for k := 3; k <= k_lim; k += 2 {
				if (i%k == 0) {
					breaked = true
					break
				}
			}
		}
		if breaked == false {
			primes = append(primes, *(big.NewInt(int64(i))))
			count++
		}
	}
	// fmt.Println(primes)
	for _, p := range primes[2:] {
		qq := p
		wg.Add(1)
		go check_p(&qq)
	}
	wg.Wait()
	fmt.Printf("finished ")
	fmt.Println(time.Now().Format(time.RFC850))
}
