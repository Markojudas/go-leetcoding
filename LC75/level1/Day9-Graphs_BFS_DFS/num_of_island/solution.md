<!-- markdownlint-disable -->

# BFS Solution

Treat the 2d map as an undirected graph and there is an edge between two horizontally or vertically adjacent nodes of value '1'.

<h2>Algo</h2>

Linear scan the 2d grid map, if a node contains a '1', then it is a root node that triggers a Breadth First Search. Put it into a queue and set its value as '0' to mark as visited node. Iteratively search the neighbors of enqueued nodes until the queue becomes empty.
