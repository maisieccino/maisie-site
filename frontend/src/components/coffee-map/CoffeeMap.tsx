import type { MapOptions } from "leaflet";
import React, { useEffect, useState } from "react";
import { MapContainer, TileLayer } from "react-leaflet"
import { MapItem } from "./Items";

type MapItem = {
  id: string
  name: string
  latitude: number
  longitude: number
  type: string
  review_url?: string
  image_url?: string
}

export const useCoffeeMapAPI = () => {
  const [items, setItems] = useState<MapItem[]>([])
  useEffect(() => {
    const fetcher = async () => {
      const response = await fetch('/api/coffee-map.json')
      const data = await response.json() as { places: MapItem[] };
      setItems(data.places)
    }
    fetcher()
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
