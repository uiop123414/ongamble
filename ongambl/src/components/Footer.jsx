import React from "react";
import styles from "./Footer.module.css";
import us_flag from "../photoes/United-states_flag_icon_round.png";

const Footer = () => {
  return (
    <div className={styles["footer"]}>
      <div className={styles["footer-left"]}>
        <h2>ongambl ©</h2>
        <div className={styles.lang}>
          <img src={us_flag} alt="change-langauge-button" />
          <p>ENG</p>
        </div>
      </div>
      <div className={styles["footer-btns"]}>
        <a href="#!">News</a>
        <a href="#!">Games</a>
      </div>
    </div>
  );
};

export default Footer;
