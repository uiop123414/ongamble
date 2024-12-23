import { useState } from "react";
import styles from "./AdminProfile.module.css";
import Input from "../form/Input";
import AdminButton from "../buttons/AdminButton";
import ButtonSecondary from "../buttons/ButtonSecondary";
import Search from "../form/Search";
import Paginator from "../paginator/Paginator";
import TableEntityUsers from "../table/TableEntityUsers";
import TableHeaderUsers from "../table/TableHeaderUsers";
import TableHeaderArticles from "../table/TableHeaderArticles";
import TableEntityArticles from "../table/TableEntityArticles";
import TableEntityGames from "../table/TableEntityGames";
import { selectCsrfToken } from "../redux/slices/csrfTokenSlice";
import TableHeaderGames from "../table/TableHeaderGames";
import { useSelector } from "react-redux";
import Button from "../buttons/Button";

const AdminProfile = () => {
  const [editFunc, setEditFunc] = useState("users");
  const csrfToken = useSelector(selectCsrfToken);

  const aloneName = {
    users: "User",
    articles: "Article",
    games: "Game",
    others: "Other",
  };

  const hrefName = {
    users: "create-new-user",
    articles: "create-new-article",
    games: "create-new-game",
  };

  const setUsers = () => {
    setEditFunc("users");
  };

  const setArticles = () => {
    setEditFunc("articles");
  };

  const setGames = () => {
    setEditFunc("games");
  };

  const setOthers = () => {
    setEditFunc("others");
  };

  const UpgradePartSubmit = () => {
      let payload = {
        article_name: "Hello",
        request: "Hello",
      };

      const requestOptions = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        credentials: "include",
        body: JSON.stringify(payload),
      };
      console.log(payload);
      fetch(`http://localhost:4000/create-ai-article`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
        })
        .catch((err) => {
          console.log(err);
        });
    
  };

  const ProfilePart = () => {
    return (
      <>
        <Button title={"Start gambling"} onClick={UpgradePartSubmit}/>
        <div className={styles["profile"]}>
          <div className={styles["profile-right"]}>
            <img alt="user-img"></img>
            <h2>Edit Functions</h2>
          </div>
          <div className={styles["profile-left"]}>
            <Input
              title={"Username"}
              placeholder={"Enter your username"}
              type={"text"}
            />
            <Input
              title={"Email"}
              placeholder={"Enter your email"}
              type={"email"}
            />
            <div className={styles["profile-btns"]}>
              <ButtonSecondary title={"Reset Password"} />
            </div>
          </div>
        </div>
        <div className={styles["btn-block"]}>
          <AdminButton
            title={"Users"}
            active={editFunc === "users" ? true : false}
            onClick={setUsers}
          />
          <AdminButton
            title={"Articles"}
            active={editFunc === "articles" ? true : false}
            onClick={setArticles}
          />
          <AdminButton
            title={"Games"}
            active={editFunc === "games" ? true : false}
            onClick={setGames}
          />
          <AdminButton
            title={"Others"}
            active={editFunc === "others" ? true : false}
            onClick={setOthers}
          />
        </div>
        {editFunc !== "others" && (
          <div className={styles["search-tab"]}>
            <AdminButton
              title={`Create New ${aloneName[editFunc]}`}
              active={true}
              className={styles["search-btn"]}
              to={`/admin/${hrefName[editFunc]}`}
            />
            <Search placeholder={"Search"} />
          </div>
        )}
      </>
    );
  };

  switch (editFunc) {
    case "users":
      return (
        <>
          {<ProfilePart />}
          <div className={styles["table"]}>
            <TableHeaderUsers />
            <div className={styles["table-content"]}>
              <TableEntityUsers
                id={1}
                username={"User"}
                email={"user@email.com"}
                is_activated={true}
              />
              <TableEntityUsers
                id={2}
                username={"User2"}
                email={"user2@email.com"}
                is_activated={false}
              />
            </div>
          </div>
          <Paginator />
        </>
      );
    case "articles":
      return (
        <>
          {<ProfilePart />}
          <div className={styles["table"]}>
            <TableHeaderArticles />
            <div className={styles["table-content"]}>
              <TableEntityArticles
                id={2}
                article_name={"Dog house Strategy"}
                last_update={"09/03/24"}
              />
            </div>
          </div>
          <Paginator />
        </>
      );
    case "games":
      return (
        <>
          {<ProfilePart />}{" "}
          <div className={styles["table"]}>
            <TableHeaderGames />
            <div className={styles["table-content"]}>
              <TableEntityGames
                id={3}
                game_name={"Dog house"}
                last_update={"09/03/24"}
              />
            </div>
          </div>
          <Paginator />
        </>
      );
    case "others":
      return <>{<ProfilePart />}</>;

    default:
      return <>{<ProfilePart />}</>;
  }
};

export default AdminProfile;
