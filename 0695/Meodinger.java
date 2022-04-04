class Meodinger {

    public int maxAreaOfIsland(int[][] grid) {
        int ans = 0, iLen = grid.length, jLen = grid[0].length;
        for (int i = 0; i < iLen; i++) {
            for (int j = 0; j < jLen; j++) {
                if (grid[i][j] == 0) continue;
                ans = Math.max(ans, findIsland(grid, i, j));
            }
        }
        return ans;
    }

    public int findIsland(int[][] ocean, int i, int j) {
        ocean[i][j] = 0;

        int acc = 1;

        if (i - 1 >= 0 && ocean[i - 1][j] == 1)
            acc += findIsland(ocean, i - 1, j);
        if (i + 1 < ocean.length && ocean[i + 1][j] == 1)
            acc += findIsland(ocean, i + 1, j);
        if (j - 1 >= 0 && ocean[i][j - 1] == 1)
            acc += findIsland(ocean, i, j - 1);
        if (j + 1 < ocean[0].length && ocean[i][j + 1] == 1)
            acc += findIsland(ocean, i, j + 1);

        return acc;
    }

}