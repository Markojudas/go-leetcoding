<!-- markdownlint-disable -->

# Richest Customer Wealth

You are given an `m x n` integer grid `accounts` where `accounts[i][j]` is the amount of money the <code>i<sup>th</sup></code> customer has in the <code>j<sup>th</sup></code> bank. Return <em>the<strong>wealth</strong></em> that the richest customer has.

A customer's **wealth** is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum **wealth**
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          accounts = [[1,2,3],[3,2,1]]
<strong>Output:</strong>         6
<strong>Explanation:</strong>    
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          accounts = [[1,5],[7,3],[3,5]]
<strong>Output:</strong>         10
<strong>Explanation:</strong>    
1st customer has wealth = 6
2nd customer has wealth = 10 
3rd customer has wealth = 8
The 2nd customer is the richest with a wealth of 10.</code></pre>

**Example 3**:

<pre><code><strong>Input:</strong>          accounts = [[2,8,7],[7,1,3],[1,9,5]]
<strong>Output:</strong>         17</code></pre>

**Constraints:**

<ul>
<li><code>m == accounts.length</code></li>
<li><code>n == accounts[i].length</code></li>
<li><code>1 <= m, n <= 50</code></li>
<li><code>1 <= accounts[i][j] <= 100</code></li>
</ul>
