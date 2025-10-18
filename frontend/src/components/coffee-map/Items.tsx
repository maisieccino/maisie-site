import { Marker, Popup } from "react-leaflet"
import type { components } from "../../lib/api";
import { icon, Marker as LeafletMarker } from "leaflet";

import "./style.css";

LeafletMarker.prototype.options.icon = icon({
  iconUrl: '/static/leaflet/marker-icon.png',
  iconRetinaUrl: '/static/leaflet/marker-icon-2x.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [0, -41],
  shadowUrl: '/static/leaflet/marker-shadow.png',
  shadowRetinaUrl: '/static/leaflet/marker-shadow.png',
  shadowSize: [41, 41],
  shadowAnchor: [12, 41],
})

export const MapItem = ({ item }: { item: components["schemas"]["place"] }) => {
  return <Marker key={item.id} position={[item.latitude, item.longitude]}>
    <Popup>
      <a href={`/places/${item.id}`}> {item.name}</a>
    </Popup>
  </Marker>
}
