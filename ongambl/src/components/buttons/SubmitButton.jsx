import React from "react";
import styles from "./SubmitButton.module.css";

const SubmitButton = (props) => {
  return (
    <div className={styles["btn-div"]}>
      <span>
        <a className={styles["btn"]} onClick={props.onClick} href="#!">
          {props.title}
        </a>
      </span>
    </div>
  );
};

export default SubmitButton;
