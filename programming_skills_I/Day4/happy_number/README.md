<!-- markdownlint-disable -->

# Happy Number

Write an algorithm to determine if a number `n` is happy.

A **happy number** is a number defined by the following process:

<ul>
    <li>Starting with any positive integer, replace the number by the sum of the squares of its digits.</li>
    <li>Repeat the process until the number equals 1 (where it will stay), or it **loops endlessly in a cycle** which does not include 1.</li>
    <li>Those numbers for which this process **ends in 1** are happy.</li>
</ul>

Return `true` if `n` is a happy number, and `false` if not.

**Example 1:**

<pre><code><strong>Input:          n = 19</strong>
<strong>Output:</strong>         true
<strong>Explanation:</strong>
1<sup>2</sup> + 9<sup>2</sup> = 82
8<sup>2</sup> + 2<sup>2</sup> = 68
6<sup>2</sup> + 8<sup>2</sup> = 100
1<sup>2</sup> + 0<sup>2</sup> + 0<sup>2</sup> = 1</code></pre>

**Example 2:**

<pre><code><strong>Input:          n = 2</strong>
<strong>Output:</strong>         false</code></pre>

**Constraints:**

<ul>
<li><code>1 <= n <= 2<sup>31</sup> - 1</code></li>
</ul>
<br>
<hr>
<br>
<h2>Algorithm Used for Solution:</h2>
<br>
<a href="https://www.geeksforgeeks.org/floyds-cycle-finding-algorithm/" target="_blank">Floyd's Cycle Finding Algortihm</a>
