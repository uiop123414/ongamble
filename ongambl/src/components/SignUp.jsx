import React, { useState } from "react";
import styles from "./SignUp.module.css";
import Input from "./form/Input";
import SubmitButton from "./buttons/SubmitButton";
import { Link } from "react-router-dom";
import AuthHeader from "./headers/AuthHeader";
import Button from "./buttons/Button";

const SignUp = () => {
  const [part, setPart] = useState(1);

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const UpgradePartSubmit = () => {
    if ((username !== "") & (email !== "") & (password !== "")) {
      let payload = {
        username: username,
        email: email,
        password: password,
      };

      const requestOptions = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(payload),
      };
      console.log(payload);
      fetch(`http://localhost:4000/create-user`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
        })
        .catch((err) => {
          console.log(err);
        });
    }
    setPart((p) => p + 1);
  };

  if (part === 1) {
    return (
      <>
        <AuthHeader />
        <div className={styles.signup}>
          <h2 className={styles.header}>Get started</h2>
          <form>
            <Input
              title={"Username*"}
              placeholder={"Enter your username"}
              type={"text"}
              onChange={(e) => setUsername(e.target.value)}
              value={username}
            />
            <Input
              title={"Email*"}
              placeholder={"Enter your email"}
              type={"email"}
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
            <SubmitButton title={"Next"} onClick={UpgradePartSubmit} />
          </form>
          <p className={styles[""]}>
            Already have an account? <Link to={"/login"}>Log In</Link>
          </p>
        </div>
      </>
    );
  } else if (part === 2) {
    return (
      <>
        <AuthHeader />
        <div className={styles.signup}>
          <h2 className={styles.header}>Your password</h2>
          <form>
            <Input
              title={"Password*"}
              placeholder={"Enter your password"}
              type={"password"}
              onChange={(e) => setPassword(e.target.value)}
              value={password}
            />
            <Input
              title={"Confirm Your password*"}
              placeholder={"Enter your password"}
              type={"password"}
              onChange={(e) => setConfirmPassword(e.target.value)}
              value={confirmPassword}
            />
            <SubmitButton title={"Next"} onClick={UpgradePartSubmit} />
          </form>
          <p className={styles[""]}>
            Already have an account? <Link to={"/login"}>Log In</Link>
          </p>
        </div>
      </>
    );
  } else if (part === 3) {
    return (
      <>
        <AuthHeader />
        <div className={styles["final"]}>
          <h1>We sent activation mail</h1>
          <h3>Just confirm that it was you</h3>
          <Button title={"Ok"} className={styles["ok-btn"]} to="/" />
        </div>
      </>
    );
  }
};

export default SignUp;
