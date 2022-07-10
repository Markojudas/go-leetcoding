<!-- markdownlint-disable -->

# DFS Solution

Treat the 2d map as an undirected graph and there is an edge between two horizontally or vertically adjacent nodes of value '1'.

<h2>Algo</h2>

Linear scan the 2d grid map, if a node contains a '1', then it is a root node that triggers a Depth First Search. During DFS, every visited node should be set as '0' to mark as visited node. Count the number of root nodes that trigger DFS, this number would be the number of islands since each DFS starting at some root identifies an island.
