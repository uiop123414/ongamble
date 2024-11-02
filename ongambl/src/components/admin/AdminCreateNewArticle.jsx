import styles from "./AdminCreateNewArticle.module.css";
import { createReactEditorJS } from "react-editor-js";
import { useRef, useCallback, useState } from "react";
import Input from "../form/Input";
import RoundCheckBox from "../form/RoundCheckBox";
import Button from "../buttons/Button";
import { data } from "./data";
import Header from "@editorjs/header";
import List from "@editorjs/list";
import ImageTool from "@editorjs/image";
import Delimiter from "@editorjs/delimiter";
import { useSelector } from "react-redux";
import { selectCsrfToken } from "../redux/slices/csrfTokenSlice";

const AdminCreateNewArticle = () => {
  const csrfToken = useSelector(selectCsrfToken);

  const [username, setUsername] = useState("");
  const [time, setTime] = useState("");
  const [articleName, setArticleName] = useState("");
  const [publish, setPublish] = useState(true);

  const EditorJs = createReactEditorJS();
  const editorJS = useRef(null);

  const handleInitialize = useCallback((instance) => {
    editorJS.current = instance;
  }, []);

  const handleSave = useCallback(async () => {
    const payload = await editorJS.current.save();
    payload.publish = publish;
    payload.username = username;
    payload.time = time;
    payload.articleName = articleName;
    console.log(payload, publish);

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
    fetch(`http://localhost:4000/admin/create-article`, requestOptions)
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
        <h1 className={styles["header"]}>Create New Article</h1>
        <EditorJs
          data={data}
          tools={{
            header: Header,
            delimiter: Delimiter,
            list: {
              class: List,
              inlineToolbar: true,
              config: {
                defaultStyle: "unordered",
              },
            },
            image: {
              class: ImageTool,
              config: {
                endpoints: {
                  byFile: "http://localhost:8008/uploadFile", // Your backend file uploader endpoint
                  byUrl: "http://localhost:8008/fetchUrl", // Your endpoint that provides uploading by Url
                },
              },
            },
          }}
          holder="custom-editor-container"
          onInitialize={handleInitialize}
        >
          <div id="custom-editor-container" />
        </EditorJs>
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

export default AdminCreateNewArticle;
