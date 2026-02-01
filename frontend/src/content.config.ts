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
    title: z.string(),
    publishedTime: z.coerce.date(),
    modifiedTime: z.coerce.date(),
    authors: z.array(z.string()).default(["Maisie Bell"]),
    tags: z.array(z.string()),
    excerpt: z.string().optional(),
    featureImage: z.string().optional(),
    featureImageAlt: z.string().optional(),
    readingTime: z.number().default(0)
  })
})

export const collections = { places, posts }

