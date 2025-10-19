import GhostContentAPI, { type PostOrPage, type PostsOrPages, type Tag } from "@tryghost/content-api";
import rehypeParse from "rehype-parse";
import rehypeRemark from "rehype-remark";
import remarkStringify from "remark-stringify";
import { unified } from "unified";
import astroConfig from "../../astro.config.mjs";
const ghostClient = new GhostContentAPI({
  url: 'https://bell-blog.ghost.io',
  key: import.meta.env.CONTENT_API_KEY,
  version: 'v6.0',
});
astroConfig.env

const limit = import.meta.env.MODE === "development" ? 50 : 100

export const posts: void | PostsOrPages = await ghostClient.posts
  .browse({
    include: ["authors", "tags"],
    limit: limit,
  })
  .catch((err) => {
    console.error(err);
  })


export const tags = await ghostClient.tags
  .browse({
    include: "count.posts",
    limit: 100,
  })
  .catch(err => {
    console.error(err)
  })

// Takes a Ghost post and converts to Markdown.
export const postToMD = async (post: PostOrPage) =>
  unified()
    .use(rehypeParse)
    .use(rehypeRemark)
    .use(remarkStringify)
    .process(post.html || "")
