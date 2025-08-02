import { ReactNode } from "react";
import { Subheading } from "../typography";

import styles from "./layout.module.css";

export interface BentoContainerProps {
  children: ReactNode
}

export interface BentoProps {
  children: ReactNode
  title?: string
  color?: string
  wide?: boolean
}

// A container for 'bento-box' style containers of content, of varying sizes.
export const BentoContainer = ({ children }: BentoContainerProps) => (
  <div className={styles.bentoContainer}>
    {children}
  </div>
)

// A colourful, rounded box highlighting some content.
export const Bento = ({ children, title, color = "#ffc700", ...props }: BentoProps) => {
  let className = styles.bento
  if (props.wide) {
    className = className + " " + styles.wideBento;
  }
  return <div className={styles.bento} style={{ backgroundColor: color }}>
    {title && <Subheading>{title}</Subheading>}
    {children}
  </div>
}
