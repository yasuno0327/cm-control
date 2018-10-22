import matplotlib.pyplot as plt
import numpy as np

x = np.arange(0, 10, 0.1)
y = 2*x + 5

plt.plot(x, y)
plt.xlabel("x(t)")
plt.ylabel("x(t+1)")

plt.show()