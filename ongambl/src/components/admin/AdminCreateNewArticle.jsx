import styles from "./AdminCreateNewArticle.module.css";
import EditorJs from "@natterstefan/react-editor-js";

import Input from "../form/Input";
import RoundCheckBox from "../form/RoundCheckBox";
import SubmitButton from "../buttons/SubmitButton";
import { data } from "./data";
import Header from "@editorjs/header";

const AdminCreateNewArticle = () => {
  const onReady = () => {
    console.log("Editor.js is ready to work!");
  };

  const onChange = () => {
    console.log("Now I know that Editor's content changed!");
  };

  return (
    <>
      <div className={styles["create-user"]}>
        <h1 className={styles["header"]}>Create New Article</h1>
        <EditorJs
          data={data}
          tools={{ header: Header }}
          holder="custom-editor-container"
          onReady={onReady}
          onChange={onChange}
        >
          <div id="custom-editor-container" />
        </EditorJs>
        <div className={styles["add-form"]}>
          <div className={styles["add-info"]}>
            <Input
              title={"Publisher's Username"}
              placeholder={"Enter publisher's username"}
              type={"text"}
            />
            <Input
              title={"Time to read"}
              placeholder={"Enter time to read"}
              type={"text"}
            />
            <Input
              title={"Article's name"}
              placeholder={"Article's name"}
              type={"text"}
            />
            <RoundCheckBox title={"Publish now ?"} is_active={true} />
          </div>
          <SubmitButton title={"Save article"} />
        </div>
      </div>
    </>
  );
};

export default AdminCreateNewArticle;
