import React from "react";
import styles from "./Card.module.css";
import aviator from "../photoes/aviatorCard.png";
import { Link } from "react-router-dom";
const Card = (props) => {
  return (
    <div className={styles["card"]}>
      <div className={styles["card-content"]}>
        <Link to={`/article/${props.id}`}>
          <img src={aviator} className={styles["card-img"]} alt="card-photo" />
        </Link>
        <Link to={`/article/${props.id}`}>{props.name}</Link>
      </div>
      <div className={styles["card-date"]}>
        <p>{props.publishAt.split("T")[0]}</p>
      </div>
    </div>
  );
};

export default Card;
