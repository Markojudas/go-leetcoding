<!-- markdownlint-disable -->

# Number of Islands

Given an `m x n` 2D binary grid `grid` which represents a map of `1`s (land) and `0`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.<br>
<br>

**Example 1:**

<pre><code><strong>Input: <strong>grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
<strong>Output:</strong></code> 1</pre>
<br>

**Example 2:**

<pre><code><strong>Input: <strong>grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
<strong>Output:</strong></code> 3</pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>m == grid.length</code></li>
    <li><code>n == grid[i].length</code></li>
    <li><code>1 <= m, n <= 300</code></li>
    <li><code>grid[i][j]</code> is <code>'0'</code> or <code>'1'</code></li>
</ul>
