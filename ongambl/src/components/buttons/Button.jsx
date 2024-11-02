import React from "react";
import styles from "./Button.module.css";
import { Link } from "react-router-dom";

const Button = (props) => {
  return (
    <div>
      <span>
        <Link
          className={styles["start-btn"]}
          type={props.type}
          to={props.to}
          onClick={props.onClick}
        >
          {props.title}
        </Link>
      </span>
    </div>
  );
};

export default Button;
