import GhostContentAPI, { type PostOrPage } from "@tryghost/content-api";
import rehypeParse from "rehype-parse";
import rehypeRemark from "rehype-remark";
import remarkGfm from "remark-gfm";
import remarkStringify from "remark-stringify";
import { sys } from "typescript";
import { unified } from "unified";
import { cwd } from "process";
import { join } from "path";
import { access } from "fs/promises";
import remarkFrontmatter from "remark-frontmatter";
import { Document, parseDocument } from "yaml";
import { write } from "to-vfile";
const ghostClient = new GhostContentAPI({
  url: "https://bell-blog.ghost.io",
  key: process.env.CONTENT_API_KEY || "",
  version: "v6.0",
});

const postsPath = join(cwd(), "src", "posts");
const LIMIT = 1;

const setFrontmatter = (post: PostOrPage) => () => (ast) => {
  const frontmatter = new Document({
    title: post.title,
    authors: post.authors?.map((a) => a.name),
    publishedTime: post.published_at,
    modifiedTime: post.updated_at,
    featureImage: post.feature_image || null,
    featureImageAlt: post.feature_image_alt || null,
    tags: post.tags?.map((t) => t.name),
    excerpt: post.custom_excerpt || null,
    readingTime: post.reading_time || 0,
  });
  ast.children.unshift({
    type: "yaml",
    value: String(frontmatter),
  });
};

// Takes a Ghost post and converts to Markdown.
const postToMD = async (post: PostOrPage) =>
  unified()
    .use(rehypeParse)
    .use(rehypeRemark)
    .use(remarkGfm)
    .use(remarkStringify)
    .use(remarkFrontmatter)
    .use(setFrontmatter(post))
    .process(post.html || "");

const posts = await ghostClient.posts.browse({
  include: ["authors", "tags"],
  limit: LIMIT,
  formats: ["html"],
});

if (posts.length == 0) {
  console.error("no posts returned");
  sys.exit(1);
}

const results = (await Promise.all(posts.map(postToMD))).forEach(
  async (f, i) => {
    const post = posts[i];
    try {
      await access(f.path);
      console.error("file exists, skipping");
      return;
    } catch {
      console.log(`writing to ${f.path}`);
      await write(f);
    }
  },
);
