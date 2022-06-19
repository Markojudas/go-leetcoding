<!-- markdownlint-disable -->

# Find Nearest Point That Has the Same X or Y Coordinate

You are given two integers, `x` and `y`, which represent your current location on a Cartesian grid: `(x, y)`. You are also given an array `points` where each <code>points[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> represents that a point exists at <code>(a<sub>i</sub>, b<sub>i</sub>)</code>. A point is **valid** if it shares the same x-coordinate or the same y-coordinate as your location.

Return the index **(0-indexed)** of the **valid** point with the smallest **Manhattan distance** from your current location. If there are multiple, return the valid point with the **smallest** index. If there are no valid points, return `-1`.

The **Manhattan distance** between two points <code>(x<sub>1</sub>, y<sub>1</sub>)</code> and <code>(x<sub>2</sub>, y<sub>2</sub>)</code> is <code>abs(x<sub>1</sub> - x<sub>2</sub>) + abs(y<sub>1</sub> - y<sub>2</sub>)</code>.

**Example 1:**

<pre><code><strong>Input:          x = 3, y = 4, points = [[1,2],[3,1],[2,4],[2,3],[4,4]]</strong>
<strong>Output:</strong>         2
<strong>Explanation:</strong>    Of all the points, only [3,1], [2,4] and [4,4] are valid. Of the valid points, [2,4] and [4,4]
have the smallest Manhattan distance from your current location, with a distance of 1. [2,4] has the smallest index,
so return 2.</code></pre>

**Example 2:**

<pre><code><strong>Input:          x = 3, y = 4, points = [[3,4]]</strong>
<strong>Output:</strong>         0
<strong>Explanation:</strong>    The answer is allowed to be on the same location as your current location.</code></pre>

**Example 3:**

<pre><code><strong>Input:          x = 3, y = 4, points = [[2,3]]</strong>
<strong>Output:</strong>         -1
<strong>Explanation:</strong>    There are no valid points.</code></pre>
