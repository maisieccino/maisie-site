import { Marker, Popup } from "react-leaflet"
import type { components } from "../../lib/api";
// import { icon } from "leaflet";

import "./style.css";

// Marker.prototype.options.icon = icon({
//   iconUrl: "/lea"
// })

export const MapItem = ({ item }: { item: components["schemas"]["place"] }) => {
  return <Marker key={item.id} position={[item.latitude, item.longitude]}>
    <Popup>
      {item.name}
    </Popup>
  </Marker>
}
