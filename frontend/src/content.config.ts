import { defineCollection, z } from "astro:content";
import { glob } from "astro/loaders";

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
  // TODO: https://docs.astro.build/en/guides/content-collections/#building-a-custom-loader
  loader: glob({ pattern: "**/*.md", base: "./src/places" }),
  schema: z.object({
    slug: z.string(),
    title: z.string(),
    publishedTime: z.date(),
    modifiedTime: z.date(),
    authors: z.string().default("Maisie Bell"),
    tags: z.array(z.string()),
    excerpt: z.string().optional(),
    featureImage: z.string().optional(),
    featureImageAlt: z.string().optional()
  })
})

export const collections = { places, posts }
