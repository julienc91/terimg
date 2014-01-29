import sys

for i in xrange(256):
    sys.stdout.write("\033[48;5;" + str(i) + "m \033[m")
sys.stdout.write("\n")
