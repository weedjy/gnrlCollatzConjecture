# coding=utf-8

#primes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127]
from math import sqrt

primes = [2]

count = 1
i = 1
while count < 1000:
    i += 2
    for k in range(2, 1+int(sqrt(i+1))):
        if i%k == 0:       
            break
    else:
        primes.append(i)
        count += 1
print(primes)

def next_step(ntt, ps):
    for q in primes:
        if q < ps:
            if (ntt // q) * q == ntt:
                return ntt // q
        else:
            break
    return ntt * ps + 1

def check_p(pp):
    print("exploring ", p)
    for n in range(5000001, 5500000, 2):
        n_started = n
        nt = n
        n_set = set()
        n_set.add(nt)
        sl = len(n_set)
        # print("checking ", pp, nt, '\t\t', end='')
        # print("checking ", pp, nt)
        while True:
            nt = next_step(nt, pp)
            n_set.add(nt)
            # print(nt, "->", end='')
            if nt == 1:
                break
            elif nt == n_started:
                print("=n with ", pp, n)
                break
            elif nt < n_started:
                break
            elif len(n_set) == sl and n_started not in n_set:
                print(">n with ", pp, n)
                # print(n_set)
                break
            elif len(n_set) == sl and n_started in n_set:
                print("=n with ", pp, n)
                # print(n_set)
                break
            else:
                sl = len(n_set)
                if sl > 10000:
                  print("warn: sl={0} for n_started={1}", sl, n_started)

for p in primes[2:]:
    check_p(p)

