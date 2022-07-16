<!-- markdownlint-disable -->

# Top K Frequent Words

Given an array of strings `words` and an integer `k`, return `k` most frequent strings.

Return the answer **sorted** by **the frequency** from highest to lowest. Sort the words with the same frequency by their **lexicographical order**<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         words = ["i","love","leetcode","i","love","coding"], k = 2
<strong>Output:</strong>        ["i","love"]
<strong>Explanation:</strong>   "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
<strong>Output:</strong>        ["the","is","sunny","day"]
<strong>Explanation:</strong>   "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= words.length <= 500</code></li>
    <li><code>1 <= words[i].length <= 10</code></li>
    <li><code>words[i]</code></li>
    <li><code>k</code> is in the range <code>[1, The number of <strong>unique</strong> words[i]]</code></li>
</ul>
