import numpy as np
import matplotlib.pyplot as plt

pow2 = lambda x: x**2

def a1(x, y):
    n = len(x)
    sumx, sumy, sumxy, sumx2 = 0, 0, 0, 0
    for i in range(n):
        sumx += x[i]
        sumy += y[i]
        sumxy += x[i] * y[i]
        sumx2 += pow2(x[i])

    return ((n * sumXY) - (sumX * sumY))/((n * sumXSq) - pow2(sumX))

def a0(x, y, a1):
    mean = lambda x: sum(x)/len(x)
    return mean(y) - a1 * mean(x)


def main():
    x = [  1,   2,   3,   4,   5,   6,   7]
    y = [0.5, 2.5, 2.0, 4.0, 3.5, 6.0, 5.5]
    _a1 = a1(x, y)
    _a0 = a0(x, y, _a1)
    print(_a1, _a0)


if __name__ == '__main__':
    main()

# def f(t):
#     return np.exp(-t) * np.cos(2*np.pi*t)
#
# t1 = np.arange(0.0, 5.0, 0.1)
# t2 = np.arange(0.0, 5.0, 0.02)
#
# plt.figure(1)
# plt.subplot(211)
# plt.plot(t1, f(t1), 'bo', t2, f(t2), 'k')
#
# # plt.subplot(212)
# plt.plot(t2, np.cos(2*np.pi*t2), 'r--')
# plt.show()
