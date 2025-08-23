import type { ReactNode } from "react";
import { Title } from "../typography";

import styles from "./layout.module.css";

export interface HeadingProps {
  title: string;
  children?: ReactNode;
}

export const Heading = ({ title, children }: HeadingProps) => (
  <div className={styles.heading}>
    <Title>{title}</Title>
    {children}
  </div>
)
