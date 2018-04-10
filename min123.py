# Goal:
# For N, find min step to get to 1 by following below rules:
# 1. Subtract 1
# 2. Divide by 2 if possible
# 3. Divide by 3 if possible


def min_path(n):
    mem = {1: 0}

    for i in range(2, n + 1):
        mem[i] = 1 + mem[i - 1]
        if i % 2 == 0:
            mem[i] = min(mem[i], 1 + mem[int(i / 2)])
        if i % 3 == 0:
            mem[i] = min(mem[i], 1 + mem[int(i / 3)])

    return mem[n]


print("Min path for 1: {}".format(min_path(1)))
print("Min path for 3: {}".format(min_path(3)))
print("Min path for 10: {}".format(min_path(10)))
