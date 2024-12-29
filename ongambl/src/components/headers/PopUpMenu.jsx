import styles from "./PopUpMenu.module.css";
import { Link } from "react-router-dom";

const PopUpMenu = (props) => {
  return (
    <div className={styles["pop-up-container"]}>
      <div className={styles["pop-up-above"]}>
        <p>{props.name}</p>
        <p>{props.email}</p>
      </div>
      <hr />
      <div className={styles["pop-up-down"]}>
        <Link to="/profile">Profile</Link>
        <Link className={styles["logout-btn"]}>Logout</Link>
      </div>
    </div>
  );
};

export default PopUpMenu;
