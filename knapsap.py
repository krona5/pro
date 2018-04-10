
def solution(vals, weights, limit):

    ratio = [vals[i]/weights[i] for i in range(len(vals))]
    curr = 0

    print(ratio)

    while True:
        max_ratio = 0
        max_at = -1
        for i, v in enumerate(ratio):
            if v is None:
                break
            elif v > max_ratio:
                max_ratio = v
                max_at = i

        if max_at == -1:
            break

        if curr + weights[max_at] <= limit:
            curr += vals[max_at]

        ratio[max_at] = None

    return curr

# CASE 1 (a:3)
vals = [1, 2, 3]
wts = [4, 5, 1]
lim = 4

solution(vals, wts, lim)

print("Case #1: val({}), weights({}), limit({}) : {}".format(vals, wts, lim, solution(vals, wts, lim)))

# CASE 1 (a:220)
vals = [60, 100, 120]
wts = [10, 20, 30]
lim = 50

print("Case #2: val({}), weights({}), limit({}) : {}".format(vals, wts, lim, solution(vals, wts, lim)))
