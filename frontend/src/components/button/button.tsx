import React, { MouseEventHandler } from "react";
import styles from "./button.module.css";

export class ButtonProps {
  action?: MouseEventHandler = function () { };
  children: React.ReactNode;
}

export const Button = (props: ButtonProps) => (
  <button className={[styles.buttonShared, styles.button].join(" ")} onClick={props.action}>
    {props.children}
  </button>
);
