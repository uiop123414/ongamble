import styles from "./RoundCheckBox.module.css";

const RoundCheckBox = (props) => {
  return (
    <div className={styles["text-input"]}>
      <label htmlFor={props.name} className={styles["form-label"]}>
        {props.title}
      </label>
      <div
        className={
          styles[`round-check-box-${props.is_active ? "active" : "disable"}`]
        }
      >
        <div
          className={
            styles[`inside-check-box-${props.is_active ? "active" : "disable"}`]
          }
        ></div>
      </div>
    </div>
  );
};

export default RoundCheckBox;
