import React, { useEffect, useState } from "react";
import styles from "./News.module.css";
import Card from "./Card";
import Paginator from "./paginator/Paginator";
import Search from "./form/Search";
const News = () => {
  const [news, setNews] = useState("");
  const [page, setPage] = useState(1);

  useEffect(() => {
    if (news === "") {
      const requestOptions = {
        method: "GET",
        credentials: "include",
      };

      fetch(`http://localhost:4000/news/${page}`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log("Data", data)
          if( data.data === null) {
            setNews([]);
          } else {
            setNews(data.data);
          }
        })
        .catch((error) => {
          console.log("user is not logged in", error);
        });
    }
  }, [news, page]);

  if (news !== "") {
    console.log(news);
    return (
      <>
        <div className={styles["search-panel"]}>
          <h2>News</h2>
          <Search placeholder={"Search"} />
        </div>
        <div className={styles.cards}>
          {news.map((v, i) => (
            <Card
              name={v.article_name}
              id={v.id}
              publishAt={v.publish_at}
              key={i}
            />
          ))}
        </div>
        <div className={styles.pagination}>
          <Paginator />
        </div>
      </>
    );
  }
};

export default News;
