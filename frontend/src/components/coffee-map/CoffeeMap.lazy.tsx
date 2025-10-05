import React, { lazy, Suspense } from "react"
import type { MapOptions } from "leaflet"

const LazyCoffeeMap = lazy(() => import("./CoffeeMap.tsx"))

export const CoffeeMap: React.FC<{
  center: [number, number]
} & MapOptions> = ({
  center,
  ...options
}) => {
    return <Suspense fallback={<div>
      <h3>Loading map...</h3>
    </div>}>
      <LazyCoffeeMap center={center} {...options}>
      </LazyCoffeeMap>
    </Suspense>
  }
