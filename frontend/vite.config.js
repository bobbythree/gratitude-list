import { defineConfig } from "vite";

export default defineConfig({
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8888",
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
