import React from "react";
import { Link } from "react-router-dom";
import styles from "./AdminButton.module.css";

const AdminButton = (props) => {
  return (
    <div
      className={`${props.className} ${styles["admin-btn"]} ${
        props.active && styles["btn-active"]
      } `}
    >
      <span>
        <Link type={props.type} onClick={props.onClick} to={props.to}>
          {props.title}
        </Link>
      </span>
    </div>
  );
};

export default AdminButton;
