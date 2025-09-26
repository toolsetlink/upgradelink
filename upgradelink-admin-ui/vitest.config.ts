import Vue from "@vitejs/plugin-vue";
import VueJsx from "@vitejs/plugin-vue-jsx";
import { configDefaults, defineConfig } from "vitest/config";

export default defineConfig({
optimizeDeps: {
    exclude: ['jiti']
},
  plugins: [Vue(), VueJsx()],
  test: {
    environment: "happy-dom",
    exclude: [...configDefaults.exclude, "**/e2e/**"],
  },
});
