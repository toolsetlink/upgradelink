import { defineConfig } from "@vben/vite-config";

export default defineConfig(async () => {
  return {
    application: {},
    vite: {
      server: {
        proxy: {
          "/fms-api": {
            changeOrigin: true,
            proxyTimeout: 600_000, // 代理连接超时时间
            rewrite: (path) => path.replace(/^\/fms-api/, ""),
            // target: "http://0.0.0.0:9102/",
            target: "http://fms.upgrade.toolsetlink.com",
            timeout: 600_000, // 代理服务器超时时间
            ws: true,
          },
          "/sys-api": {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/sys-api/, ""),
            target: "http://core.upgrade.toolsetlink.com",
            // target: "http://0.0.0.0:9100",
            ws: true,
          },
          "/upgrade": {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/upgrade/, ""),
            // mock代理目标地址
            target: "http://upgrade.upgrade.toolsetlink.com",
            // target: "http://0.0.0.0:8181",
            ws: true,
          },
        },
      },
    },
  };
});
