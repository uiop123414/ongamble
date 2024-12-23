import React, { useEffect } from "react";
import styles from "./Main.module.css";
import image from "../photoes/casionCardsCropped.png";
import Button from "./buttons/Button";
const Main = () => {
  useEffect(() => {
    document.title = "Ongambl";
  }, []);
  return (
    <div className={styles["main"]}>
      <div className={styles["left-side"]}>
        <h2>
          Your <span>assistant</span> in the world of gambling
        </h2>
        <Button title={"Start gambling"} to={"/login"} />
      </div>
      <div className={styles["photo"]}>
        <img src={image} alt="poker cards and chips"></img>
      </div>
    </div>
  );
};

export default Main;
