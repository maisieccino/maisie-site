import type { MapOptions } from "leaflet";
import React, { useEffect, useState } from "react";
import { MapContainer, Marker, TileLayer } from "react-leaflet"
import { MapItem } from "./Items";
import type { components } from "../../lib/api";

type MapItem = {
  id: string
  name: string
  x: number
  y: number
}

export const useCoffeeMapAPI = () => {
  const [items, setItems] = useState<components["schemas"]["place"][]>([])
  useEffect(() => {
    // TODO: Use maisie site API to fetch this.
    setItems([
      {
        id: "lsjdflsfh",
        type: "coffee-shop",
        name: "WatchHouse Canary Wharf",
        latitude: 51.5050643,
        longitude: -0.0211773,
      },
      {
        id: "sldjkfhsdlf",
        type: "coffee-shop",
        name: "Prufrock",
        latitude: 51.5199272,
        longitude: -0.1120488,
      }
    ])
    return () => { }
  }, [])
  return items;
}

export const CoffeeMap = ({
  center,
  children
}: {
  center?: [number, number],
  children?: React.ReactNode
  & MapOptions
}) => {

  const items = useCoffeeMapAPI()

  return <MapContainer zoom={13} className="map" center={center}>
    <TileLayer
      attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>'
      url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png"
    />
    {children && children}

    {items.map((item, _) => (
      <MapItem item={item} />
    ))}
  </MapContainer>
}

export default CoffeeMap;
