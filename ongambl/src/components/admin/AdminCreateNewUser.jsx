import styles from "./AdminCreateNewUser.module.css";
import Input from "../form/Input";
import SubmitButton from "../buttons/SubmitButton";
const AdminCreateNewUser = () => {
  return (
    <>
      <div className={styles["create-user"]}>
        <h1 className={styles["header"]}>Create User</h1>
        <form>
          <Input
            title={"Username*"}
            placeholder={"Enter your username"}
            type={"text"}
          />
          <Input
            title={"Email*"}
            placeholder={"Enter your email"}
            type={"email"}
          />
          <Input
            title={"Password*"}
            placeholder={"Enter your Password"}
            type={"password"}
          />
          <SubmitButton title={"Create User"} onClick={() => {}} />
        </form>
      </div>
    </>
  );
};

export default AdminCreateNewUser;
