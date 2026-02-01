// @ts-check
import { defineConfig } from 'astro/config';

import react from '@astrojs/react';

import mdx from "@astrojs/mdx";
import rehypeFigure from 'rehype-figure';

// https://astro.build/config
export default defineConfig({
  integrations: [react(), mdx()],
  markdown: {
    rehypePlugins: [
      rehypeFigure
    ]
  },
  site: "https://mbell.dev",
  redirects: {
    "/blog": "/blog/1",
  },
  vite: {
    server: {
      allowedHosts: [
        "maisie-desktop",
        "maisie-desktop.local",
        "macsie.local",
        "macsie.corvus-buri.ts.net"
      ]
    }
  }
});
