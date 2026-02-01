import { defineCollection, z } from "astro:content";
import { glob, type Loader } from "astro/loaders";
import { posts as blogPosts, postToMD } from "./util/ghost";
import { parseISO } from "date-fns";

const globLoader = glob({ pattern: "**/*.md", base: "./src/posts" })

const postLoader: Loader = {
  name: "post-loader",
  async load(context) {
    await globLoader.load(context)
    const { store, parseData, logger, renderMarkdown } = context
    await Promise.all(
      (blogPosts || [])
        .map(post => (
          {
            post,
            data: {
              id: post.slug,
              slug: post.slug,
              title: post.title || "",
              publishedTime: parseISO(post.published_at || ""),
              modifiedTime: parseISO(post.updated_at || ""),
              authors: (post.authors || []).map(a => a.name),
              tags: (post.tags || []).map(t => t.name),
              excerpt: post.custom_excerpt || post.excerpt || "",
              featureImage: post.feature_image || "",
              featureImageAlt: post.feature_image_alt || "",
              rendered: post.html || "",
              readingTime: post.reading_time || 0,
            }
          }))
        .map(async post => {
          if (post.data.slug === "") {
            logger.warn(`No slug for post ${post.data.id}`)
            return
          }
          if (store.has(post.data.slug)) {
            return
          }
          const data = await parseData({ id: post.data.slug, data: post.data })
          logger.debug(`id is ${data.id}`)
          const md = await postToMD(post.post)
          store.set({
            id: post.data.slug,
            data,
            rendered: await renderMarkdown(String(md)),
          })
        })
    )
  },
}

const places = defineCollection({
  loader: glob({ pattern: "**/*.md", base: "./src/places" }),
  schema: z.object({
    name: z.string(),
    latitude: z.number(),
    longitude: z.number(),
    type: z.enum(["coffee-shop"]).default("coffee-shop"),
    url: z.string().optional(),
    address: z.string().optional(),
    image: z.string().optional()
  })
})
const posts = defineCollection({
  loader: postLoader,
  schema: z.object({
    title: z.string().describe("Title of the post"),
    publishedTime: z.coerce.date().default(new Date()).describe("Date the post was published"),
    modifiedTime: z.coerce.date().default(new Date()),
    authors: z.array(z.string()).default(["Maisie Bell"]).describe("Authors who contributed to the post"),
    tags: z.array(z.string()).describe("Tags that describe the post"),
    excerpt: z.string().optional().describe("A paragraph-sized excerpt from the post"),
    featureImage: z.string().optional().describe("An image to appear at the top of the post"),
    featureImageAlt: z.string().optional().describe("Description of the post's featured image"),
    readingTime: z.number().default(0).describe("Approximate time in minutes to read the post")
  })
})

export const collections = { places, posts }

