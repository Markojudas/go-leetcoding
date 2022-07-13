<!-- markdownlint-disable -->

# Min Cost Climbing Stairs

You are given an integer array `cost` where `cost[i]` is the cost of <code>i<sup>th</sup></code> step on a staircase. Once you pay the cost, you can either climb one of two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          cost = [10,<u>15</u>,20]
<strong>Output:</strong>         15
<strong>Explanation:</strong>    You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          cost = [<u>1</u>,100,<u>1</u>,1,<u>1</u>,100,<u>1</u>,<u>1</u>,100,<u>1</u>]
<strong>Output:</strong>         6
<strong>Explanation:</strong>    You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>2 <= cost.length <= 1000</code></li>
    <li><code>0 <= cost[i] <= 999</code></li>
</ul>
