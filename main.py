import sys
# import networkx as nx
from common import read_input
from solver_greedy import distance, solve

# Pythonでも試した記録…

def distance_all(cities):
    distances = {}
    for i in range(len(cities)):
        for j in range(i+1, len(cities)):
            if i not in distances:
                distances[i] = {}
            if j not in distances:
                distances[j] = {}
            distances[i][j] = distance(cities[i], cities[j])
            distances[j][i] = distance(cities[i], cities[j])
    return distances

# def make_complete_graph(cities):
#     G = nx.complete_graph(len(cities))
#     distances = distance_all(cities)
#     for (u, v) in G.edges():
#         G[u][v]['weight'] = distances[u][v]
#     return G

def improve(cities, distances):    
    solution = solve(cities)
    print(total(solution, distances))
    print(len(solution))
    for i in range(len(solution)-2):
        u1 = solution[i]
        v1 = solution[i+1]
        for j in range(i+2, len(solution)):
            u2 = solution[j]
            v2 = solution[(j+1)%len(solution)]
            if distances[u1][u2] + distances[v1][v2] < distances[u1][v1] + distances[u2][v2]:
                l = i+1
                r = j
                while l < r:
                    solution[l], solution[r] = solution[r], solution[l]
                    l += 1
                    r -= 1
    return solution

                
def total(solution, distances):
    length = 0
    for i in range(len(solution)):
        length += distances[solution[i-1]][solution[i % len(solution)]]
    return length

if __name__ == '__main__':
    cities = read_input('input_%s.csv'%sys.argv[1])
    distances = distance_all(cities)
    solution = improve(cities, distances)
    print(len(solution))
    print(total(solution, distances))
    # print(distance_all(cities))
    # G = make_complete_graph(cities)
    # print(G.edges(data=True))


