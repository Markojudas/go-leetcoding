<!-- markdownlint-disable -->

# Validate Binary Search Tree

Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

A **valid BST** is definied as follows:

<ul>
    <li>The left subtree of a node contains only nodes with keys <strong>less than</strong> the node's key.</li>
    <li>The right subtree of a node contains only nodes with keys <strong>greater than</strong> the node's key</li>
    <li>Both the left and right subtrees must also by binary search trees.</li>
</ul><br>
<br>

**Example 1:**

<img src="./img/tree1.jpg">

<pre><code><strong>Input:</strong>      root = [2,1,3]
<strong>Output:</strong>     true</code></pre>

<br>

**Example 2:**

<img src="./img/tree2.jpg">

<pre><code><strong>Input:</strong>         root = [5,1,4,null,null,3,6]
<strong>Output:</strong>        false
<strong>Explanation:</strong>   The root node's value is 5 but its right child's value is 4.</code></pre>

<br>
<br>

**Constraints:**

<ul>
    <li>The number of nodes in the tree is in the range of <code>[1, 10<sup>4</sup>]</code>.</li>
    <li><code>-2<sup>31</sup> <= Node.val <= 2<sup>31</sup> - 1</code></li>
</ul>
