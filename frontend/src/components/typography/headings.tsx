import React from "react";
import styles from "./typography.module.css";

export class Props {
  children: React.ReactNode;
}

export const Heading = (props: Props) => (
  <h2 className={styles.h2}>{props.children}</h2>
);
