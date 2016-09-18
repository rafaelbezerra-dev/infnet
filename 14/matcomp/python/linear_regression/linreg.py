import numpy as np
import matplotlib.pyplot as plt

pow2 = lambda x: x**2
mean = lambda x: sum(x)/len(x)

def a1(x, y):
    n = len(x)
    sumx, sumy, sumxy, sumx2 = 0, 0, 0, 0
    for i in range(n):
        sumx += x[i]
        sumy += y[i]
        sumxy += x[i] * y[i]
        sumx2 += pow2(x[i])

    return ((n * sumxy) - (sumx * sumy))/((n * sumx2) - pow2(sumx))

def a0(x, y, a1):
    return mean(y) - a1 * mean(x)


def main():
    x = [  1,   2,   3,   4,   5,   6,   7]
    y = [0.5, 2.5, 2.0, 4.0, 3.5, 6.0, 5.5]
    _a1 = a1(x, y)
    _a0 = a0(x, y, _a1)
    print('_a1:', _a1, ', _a0:', _a0)
    print('y = '+ str(_a1) + '*x + ' + str(_a0))

    n = len(x)
    st, sr = 0, 0
    meany = mean(y)
    for i in range(n):
        st = st + pow2(y[i] - meany)
        sr = sr + pow2(y[i] - _a1*x[i] - _a0)

    syx = pow(sr/(n - 2), 0.5)
    r2 = (st - sr)/st
    print('st:', st)
    print('sr:', sr)
    print('syx:', syx)
    print('r²:', r2)

    sol = lambda n: _a1 * n + _a0

    plt.plot(x, y, 'ro')
    t1 = np.arange(min(x) - 1, max(x) + 1, 0.1)
    plt.plot(t1, sol(t1), 'k')

    plt.axis([min(x) - 1, max(x) + 1, min(y) - 1, max(y) + 1])
    plt.show()


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
