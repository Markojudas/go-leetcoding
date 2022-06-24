<!-- markdownlint-disable -->

# Merge Strings Alternately

You are given twon strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.<br>

<br>

**Example 1:**

<pre><code><strong>Input:</strong>          word1 = "abc", word2 = "pqr"
<strong>Output:</strong>         "apbqcr"
<strong>Explanation:</strong>    The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          word1 = "ab", word2 = "pqrs"
<strong>Output:</strong>         "apbqrs"
<strong>Explanation:</strong>    Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b 
word2:    p   q   r   s
merged: a p b q   r   s</code></pre>

**Example 3**:

<pre><code><strong>Input:</strong>          word1 = "abcd", word2 = "pq"
<strong>Output:</strong>         "apbqcd"
<strong>Explanation:</strong>    Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q 
merged: a p b q c   d</code></pre>

**Constraints:**

<ul>
<li><code>1 <= word1.length, word2.length <= 100</code></li>
<li><code>word1</code> and <code>word2</code> consist of lowercase English letters.</li>
</ul>
