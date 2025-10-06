import GhostContentAPI, { type PostsOrPages } from "@tryghost/content-api";
const ghostClient = new GhostContentAPI({
  url: 'https://bell-blog.ghost.io',
  key: import.meta.env.CONTENT_API_KEY,
  version: 'v6.0',
});

export const posts: void | PostsOrPages = await ghostClient.posts
  .browse({
    limit: 5,
  })
  .catch((err) => {
    console.error(err);
  })

