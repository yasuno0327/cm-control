import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

dn = np.random.randint(2, size=1000)*2-1
p0 = 100
swalk = np.cumsum(dn)+p0

pd.Series(swalk).plot()
plt.ylabel("x(t+1)")
plt.xlabel("x(t)")
plt.show()