import React from "react";
import { Link } from "react-router-dom";
import { GrNext, GrPrevious } from "react-icons/gr";
import styles from "./Paginator.module.css";

const Paginator = () => {
  return (
    <div className={styles.paginator}>
      <Link className={styles["previous-btn"]}>
        <GrPrevious />
      </Link>
      <Link className={`${styles.num} ${styles.current}`}>1</Link>
      <Link className={styles.num}>2</Link>
      <Link className={styles.num}>3</Link>
      <Link className={styles["next-btn"]}>
        <GrNext />
      </Link>
    </div>
  );
};

export default Paginator;
