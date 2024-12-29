import React from "react";
import styles from "./ButtonSecondary.module.css";
import { Link } from "react-router-dom";

const ButtonSecondary = (props) => {
  return (
    <div>
      <span>
        <Link
          className={styles["secondary-btn"]}
          type={props.type}
          to={props.to}
        >
          {props.title}
        </Link>
      </span>
    </div>
  );
};

export default ButtonSecondary;
