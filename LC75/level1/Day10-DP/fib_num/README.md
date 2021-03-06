<!-- markdownlint-disable -->

# Fibonacci Number

The **Fibonacci numbers**, commonly denoted `F(n)` form a sequence, called the **Fibonacci sequence**, such that each number is the sum of the two preceding ones, starting from `0` and `1`. That is,

<pre><code>F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1</code></pre>

Given `n`, calculate `F(n)`.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         n = 2
<strong>Output:</strong>        1
<strong>Explanation:</strong>   F(2) = F(1) + F(0) = 1 + 0 = 1</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         n = 3
<strong>Output:</strong>        2
<strong>Explanation:</strong>   F(3) = F(2) + F(1) = 1 + 1 = 2</code></pre>
<br>

**Example 3:**

<pre><code><strong>Input:</strong>         n = 4
<strong>Output:</strong>        3
<strong>Explanation:</strong>   F(4) = F(3) + F(2) = 2 + 1 = 3</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>0 <= n <= 30</code></li>
</ul>
