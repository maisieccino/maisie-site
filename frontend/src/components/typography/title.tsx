import React from "react";
import { Logo } from "../logo";

import styles from "./typography.module.css";

export class TextProps {
  children: React.ReactNode;
}

export const Title = (props: TextProps) => (
  <h1 className={styles.h1}>{props.children}</h1>
);
