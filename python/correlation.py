import statistics
import math
import numpy as np
# import matplotlib.pyplot as plt

def big(median, data, n): # 1 = + -1 = -
    for value in data:
        if median > value:
            n.append(-1)
        elif median == value:
            n.append(0)
        else:
            n.append(1)

def small(data, n):
    for i in range(len(data)):
        if i == 0:
            continue
        else:
            if data[i] > data[i - 1]:
                n.append(1)
            elif data[i] < data[i - 1]:
                n.append(-1)
            else:
                n.append(0)

def sign_test(n, nplus, nminus, types):
    print(types)
    ns = n / 2 + 1.96 * math.sqrt(n) / 2 + 0.459
    print(ns, nplus, nminus)
    if ns <= nplus:
        print("正の相関")
    elif ns <= nminus:
        print("負の相関")
    else:
        print("相関なし")

x = [
        24.59, 24.56, 24.47, 24.52, 24.54, 24.53, 24.51, 24.52, 24.43, 24.58, 24.49, 24.38, 24.51, 24.53, 24.44, 24.57, 24.44, 24.54, 24.55, 24.49,
        24.51, 24.50, 24.52, 24.47, 24.49, 24.56, 24.51, 24.51, 24.53, 24.52, 24.47, 24.46, 24.52, 24.52, 24.45, 24.47, 24.52, 24.50, 24.47, 24.49
    ]
y = [
        3.55, 3.62, 3.34, 3.46, 3.52, 3.46, 3.39, 3.43, 3.29, 3.60, 3.40, 3.3, 3.44, 3.44, 3.27, 3.54, 3.29, 3.41, 3.55, 3.44,
        3.44, 3.42, 3.43, 3.33, 3.49, 3.61, 3.34, 3.48, 3.51, 3.49, 3.35, 3.36, 3.46, 3.49, 3.30, 3.30, 3.42, 3.31, 3.39, 3.44
    ]
z = [
        0.91, 0.81, 0.78, 0.79, 0.91, 0.83, 0.91, 1.09, 1.03, 1.58, 1.41, 1.16, 0.78, 1.02, 0.78, 0.87, 0.80, 0.85, 1.09, 0.87,
        1.33, 0.87, 0.95, 0.89, 0.98, 1.34, 1.03, 0.84, 0.89, 1.34, 1.14, 0.85, 0.77, 1.24, 0.83, 0.99, 1.39, 1.08, 1.03, 1.20
    ]

nx = []
ny = []
nz = []

medianx = statistics.median(x)
mediany = statistics.median(y)
medianz = statistics.median(z)

# big(medianx, x, nx)
# big(mediany, y, ny)
# big(medianz, z, nz)
small(x, nx)
small(y, ny)
small(z, nz)

nxy = np.array(nx) * np.array(ny)
nxz = np.array(nx) * np.array(nz)

plus_count_nxy = np.sum(nxy == 1)
minus_count_nxy = np.sum(nxy == -1)
count_nxy = plus_count_nxy + minus_count_nxy

plus_count_nxz = np.sum(nxz == 1)
minus_count_nxz = np.sum(nxz == -1)
count_nxz = plus_count_nxz + minus_count_nxz

sign_test(count_nxy, plus_count_nxy, minus_count_nxy, "外径、部品肉厚")
sign_test(count_nxz, plus_count_nxz, minus_count_nxz, "外径、保持時間")

# plt.plot(np.arange(40), z, '-o')
# plt.plot([0, 40], [medianz, medianz], "red", linestyle='dashed')
# plt.show()