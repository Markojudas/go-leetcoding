<!-- markdownlint-disable -->

# Linked List Cycle II

Given the `head` of a linked list, return the node where the cycle begins. if there is no cycle, return `nil`

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (**0-index**). It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter**.

**Do not modify** the linked list.<br>

**Example 1:**

<img src="./img/example1.png">

<pre><code><strong>Input:</strong>          head = [3,2,0,-4], pos = 1
<strong>Output:</strong>         tail connects to node index 1
<strong>Explanation:</strong>    There is a cycle in the linked list, where tail connects to the second node.</code></pre>
<br>

**Example 2:**

<img src="./img/example2.png">

<pre><code><strong>Input:</strong>          head = [1,2], pos = 0
<strong>Output:</strong>         tail connects to node index 0
<strong>Explanation:</strong>    There is a cycle in the linked list, where tail connects to the first node.</code></pre>
<br>

**Example 3:**

<img src="./img/example3.png">

<pre><code><strong>Input:</strong>          head = [1], pos = -1
<strong>Output:</strong>         no cycle
<strong>Explanation:</strong>    There is no cycle in the linked list.</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li>The number of nodes in the list is in the range <code>[0, 10<sup>4</sup>]</code>.</li>
    <li><code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code></li>
    <li><code>pos</code> is <code>-1</code> or a <strong>valid index</strong> in the linked list.</li>
</ul>
