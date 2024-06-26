#### Check If N and Its Double Exist
Given an array `arr` of integers, check if there exists two integers `N` and `M` such that `N` is the double of `M` ( i.e. `N = 2 * M`).

More formally check if there exists two indices `i` and `j` such that :
* `i != j`
* `0 <= i, j < arr.length`
* `arr[i] == 2 * arr[j]`

**Example 1**:
<pre><code><b>Input</b>: arr = [10,2,5,3]
<b>Output</b>: true
<b>Explanation</b>: N = 10 is the double of M = 5, that is, 10 = 2 * 5.
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: arr = [7,1,14,11]
<b>Output</b>: true
<b>Explanation</b>: N = 14 is the double of M = 7, that is, 14 = 2 * 7.
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: arr = [3,1,7,11]
<b>Output</b>: false
<b>Explanation</b>: In this case does not exist N and M, such that N = 2 * M.
</code></pre>

**Constraints:**
* `2 <= arr.length <= 500`
* <code>-10<sup>3</sup> <= nums1.length <= 10<sup>3</sup></code>
