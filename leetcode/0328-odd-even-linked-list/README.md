#### Odd Even Linked List
Given the  `head`  of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return  _the reordered list_.

The  **first**  node is considered  **odd**, and the  **second**  node is  **even**, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in  `O(1)` extra space complexity and  `O(n)`  time complexity.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: head = [1,2,3,4,5]
<b>Output</b>: [1,3,5,2,4]
</code></pre>

**Example 2**:

![](example_2.jpg)
<pre><code><b>Input</b>: head = [2,1,3,5,6,4,7]
<b>Output</b>: [2,3,6,7,1,5,4]
</code></pre>

**Constraints**:
* `n ==` number of nodes in the linked list
* <code>0 <= n <= 10<sup>4</sup></code>
* <code>-10<sup>6</sup> <= Node.val <= 10<sup>6</sup>
