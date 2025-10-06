import createClient, { type Client } from "openapi-fetch"
import type { paths, components } from "../lib/api"

type place = components["schemas"]["place"]

export const client = (baseURL: string) =>
  createClient<paths>({ baseUrl: baseURL })

type getPlacesResponse = {
  places?: place[];
  error?: Error;
};

export const getPlaces = async (client: Client<paths>): Promise<getPlacesResponse> => {
  const { data, error } = await client.GET("/coffee/places", {
  });
  return { places: data, error }
};
