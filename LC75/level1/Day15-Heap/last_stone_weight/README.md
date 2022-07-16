<!-- markdownlint-disable -->

# Last Stone Weight

You are given an array of integers `stones` where `stones[i]` is the weight of the <code>i<sup>th</sup></code> stone.

We are playing a game with the stones. On each turn, we choose the **heaviest two stones** and smash them together. Suppose the heaviest two stones have weights `x` and `y` with `x <= y`. The result of this smash is:

<ul>
    <li>If <code>x == y</code>, both stones are destroyed, and</li>
    <li>If <code>x != y</code>, the stone of weight <code>x</code> is destroyed, and the stone of weight <code>y</code> has new weight <code>y - x</code>.</li>
</ul>

At the end of the game, there is **at most one** stone left.

Return the <em>the weight of the last remaining stone</em>. If there are no stones left, return 0.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>     stones = [2,7,4,1,8,1]
<strong>Output:</strong>    1
<strong>Explanation:</strong>
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>     stones = [1]
<strong>Output:</strong>    1</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= stones.length <= 30</code></li>
    <li><code>1 <= stones[i] <= 1000</code></li>
</ul>
