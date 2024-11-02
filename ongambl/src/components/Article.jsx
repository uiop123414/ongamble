import styles from "./Article.module.css";
import BlockParser from "editor-react-parser";
import Button from "./buttons/Button";
import Icon from "../photoes/pngtree.jpg";
import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";

const Article = () => {
  const params = useParams();

  const [article, setArticle] = useState(false);

  useEffect(() => {
    if (params.id !== "") {
      console.log(params.id);
      const requestOptions = {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      };

      fetch(`http://localhost:4000/article/${params.id}`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          data.data.blocks = JSON.parse(data.data.html_list);
          setArticle(data.data);
        })
        .catch((error) => {
          console.log("article is not found", error);
        });
    }
  }, []);

  if (article) {
    return (
      <div className={styles["article-content"]}>
        <h1>{article.article_name}</h1>
        <div className={styles["extra-info"]}>
          <div className={styles["user-content"]}>
            <img src={Icon} alt={"user-icon"} />
            <div className={styles["user-text"]}>
              <p>{article.username}</p>
              <p>{"Casino expert"}</p>
            </div>
          </div>
          <div className={styles["meta-info"]}>
            <p>{article.publish_date}</p>
            <p>{article.reading_time} min read</p>
            <p>{article.last_updated}</p>
          </div>
        </div>
        <div className={styles["article-container"]}>
          <BlockParser data={article} />
        </div>
        <div className={styles["btn-link"]}>
          <Button title="Play aviator" to="#!" />
        </div>
      </div>
    );
  }
};

export default Article;
