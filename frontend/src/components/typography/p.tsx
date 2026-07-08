import React from "react";

import Styles from "./typography.module.css";

export class BodyProps {
  children: React.ReactNode;
}

export const P = (props: BodyProps) => (
  <p className={Styles.body}>{props.children}</p>
)

export const CodeP = (props: BodyProps) => (
  <p className={Styles.pcode}>{props.children}</p>
)
