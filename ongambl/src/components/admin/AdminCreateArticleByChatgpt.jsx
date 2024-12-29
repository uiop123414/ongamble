import { useRef, useCallback, useState } from "react";
import { createReactEditorJS } from "react-editor-js";
import { useSelector } from "react-redux";
import Delimiter from "@editorjs/delimiter";
import ImageTool from "@editorjs/image";
import Header from "@editorjs/header";
import List from "@editorjs/list";
import styles from "./AdminCreateNewArticle.module.css";
import { selectCsrfToken } from "../redux/slices/csrfTokenSlice";
import RoundCheckBox from "../form/RoundCheckBox";
import Button from "../buttons/Button";
import { data } from "./data";
import Input from "../form/Input";

const AdminCreateArticleByChatgpt = () => {
  const csrfToken = useSelector(selectCsrfToken);

  const [prompt, setPrompt] = useState("");
  const [username, setUsername] = useState("");
  const [time, setTime] = useState("");
  const [articleName, setArticleName] = useState("");
  const [publish, setPublish] = useState(true);

  const handleSave = useCallback(async () => {
    let payload = {
      article_name: articleName,
      request: prompt,
      type: "create-chatgpt-article",
    };

    console.log(publish);

    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      credentials: "include",
      body: JSON.stringify(payload),
    };
    fetch(`http://localhost:4000/create-ai-article`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [publish, username, time, articleName, csrfToken]);

  return (
    <>
      <div className={styles["create-user"]}>
        <h1 className={styles["header"]}>Create New Article By Chatgpt</h1>
        <Input
          title={"Prompt text"}
          placeholder={"Enter prompt text"}
          type={"text"}
          value={prompt}
          onChange={(e) => setPrompt(e.target.value)}
        />
        <div className={styles["add-form"]}>
          <div className={styles["add-info"]}>
            <Input
              title={"Publisher's Username"}
              placeholder={"Enter publisher's username"}
              type={"text"}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <Input
              title={"Time to read"}
              placeholder={"Enter time to read"}
              type={"text"}
              value={time}
              onChange={(e) => setTime(e.target.value)}
            />
            <Input
              title={"Article's name"}
              placeholder={"Article's name"}
              type={"text"}
              value={articleName}
              onChange={(e) => setArticleName(e.target.value)}
            />
            <RoundCheckBox
              title={"Publish now ?"}
              is_active={publish}
              onClick={() => {
                setPublish(!publish);
              }}
            />
          </div>
          <Button title={"Save article"} onClick={handleSave} />
        </div>
      </div>
    </>
  );
};

export default AdminCreateArticleByChatgpt;
