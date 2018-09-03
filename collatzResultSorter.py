# coding=utf-8
import codecs

xml_file = codecs.open("collatz.txt", "r", "utf-8")
results = xml_file.readlines()
xml_file.close()
results.sort(reverse=True)
checked_primes_set = set()
primes_to_check = set()
for ln in results:
	words = ln.strip(' \n\r')
	words = words.split()
	if words[0] == u'exploring':
		checked_primes_set.add(int(words[1]))
		primes_to_check.add(int(words[1]))
	else:
		item_to_delete = int(words[2])
		if item_to_delete in primes_to_check:
			primes_to_check.remove(item_to_delete)

primes_to_check_file = codecs.open("primes-to-check.txt", "w", "utf-8")
output = list(primes_to_check)
output.append(4079)  # very:) interesting number, continue to check it
output.sort()
primes_to_check_file.write(u'size = ' + str(len(output)) + u'\n')
primes_to_check_file.write(str(output))
primes_to_check_file.close()
