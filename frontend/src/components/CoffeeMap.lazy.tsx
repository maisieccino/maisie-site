import React, { lazy, Suspense } from "react"
import type { MapOptions } from "leaflet"

const LazyCoffeeMap = lazy(() => import("./CoffeeMap.tsx"))

export const CoffeeMap: React.FC<{
  center: [number, number]
} & MapOptions> = ({
  center,
  ...options
}) => {
    return <Suspense fallback={<div>Loading</div>}>
      <LazyCoffeeMap center={center} {...options}>
      </LazyCoffeeMap>
    </Suspense>
  }
