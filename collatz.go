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
var SIX *(big.Int) = big.NewInt(6)

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

	int_n_start := 325000001         // should be the odd number >= 1
	n_limit := big.NewInt(326000000)

	int_n_start = int_n_start/6
	int_n_start = int_n_start*6+1
	n_start := big.NewInt(int64(int_n_start))
	var STEP *(big.Int) = big.NewInt(2)

	for n := n_start; n.Cmp(n_limit) == -1; n.Add(n, STEP) {
		STEP = STEP.Sub(SIX, STEP)   // skip 6k+3 numbers
		n_started := n
		// fmt.Printf("%v n_started=%v STEP=%v\n", pp, n_started, STEP)
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
	
	primes_to_check := []int64{5, 7, 29, 41, 79, 89, 107, 109, 127, 149, 157, 179, 191, 199, 211, 307, 311, 337, 359, 397, 419, 433, 439, 443, 463, 479, 487, 521, 557, 569, 631, 643, 647, 653, 733, 773, 811, 823, 827, 853, 887, 929, 941, 977, 997, 1009, 1019, 1021, 1033, 1061, 1093, 1123, 1181, 1187, 1217, 1223, 1249, 1279, 1283, 1289, 1291, 1301, 1327, 1367, 1373, 1433, 1471, 1483, 1489, 1499, 1523, 1531, 1549, 1559, 1567, 1601, 1607, 1609, 1613, 1619, 1627, 1657, 1667, 1669, 1699, 1733, 1747, 1753, 1787, 1811, 1831, 1861, 1867, 1933, 1949, 1987, 1993, 2003, 2029, 2039, 2081, 2087, 2089, 2113, 2141, 2143, 2161, 2207, 2237, 2239, 2243, 2251, 2267, 2269, 2273, 2281, 2351, 2357, 2389, 2399, 2411, 2423, 2441, 2503, 2539, 2591, 2633, 2657, 2753, 2777, 2819, 2837, 2851, 2857, 2861, 2887, 2917, 2927, 2953, 2969, 2971, 3011, 3023, 3037, 3041, 3049, 3079, 3083, 3109, 3209, 3251, 3257, 3259, 3271, 3313, 3323, 3347, 3359, 3371, 3373, 3389, 3413, 3433, 3449, 3467, 3469, 3491, 3547, 3557, 3571, 3581, 3583, 3607, 3613, 3617, 3659, 3677, 3701, 3709, 3719, 3733, 3761, 3793, 3821, 3823, 3833, 3847, 3863, 3877, 3881, 3889, 4021, 4057, 4079, 4091, 4093, 4153, 4201, 4211, 4217, 4219, 4271, 4283, 4327, 4349, 4357, 4421, 4441, 4451, 4483, 4493, 4513, 4523, 4567, 4603, 4651, 4657, 4663, 4673, 4729, 4733, 4783, 4787, 4799, 4813, 4817, 4877, 4889, 4903, 4919, 4937, 4969, 4993, 5039, 5087, 5101, 5107, 5113, 5227, 5231, 5237, 5279, 5303, 5309, 5323, 5381, 5419, 5477, 5519, 5557, 5563, 5569, 5581, 5623, 5647, 5653, 5659, 5689, 5693, 5711, 5717, 5737, 5743, 5783, 5807, 5813, 5857, 5881, 5897, 5923, 5939, 5987, 6007, 6011, 6037, 6047, 6067, 6073, 6113, 6121, 6143, 6151, 6173, 6199, 6217, 6221, 6229, 6263, 6311, 6353, 6367, 6373, 6397, 6521, 6529, 6599, 6607, 6653, 6659, 6661, 6673, 6679, 6701, 6719, 6737, 6763, 6781, 6793, 6803, 6823, 6827, 6829, 6857, 6871, 6883, 6899, 6959, 6967, 6977, 6991, 7027, 7043, 7057, 7079, 7127, 7129, 7151, 7193, 7219, 7253, 7297, 7309, 7333, 7351, 7393, 7417, 7481, 7489, 7517, 7523, 7541, 7549, 7559, 7561, 7573, 7589, 7603, 7639, 7643, 7669, 7673, 7687, 7703, 7741, 7753, 7789, 7793, 7817, 7829, 7853, 7873, 7879}
	fmt.Printf("check size = %d\n", len(primes_to_check))
	for q := 0; q < len(primes_to_check); q++ {
		qq := *(big.NewInt(primes_to_check[q]))
		wg.Add(1)
		go check_p(&qq)
	}
	
	wg.Wait()
	fmt.Printf("finished ")
	fmt.Println(time.Now().Format(time.RFC850))
}

