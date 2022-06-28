<!-- markdownlint-disable -->

# Sort Integers by The Number of 1 Bits

<p>You are given an integer array <code>arr</code>. Sort the integers in the array in ascending order by the number of <code>1</code>'s in their binary representation and in case of two or more integers have the same number of <code>1<code>'s you have to sort them in ascending order.</p>
<p>Return the array after sorting it</p>
<br>

**Example 1:**

<pre><code><strong>Input:</strong>             arr = [0,1,2,3,4,5,6,7,8]
<strong>Output:</strong>            [0,1,2,4,8,3,5,6,7]
<strong>Explanation:</strong>       [0] is the only integer with 0 bits.
[1,2,4,8] all have 1 bit.
[3,5,6] have 2 bits.
[7] has 3 bits.
The sorted array by bits is [0,1,2,4,8,3,5,6,7]</code> </code></pre>

**Example 2:**

<pre><code><strong>Input:</strong>             arr = [1024,512,256,128,64,32,16,8,4,2,1]
<strong>Output:</strong>            [1,2,4,8,16,32,64,128,256,512,1024]
<strong>Explanation:</strong>       All integers have 1 bit in the binary representation, you should just sort them in ascending order.</code> </code></pre>
<br>

**Constraints:**

<ul>
    <li><code>1 <= arr.length <= 500</code></li>
    <li><code>0 <= arr[i] <= 10<sup>4</sup></code></li>
</ul>
