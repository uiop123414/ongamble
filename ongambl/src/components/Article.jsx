import styles from "./Article.module.css";
import BlockParser from "editor-react-parser";

const Article = () => {
  const editorJsData = {
    version: "2.29.1",
    time: new Date().getTime(),
    blocks: [
      {
        id: "Kp5hXEi74T",
        type: "paragraph",
        data: {
          text: "test paragraph",
        },
      },
    ],
  };

  return (
    <div className={styles["article-content"]}>
      <h1>How gamble rocket ?</h1>
      <div className={styles["user-content"]}></div>
      <div className={styles["text-content"]}>
        <BlockParser data={editorJsData} />
      </div>
    </div>
  );
};

export default Article;
