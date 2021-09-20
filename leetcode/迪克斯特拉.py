def dijkstra(graph):
    pass


def find_lowest_cost_node(costs, processed):
    pass


if __name__ == '__main__':
    graph = dict()
    graph["start"] = {}
    graph["start"]["a"] = 5
    graph["start"]["b"] = 2

    graph["a"] = {}
    graph["a"]["c"] = 4
    graph["a"]["d"] = 2

    graph["b"] = {}
    graph["b"]["a"] = 8
    graph["b"]["d"] = 7

    graph["c"] = {}
    graph["c"]["d"] = 6
    graph["c"]["end"] = 3

    graph['d'] = {}
    graph['d']['end'] = 1

    graph['end'] = {}
    costs = dijkstra(graph)
    print(costs)
