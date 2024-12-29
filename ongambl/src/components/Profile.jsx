import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import Input from "./form/Input";
import Button from "./buttons/Button";
import styles from "./Profile.module.css";
import ButtonSecondary from "./buttons/ButtonSecondary";
import { selectJwtToken } from "./redux/slices/jwtTokenSlice";
import avator from "../photoes/pngtree.jpg";
import { useNavigate } from "react-router-dom";

const Profile = () => {
  const jwtToken = useSelector(selectJwtToken);
  const navigate = useNavigate();
  const [user, setUser] = useState({});
  const [count, setCount] = useState(0);

  useEffect(() => {
    if (jwtToken !== "") {
      const requestOptions = {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + jwtToken,
        },
      };

      fetch("http://localhost:4000/user/data", requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          setUser(data.data);
        })
        .catch((error) => {
          console.log("user is not logged in", error);
        });
    } else {
      if (count !== 1) {
        setCount((count) => (count += 1));
      } else {
        navigate("/login");
      }
    }
  }, [jwtToken]);
  return (
    <>
      <div className={styles["profile"]}>
        <div className={styles["profile-right"]}>
          <img src={avator} alt="user-img"></img>
          <h2>Rank V</h2>
        </div>
        <div>
          <Input
            title={"Username*"}
            placeholder={"Enter your username"}
            type={"text"}
            value={user.username}
          />
          <Input
            title={"Password*"}
            placeholder={"Enter your password"}
            type={"password"}
          />
          <div className={styles["profile-btns"]}>
            <Button title={"Update data"} />
            <ButtonSecondary title={"Reset Password"} />
          </div>
        </div>
      </div>
    </>
  );
};

export default Profile;
