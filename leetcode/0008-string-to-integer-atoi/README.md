#### String to Integer (atoi)

Implement `atoi` which converts a string to an integer.

The function first discards as many whitespace characters as necessary until
the first non-whitespace character is found. Then, starting from this
character, takes an optional initial plus or minus sign followed by as many
numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral
number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid
integral number, or if no such sequence exists because either str is empty or
it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

**Note**:
* Only the space character ' ' is considered as whitespace character.
* Assume we are dealing with an environment which could only store integers
  within the 32-bit signed integer range: [−2<sup>31</sup>,  2<sup>31</sup> −
  1]. If the numerical value is out of the range of representable values,
  INT\_MAX (2<sup>31</sup> − 1) or INT\_MIN (−2<sup>31</sup>) is returned.

**Example 1**:
<pre><code><b>Input</b>: "42"
<b>Output</b>: 42
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: "   -42"
<b>Output</b>: -42
<b>Explanation</b>: The first non-whitespace character is '-', which is the minus sign.  Then take
             as many numerical digits as possible, which gets 42.
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: "4193 with words"
<b>Output</b>: 4193
<b>Explanation</b>: Conversion stops at digit '3' as the next character is not a numerical digit.
</code></pre>

**Example 4**:
<pre><code><b>Input</b>: "words and 987"
<b>Output</b>: 0
<b>Explanation</b>: The first non-whitespace character is 'w', which is not a numerical digit or a
             +/- sign. Therefore no valid conversion could be performed.  </code></pre>

**Example 5**:
<pre><code><b>Input</b>: "-91283472332"
<b>Output</b>: -2147483648
<b>Explanation</b>: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−2<sup>31</sup>) is returned.
</code></pre>
