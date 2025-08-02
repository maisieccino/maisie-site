import React from "react";

import styles from "./layout.module.css";

export { Bento, BentoContainer } from "./bento";
export { Heading } from "./heading";

export class ContainerProps {
  children: React.ReactNode;
}

export const Container = (props: ContainerProps) => (
  <div className={styles.container}>{props.children}</div>
);
