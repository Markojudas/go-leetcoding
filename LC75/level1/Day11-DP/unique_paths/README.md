<!-- markdownlint-disable -->

# Unique Paths

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][1]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point of time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.<br>
<br>

**Example 1:**

<img src="./img/robot_maze.png">

<pre><code><strong>Input:</strong>      m = 3, n = 7
<strong>Output:</strong>     28</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         m = 3, n = 2
<strong>Output:</strong>        3
<strong>Explanation:</strong>   From the top-left corner, there are a total of 
3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= m, n <= 100</code></li>
</ul>
