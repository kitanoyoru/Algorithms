// Solved by @kitanoyoru
// https://leetcode.com/problems/linked-list-in-binary-tree/

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

const check = (head: ListNode | null, root: TreeNode | null): boolean => {
  if (!head) return true
  if (!root) return false

  return (
    head.val === root.val &&
    (check(head.next, root.left) || check(head.next, root.right))
  )
}

const isSubPath = (head: ListNode | null, root: TreeNode | null): boolean => {
  if (!head) return true
  if (!root) return false

  return (
    check(head, root) ||
    isSubPath(head, root.left) ||
    isSubPath(head, root.right)
  )
}
