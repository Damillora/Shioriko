import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],

  css: {
    preprocessorOptions: {
      scss: {
        additionalData: '@use "src/variables.scss" as *;',
      },
    },
  },
  server: {
    allowedHosts: ["5173.shinosawa.nanao.moe"],
    proxy: {
      // string shorthand:
      // http://localhost:5173/foo
      //   -> http://localhost:4567/foo
      '/api': 'http://localhost:5050',
      '/data': 'https://booru.nanao.moe',
    }
  }
});
