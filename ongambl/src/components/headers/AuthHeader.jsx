import React from "react";
import { Link } from "react-router-dom";
import styles from "./Header.module.css";

const AuthHeader = () => {
  return (
    <div className={styles.header}>
      <h1>
        <Link to="/">ongambl</Link>
      </h1>
    </div>
  );
};

export default AuthHeader;
