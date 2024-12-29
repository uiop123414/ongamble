import styles from "./TableEntityUsers.module.css";
import { IoMdMore } from "react-icons/io";
import { LuTrash } from "react-icons/lu";

const TableEntityUsers = (props) => {
  return (
    <div className={styles["entity"]}>
      <h4>{props.id}</h4>
      <h4>{props.username}</h4>
      <h4>{props.email}</h4>
      <h4
        className={
          styles[`${props.is_activated ? "activated-true" : "activated-false"}`]
        }
      >
        {props.is_activated ? "True" : "False"}
      </h4>
      <a href={"#!"}>{<IoMdMore />}</a>
      <a href={"#!"} className={styles["trash"]}>
        {<LuTrash />}
      </a>
    </div>
  );
};

export default TableEntityUsers;
