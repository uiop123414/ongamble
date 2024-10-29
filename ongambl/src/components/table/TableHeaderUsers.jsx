import styles from "./TableHeaderUsers.module.css";

const TableHeaderUsers = () => {
  return (
    <div className={styles["table-head"]}>
      <h4>ID:</h4>
      <h4>User Name:</h4>
      <h4>Email:</h4>
      <h4>Is activated:</h4>
      <h4>More:</h4>
      <h4>Ban:</h4>
    </div>
  );
};
export default TableHeaderUsers;
