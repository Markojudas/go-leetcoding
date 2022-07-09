<!-- markdownlint-disable -->

# Lowest Common Ancestor of a Binary Search Tree

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the <a href="https://en.wikipedia.org/wiki/Lowest_common_ancestor" target="_blank" rel="noopener noreferrer">definition of LCA on Wikipedia</a>: "The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**)."<br>
<br>

**Example 1:**

<img src="./img/binarysearchtree_improved.png">

<pre><code><strong>Input:</strong>       root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
<strong>Output:</strong>      6    
<strong>Explanation:</strong> The LCA of nodes 2 and 8 is 6.</code></pre>

<br>

**Example 2:**

<img src="./img/binarysearchtree_improved2.png">

<pre><code><strong>Input:</strong>       root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
<strong>Output:</strong>      2    
<strong>Explanation:</strong> The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.</code></pre>
<br>

**Example 3:**

<pre><code><strong>Input:</strong>       root = [2,1], p = 2, q = 1
<strong>Output:</strong>      2</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li>The number of nodes in the tree is in the range <code>[2, 10<sup>5</sup>]</code>.</li>
    <li><code>-10<sup>9</sup> <= Node.val <= 10</sup>9</sup></code></li>
    <li>All <code>Node.val</code> are <strong>unique</strong>.</li>
    <li><code>p != q</code></li>
    <li><code>p</code> and <code>q</code> will exist in the BST.</li>
<ul>
