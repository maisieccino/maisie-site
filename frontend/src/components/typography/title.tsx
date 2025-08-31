import React from "react";

import styles from "./typography.module.css";

export class TextProps {
  children: React.ReactNode;
}

export const Title = (props: TextProps) => (
  <h1 className={[styles.h1, styles.title].join(" ")}>{props.children}</h1>
);
