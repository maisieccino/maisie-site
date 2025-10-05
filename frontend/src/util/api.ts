import createClient, { type Client } from "openapi-fetch"
import type { paths, components } from "../lib/api"

const client = (baseURL: string) =>
  createClient<paths>({ baseUrl: baseURL })

export const getPlaces = async (client: Client<paths>) => {
  const { data, error } = await client.GET("/coffee/places", {
  });
  return { data, error }
};
