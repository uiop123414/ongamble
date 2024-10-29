import styles from "./TableEntityUsers.module.css";
import { IoMdMore } from "react-icons/io";
import { LuTrash } from "react-icons/lu";

const TableEntityArticles = (props) => {
  return (
    <div className={styles["entity"]}>
      <h4>{props.id}</h4>
      <h4>{props.article_name}</h4>
      <h4>{props.last_update}</h4>
      <a href={"#!"}>{<IoMdMore />}</a>
      <a href={"#!"} className={styles["trash"]}>
        {<LuTrash />}
      </a>
    </div>
  );
};

export default TableEntityArticles;
