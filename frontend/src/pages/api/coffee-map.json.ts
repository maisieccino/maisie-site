import type { APIRoute } from "astro";
import { getCollection } from "astro:content"

export const GET: APIRoute = async () => {
  const places = await getCollection("places");
  return new Response(JSON.stringify({
    places: places.map(x => ({
      ...x.data,
      id: x.id,
    })),
  }))
}
