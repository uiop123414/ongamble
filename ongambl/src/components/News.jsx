import React from "react";
import styles from "./News.module.css";
import Card from "./Card";
import Paginator from "./paginator/Paginator";
import Search from "./form/Search";
const News = () => {
  return (
    <>
      <div className={styles["search-panel"]}>
        <h2>News</h2>
        <Search placeholder={"Search"} />
      </div>
      <div className={styles.cards}>
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
      </div>
      <div className={styles.pagination}>
        <Paginator />
      </div>
    </>
  );
};

export default News;
