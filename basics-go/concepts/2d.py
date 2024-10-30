def traverse_grid(grid, row, col, visited):
    # Check if out of bounds or already visited
    if (
        row < 0
        or row >= len(grid)
        or col < 0
        or col >= len(grid[0])
        or visited[row][col]
    ):
        return

    # Mark the cell as visited
    visited[row][col] = True
    print(f"Visiting cell: {row}, {col}")

    # Explore in all 4 directions: up, down, left, right
    traverse_grid(grid, row + 1, col, visited)  # Down
    traverse_grid(grid, row - 1, col, visited)  # Up
    traverse_grid(grid, row, col + 1, visited)  # Right
    traverse_grid(grid, row, col - 1, visited)  # Left


# Example usage
grid = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
visited = [[False] * len(grid[0]) for _ in range(len(grid))]
traverse_grid(grid, 0, 0, visited)
