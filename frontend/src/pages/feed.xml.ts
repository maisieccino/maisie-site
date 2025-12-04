import rss, { type RSSFeedItem } from "@astrojs/rss";
import { posts as blogPosts, postToMD } from "../util/ghost"
import type { APIRoute } from "astro";
import type { PostOrPage } from "@tryghost/content-api";
import { createMarkdownProcessor } from "@astrojs/markdown-remark";
import { parseISO } from "date-fns";

export const GET: APIRoute = async (context) => {

  const postToFeedItem = async (post: PostOrPage): Promise<RSSFeedItem> => {
    const proc = await createMarkdownProcessor()
    const md = await postToMD(post)

    const render = await proc.render(String(md))

    return {
      author: (post.authors || []).join(","),
      content: render.code,
      categories: (post.tags || []).map(t => t.name || ""),
      link: `/post/${post.slug}`,
      title: post.title || "",
      pubDate: parseISO(post.published_at || ""),
      customData: `<image>
        <url>${post.feature_image}</url>
        <title>${post.feature_image_alt}</title>
      </image>`
    }
  }

  return rss({
    title: "Maisieccino",
    description: "Sharing longer thoughts & stories about tech, gender, health, climbing",
    site: context.site || "",
    items: await Promise.all(
      (blogPosts || [])
        .filter(isPost)
        .map(postToFeedItem)
    ),
  })
}
const isPost = (p: any): p is PostOrPage => {
  return p.uuid !== undefined;
}
