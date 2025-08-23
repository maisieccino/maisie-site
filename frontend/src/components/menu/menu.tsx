import React from "react";

import styles from "./menu.module.css";
import buttonStyles from "../button/button.module.css";

export interface Props {
  children: React.ReactNode;
}
export const Menu = ({ children }: Props) => (
  <>
    <button className={[buttonStyles.buttonShared, styles.menuButton].join(" ")}>Menu</button>
    <nav className={styles.menuNav}>
      {children}
    </nav>
  </>
)
