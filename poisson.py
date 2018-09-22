import datetime
import matplotlib.pyplot as plt
import numpy as np

# ポアソン分布、確率密度関数
def poisson(k, lamb):
    ret = np.e**(-lamb)
    for i in range(k):
        ret *= lamb / (i + 1)
    return ret

x = np.arange(0, 100)
y = [poisson(k, 50) for k in x]

plt.plot(x, y)
plt.xlabel("Request")

plt.show()

