from mpl_toolkits import mplot3d
import numpy as np
import matplotlib.pyplot as plt

clust0x = []
clust0y = []
clust0z = []

clust1x = []
clust1y = []
clust1z = []

clust2x = []
clust2y = []
clust2z = []

while True:
    try:
        val = input().split();
        x = float(val[0])
        y = float(val[1])
        z = float(val[2])
        clust = int(val[3])
        print(x, y, z, clust)
        if clust == 0:
            clust0x.append(x)
            clust0y.append(y)
            clust0z.append(z)
        elif clust == 1:
            clust1x.append(x)
            clust1y.append(y)
            clust1z.append(z)
        elif clust == 2:
            clust2x.append(x)
            clust2y.append(y)
            clust2z.append(z)
    except EOFError:
        break

fig = plt.figure()
ax = plt.axes(projection='3d')
ax.scatter3D(clust0x, clust0y, clust0z, color='r');
ax.scatter3D(clust1x, clust1y, clust1z, color='g');
ax.scatter3D(clust2x, clust2y, clust2z, color='b');
plt.show()