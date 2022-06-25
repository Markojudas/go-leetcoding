<!-- markdownlint-disable -->

# Decrypt String from Alphabet to Integer Mapping

You are given a string `s` formed by digits and `'#'`. We want to map `s` to English lowercase characters as follows:

<ul>
    <li>Characters (<code>'a'</code> to <code>'i'</code>) are represented by (<code>'1'</code> to <code>'9'</code>) respectively.</li>
    <li>Characters (<code>'j'</code> to <code>'z'</code>) are represented by (<code>'10#'</code> to <code>'26#'</code>) respectively.</li>   
</ul>

Return the string formed after mapping.

The test cases are generated so that a unique mapping will always exist.
<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>          s = "10#11#12"
<strong>Output:</strong>         "jkab"
<strong>Explanation:</strong>    "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".</code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>          s = "1326#"
<strong>Output:</strong>         "acz"</code></pre>

<br>
**Constraints:**

<ul>
    <li><code>1 <= s.length <= 1000</code></li>
    <li><code>s</code> consists of digits and the <code>'#'</code> letter.</li>
    <li><code>s</code> will be a valid string such that mapping is always possible</li>

</ul>
