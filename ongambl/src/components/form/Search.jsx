import { forwardRef } from "react";
import styles from "./Search.module.css";
import { FaMagnifyingGlass } from "react-icons/fa6";

const Search = forwardRef((props, ref) => {
  return (
    <div className={styles["search"]}>
      <div className={styles["search-magnifer"]}>
        <FaMagnifyingGlass />
      </div>
      <div className={styles["search-right"]}>
        <input
          type={props.type}
          placeholder={props.placeholder}
          className={styles["search-input"]}
        />
      </div>
    </div>
  );
});

export default Search;
