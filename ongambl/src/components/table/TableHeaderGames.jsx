import styles from "./TableHeaderGames.module.css";

const TableHeaderGames = () => {
  return (
    <div className={styles["table-head"]}>
      <h4>ID:</h4>
      <h4>Game Name:</h4>
      <h4>Last Update:</h4>
      <h4>Edit Game:</h4>
      <h4>Delete Game:</h4>
    </div>
  );
};
export default TableHeaderGames;
