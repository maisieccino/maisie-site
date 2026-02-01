import rss from "@astrojs/rss";
import type { APIRoute } from "astro";
import { getImage } from "astro:assets";
import { getCollection } from "astro:content";

export const GET: APIRoute = async (context) => {
  const posts = await getCollection("posts")

  return rss({
    title: "Maisieccino",
    description: "Sharing longer thoughts & stories about tech, gender, health, climbing",
    site: context.site || "",
    items: await Promise.all(posts.sort((a, b) => b.data.publishedTime.valueOf() - a.data.publishedTime.valueOf())
      .map(async post => ({
        title: post.data.title,
        pubDate: post.data.publishedTime,
        description: post.data.excerpt || "",
        link: `/post/${post.id}/`,
        enclosure: {
          type: "image",
          url: (await featureImage(post.data.featureImage || "")).default.src,
          length: 0
        }
      }))),
  })
}

const featureImage = async (url: string) => {
  if (url === "") {
    return { default: { src: url } }
  }
  if (URL.parse(url)?.host === "") {
    return { default: { src: url } }
  }

  const images = import.meta.glob<{ default: ImageMetadata }>('/src/assets/img/**/*.{jpeg,jpg,png,gif}')
  if (!images[url]) {
    return { default: { src: "" } }
  }
  return await images[url]()
}
