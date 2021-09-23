def dijkstra(graph):
    costs, processed = dict(), set()
    processed.add('start')
    for child, cost in graph.get('start', dict()).items():
        costs[child] = cost
    node = find_lowest_cost_node(costs, processed)
    while node is not None:
        for next_node, next_cost in graph.get(node, dict()).items():
            costs[next_node] = min(
                costs.get(next_node, float('inf')),
                costs[node] + next_cost
            )
        processed.add(node)
        node = find_lowest_cost_node(costs, processed)
    return costs


def find_lowest_cost_node(costs, processed):
    lowest_cost = float('inf')
    lowest_cost_node = None
    for node, cost in costs.items():
        if node in processed:
            continue
        if cost < lowest_cost:
            lowest_cost = cost
            lowest_cost_node = node
    return lowest_cost_node


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
