import java.util.*;

public class BFSExample {
    static int[][] graph = {
        {1, 1, 0, 1},
        {1, 0, 1, 1},
        {1, 1, 1, 0},
        {0, 1, 1, 1}
    };
    static boolean[][] visited;
    static int n = 4; // 행
    static int m = 4; // 열

    // 방향: 상, 하, 좌, 우
    static int[] dx = {-1, 1, 0, 0};
    static int[] dy = {0, 0, -1, 1};

    public static void main(String[] args) {
        visited = new boolean[n][m];
        bfs(0, 0); // 시작점 (0, 0)
    }

    static void bfs(int x, int y) {
        Queue<int[]> queue = new LinkedList<>();
        queue.offer(new int[]{x, y});
        visited[x][y] = true;

        while (!queue.isEmpty()) {
            int[] now = queue.poll();
            int cx = now[0];
            int cy = now[1];

            System.out.println("방문: (" + cx + ", " + cy + ")");

            for (int i = 0; i < 4; i++) {
                int nx = cx + dx[i];
                int ny = cy + dy[i];

                if (nx < 0 || ny < 0 || nx >= n || ny >= m) continue; // 범위 밖
                if (visited[nx][ny]) continue; // 이미 방문
                if (graph[nx][ny] == 0) continue; // 갈 수 없는 길

                visited[nx][ny] = true;
                queue.offer(new int[]{nx, ny});
            }
        }
    }
}

