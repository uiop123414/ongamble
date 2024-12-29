import styles from "./TableEntityGames.module.css";
import { IoMdMore } from "react-icons/io";
import { LuTrash } from "react-icons/lu";

const TableEntityGames = (props) => {
  return (
    <div className={styles["entity"]}>
      <h4>{props.id}</h4>
      <h4>{props.game_name}</h4>
      <h4>{props.last_update}</h4>
      <a href={"#!"}>{<IoMdMore />}</a>
      <a href={"#!"} className={styles["trash"]}>
        {<LuTrash />}
      </a>
    </div>
  );
};

export default TableEntityGames;
