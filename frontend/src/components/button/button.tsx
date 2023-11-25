import React, { MouseEventHandler } from "react";

export class ButtonProps {
    action: MouseEventHandler = function() {}
    children: React.ReactNode
}

export const Button = (props: ButtonProps) => {
    <button onClick={props.action}>
        {props.children}
    </button>
}