import { ReactNode } from "react";
import { Heading, Subheading } from "../typography";

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

export const BentoContainer = ({ children }: BentoContainerProps) => (
  <div className={styles.bentoContainer}>
    {children}
  </div>
)

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
