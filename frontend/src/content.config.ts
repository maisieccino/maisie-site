import { defineCollection, z } from "astro:content";
import { glob, file } from "astro/loaders";

const places = defineCollection({
  loader: glob({ pattern: "**/*.md", base: "./src/places" }),
  schema: z.object({
    name: z.string(),
    latitude: z.number(),
    longitude: z.number(),
    type: z.enum(["coffee-shop"]).default("coffee-shop"),
    url: z.string().optional(),
    address: z.string().optional(),
  })
})

export const collections = { places }
