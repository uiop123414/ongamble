import styles from "./TableHeaderArticles.module.css";

const TableHeaderArticles = () => {
  return (
    <div className={styles["table-head"]}>
      <h4>ID:</h4>
      <h4>Article Name:</h4>
      <h4>Last Update:</h4>
      <h4>Edit Article:</h4>
      <h4>Delete Article:</h4>
    </div>
  );
};
export default TableHeaderArticles;
