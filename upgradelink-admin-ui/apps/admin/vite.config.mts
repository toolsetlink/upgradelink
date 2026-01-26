import { defineConfig } from '@vben/vite-config';

export default defineConfig(async () => {
  return {
    application: {},
    vite: {
      server: {
        proxy: {
          '/api': {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/api/, ''),
            // 代理目标地址
            target: 'http://localhost:9110',
            ws: true,
            proxyTimeout: 600_000, // 代理连接超时时间
          },
        },
      },
    },
  };
});
