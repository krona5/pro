
def fac_rec(number):
    if number <= 1:
        return 1
    else:
        return fac_rec(number - 1) * number


def fac(num):
    f = 1
    for i in range(1, num + 1):
        f *= i
    return f


print("Factorial 3: {}".format(fac(3)))
print("Factorial 10: {}".format(fac(10)))
