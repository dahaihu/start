def print_binary_tree():
    from binarytree import Node
    root = Node(5)
    root.left = Node(3)
    root.right = Node(7)
    root.left.left = Node(2)
    root.left.right = Node(4)
    root.right.left = Node(6)
    root.right.right = Node(8)
    print(root)


if __name__ == '__main__':
    print_binary_tree()
