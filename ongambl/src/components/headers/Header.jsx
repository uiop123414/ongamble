import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { selectJwtToken } from "../redux/slices/jwtTokenSlice";
import styles from "./Header.module.css";
import avator from "../../photoes/pngtree.jpg";
import PopUpMenu from "./PopUpMenu";

const Header = () => {
  const jwtToken = useSelector(selectJwtToken);
  const [clicked, setClicked] = useState(false);

  const loginBtn = () => {
    if (jwtToken === "") {
      return (
        <Link className={styles["header-log-in"]} to="/login">
          Log in
        </Link>
      );
    } else {
      return (
        <div
          className={styles["header-log-in-logged-in"]}
          onClick={() => setClicked(!clicked)}
        >
          <Link>
            <img src={avator} alt="avatar" />
          </Link>
          {clicked && <PopUpMenu name="User" email="email@mail.com" />}
        </div>
      );
    }
  };

  return (
    <div className={styles.header}>
      <h1>
        <Link to="./">ongambl</Link>
      </h1>
      <div className={styles["header-btns"]}>
        <Link to="/news">News</Link>
        <Link to="/games">Games</Link>
        {loginBtn()}
      </div>
    </div>
  );
};

export default Header;
