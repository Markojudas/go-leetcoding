<!-- markdownlint-disable -->

# Bulls and Cows

You are playing the <a href="https://en.wikipedia.org/wiki/Bulls_and_Cows">Bulls and Cows</a> game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess you provide a hint with the following info:

<ul>
    <li>The number of "bulls", which are digits in the guess that are in the correct position.</li>
    <li>The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.</li>
</ul>

Given the secret number `secret` and your friend's guess `guess`, return the hint for your friend's guess.

The hint should be formatted as `xAyB`, where `x` is the number of bulls and `y` is the number of cows. Note that both `secret` and `guess` may contain duplicate digits.<br>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>         secret = "1807", guess = "7810"
<strong>Output:</strong>        "1A3B"
<strong>Explanation:</strong>   Bulls are connected with a '|' and cows are underlined:
"1807"
  |
"<u>7</u>8<u>10</u>"</code></pre>
<br>

**Example 2:**

<pre><code><strong>Input:</strong>         secret = "1123", guess = "0111"
<strong>Output:</strong>        "1A1B"
<strong>Explanation:</strong>   Bulls are connected with a '|' and cows are underlined:
"1123"        "1123"
  |      or     |
"01<u>1</u>1"        "011<u>1</u>"
Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.</code></pre>
<br>
<br>

**Constraints:**

<ul>
    <li><code>1 <= secret.length, guess.length <= 1000</code></li>
    <li><code>secret.length == guess.length</code></li>
    <li><code>secret</code> and <code>guess</code> consist of digits only</li>
</ul>
