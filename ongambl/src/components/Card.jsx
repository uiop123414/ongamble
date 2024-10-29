import React from "react";
import styles from "./Card.module.css";
import aviator from "../photoes/aviatorCard.png";
const Card = () => {
  return (
    <div className={styles["card"]}>
      <div className={styles["card-content"]}>
        <img src={aviator} className={styles["card-img"]} alt="card-photo" />
        <a href="#!">Strategies in Aviator</a>
      </div>
      <div className={styles["card-date"]}>
        <p>06.06.2024</p>
      </div>
    </div>
  );
};

export default Card;
