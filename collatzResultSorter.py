# coding=utf-8
import codecs
import collections
import matplotlib.pyplot as plt


xml_file = codecs.open("collatz.txt", "r", "utf-8")
results = xml_file.readlines()
xml_file.close()
results.sort(reverse=True)
checked_primes_set = set()
primes_to_check = set()
cycles = dict()
for ln in results:
	words = ln.strip(' \n\r')
	words = words.split()
	if words[0] == u'exploring':
		current_prime = int(words[1])
		checked_primes_set.add(current_prime)
		primes_to_check.add(current_prime)
		cycles[current_prime] = []
	else:
		checked_prime = int(words[2])
		if checked_prime in primes_to_check:
			primes_to_check.remove(checked_prime)
		cycles[checked_prime].append(int(words[3]))

primes_to_check_file = codecs.open("primes-to-check.txt", "w", "utf-8")
output = list(primes_to_check)
output.append(4079)  # very:) interesting number, continue to check it
output.sort()
primes_to_check_file.write(u'size = ' + str(len(output)) + u'\n')
primes_to_check_file.write(str(output))
primes_to_check_file.close()

sorted_cycles = collections.OrderedDict()
for key in sorted(cycles.keys()):
	sorted_cycles[key] = sorted(cycles[key])


for key, val in sorted_cycles.items():
	if len(val) > 0:
		print(key, val)


x = collections.OrderedDict()
for key in sorted_cycles.keys():
	x[key] = len(sorted_cycles[key])


plt.bar(x.keys(), x.values(), width=20.0, color='b')
plt.show()
