from __future__ import division
import matplotlib.pyplot as plt
import networkx as nx

DG = nx.DiGraph()
edges = [(i, i+1, 1) for i in range(3)] + [(3, 0, 1)]
DG.add_weighted_edges_from(edges)
subax1 = plt.subplot(121)
nx.draw(DG, with_labels=True, font_weight='bold')
plt.show()