class Solution:
    def minCost(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]  # right, left, down, up
        costs = {(0, 0): 0}
        queue = [(0, 0, 0)]  # format: (cost, row, column)

        while queue:
            cost, r, c = heappop(queue)

            # If we've reached the bottom-right cell, return the cost to get there
            if r == m - 1 and c == n - 1:
                return cost

            # Explore the four possible directions
            for i, (dr, dc) in enumerate(directions):
                nr, nc = r + dr, c + dc

                # Check if the new position is within bounds
                if 0 <= nr < m and 0 <= nc < n:
                    # Determine the cost to move to the new cell
                    next_cost = cost if grid[r][c] == i + 1 else cost + 1

                    # If the new cost is lower than any previously recorded cost for this cell
                    if (nr, nc) not in costs or next_cost < costs[(nr, nc)]:
                        costs[(nr, nc)] = next_cost
                        heappush(queue, (next_cost, nr, nc))

        return -1  # This should never be reached because there is always at least one path
