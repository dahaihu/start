def dijkstra(graph):
    costs, processed = dict(), set()
    costs['start'] = 0
    processed.add('start')
    for key, cost in graph['start'].items():
        costs[key] = cost

    node = find_lowest_cost_node(costs, processed)
    while node is not None:
        for next_node, cost in graph.get(node, dict()).items():
            new_cost = cost + costs[node]
            if new_cost < costs.get(next_node, float('inf')):
                costs[next_node] = new_cost
        processed.add(node)
        node = find_lowest_cost_node(costs, processed)
    return costs


def find_lowest_cost_node(costs, processed):
    lowest_cost = float('inf')
    lowest_node = None
    for node, cnt in costs.items():
        if node in processed:
            continue
        if cnt < lowest_cost:
            lowest_cost = cnt
            lowest_node = node
    return lowest_node


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
