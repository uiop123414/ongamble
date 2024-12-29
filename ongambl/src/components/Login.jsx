import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { selectCsrfToken } from "./redux/slices/csrfTokenSlice";
import { selectJwtToken } from "./redux/slices/jwtTokenSlice";
import { Link, useNavigate } from "react-router-dom";
import styles from "./Login.module.css";
import Input from "./form/Input";
import SubmitButton from "./buttons/SubmitButton";
import AuthHeader from "./headers/AuthHeader";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isJwt, setIsJwt] = useState(false);
  const csrfToken = useSelector(selectCsrfToken);
  const jwtToken = useSelector(selectJwtToken);

  const navigate = useNavigate();

  useEffect(() => {
    if (isJwt === false) {
      setIsJwt(true);
    } else {
      navigate("/profile");
    }
  }, [jwtToken]);

  const UpgradePartSubmit = () => {
    if ((username !== "") & (password !== "")) {
      let payload = {
        username: username,
        password: password,
      };
      console.log(csrfToken);
      const requestOptions = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        credentials: "include",
        body: JSON.stringify(payload),
      };
      fetch(`http://localhost:4000/login`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          navigate("/profile");
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  return (
    <>
      <AuthHeader />
      <div className={styles.login}>
        <h2 className={styles.header}>Log in</h2>
        <form>
          <Input
            title={"Username*"}
            placeholder={"Enter your username"}
            type={"text"}
            onChange={(e) => setUsername(e.target.value)}
            value={username}
          />
          <Input
            title={"Password*"}
            placeholder={"Enter your password"}
            type={"password"}
            onChange={(e) => setPassword(e.target.value)}
            value={password}
          />
          <SubmitButton title={"Next"} onClick={UpgradePartSubmit} />
        </form>
        <p className={styles[""]}>
          Already have an account? <Link to={"/sign-up"}>Sign Up</Link>
        </p>
      </div>
    </>
  );
};

export default Login;
