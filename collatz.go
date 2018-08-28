package main

import "fmt"
import "math"
import "math/big"

var primes []big.Int

func next_step(ntt *big.Int, ps *big.Int) *big.Int {
	for _, q := range primes {
		if q.Cmp(ps) == -1 {
			ZERO := big.NewInt(0)
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
	ans.Add(ans, big.NewInt(1))
	return ans
}

func check_p(pp *big.Int) {
    fmt.Println("exploring ", pp)
	n_limit := big.NewInt(100000)
	n_start := big.NewInt(1)
	n_step := big.NewInt(2)
	ONE := big.NewInt(1)
    for n := n_start; n.Cmp(n_limit) == -1; n.Add(n, n_step) {
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
                if sl > 10000 {
					fmt.Printf("warn: sl=%d for n_started=%v\n", sl, n_started)
					fmt.Println(pp, nt)
				}
			}
		}
	}
}

func main() {
	primes = make([]big.Int, 0)
	p2 := big.NewInt(2)
	primes = append(primes, *p2)
	count := 1
	i := 1
	for count < 1000 {
		i += 2
		k_lim := 1 + int(math.Sqrt(float64(i+1)))
		breaked := false
		for k := 2; k < k_lim; k++ {
			if (i%k == 0) {
				breaked = true
				break
			}
		}
		if breaked == false {
			pii := big.NewInt(int64(i))
			primes = append(primes, *pii)
			count++
		}
	}
	//fmt.Println(primes)
	for _, p := range primes[2:] {
	    check_p(&p)		
	}
}

