import React from "react";
import styles from "./Games.module.css";
import Paginator from "./paginator/Paginator";
import Search from "./form/Search";

const Games = () => {
  return (
    <>
      <div className={styles["search-panel"]}>
        <h2>Games</h2>
        <Search type={"text"} placeholder={"Search"} />
      </div>
      <div className={styles.cards}></div>
      <div className={styles.pagination}>
        <Paginator />
      </div>
    </>
  );
};

export default Games;
