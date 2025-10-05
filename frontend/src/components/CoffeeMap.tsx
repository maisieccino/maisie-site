import type { MapOptions } from "leaflet";
import React, { useEffect, useState } from "react";
import { MapContainer, Marker, TileLayer } from "react-leaflet"

type MapItem = {
  id: string
  name: string
  x: number
  y: number
}

export const useCoffeeMapAPI = () => {
  const [items, setItems] = useState<MapItem[]>([])
  useEffect(() => {
    // TODO: Use maisie site API to fetch this.
    setItems([
      {
        id: "lsjdflsfh",
        name: "WatchHouse Canary Wharf",
        x: 51.5050643,
        y: -0.0211773,
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
      <Marker key={item.id} position={[item.x, item.y]} title={item.name}></Marker>
    ))}
  </MapContainer>
}

export default CoffeeMap;
