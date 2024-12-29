import { BrowserRouter, Route, Routes } from "react-router-dom";
import { useState } from "react";
import { useCallback } from "react";
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import "./App.css";
import Main from "./components/Main";
import News from "./components/News";
import Games from "./components/Games";
import Article from "./components/Article";
import MainLayouts from "./layouts/MainLayouts";
import AdminLayouts from "./layouts/AdminLayouts";
import Login from "./components/Login";
import Signup from "./components/SignUp";
import Profile from "./components/Profile";
import AdminProfile from "./components/admin/AdminProfile";
import AdminCreateNewUser from "./components/admin/AdminCreateNewUser";
import AdminCreateNewArticle from "./components/admin/AdminCreateNewArticle";
import AdminCreateArticleByChatgpt from "./components/admin/AdminCreateArticleByChatgpt";
import { setCsrfToken } from "./components/redux/slices/csrfTokenSlice";
import {
  setJwtToken,
  selectJwtToken,
} from "./components/redux/slices/jwtTokenSlice";

function App() {
  const jwtToken = useSelector(selectJwtToken);
  const [tickInterval, setTickInterval] = useState();

  const dispatch = useDispatch();
  // const nanvigate = useNavigate();

  const logOut = () => {
    const requestOptions = {
      metod: "GET",
      credentials: "include",
    };

    fetch("http://localhost:4000/logout", requestOptions)
      .catch((error) => {
        console.log("error logging out", error);
      })
      .finally(() => {
        dispatch(setCsrfToken(""));
        toggleRefresh(false);
      });
    // navigate("/login");
  };

  const toggleRefresh = useCallback(
    (status) => {
      console.log("clicked");

      if (status) {
        console.log("turning on ticking");
        let i = setInterval(() => {
          const requestOptions = {
            method: "GET",
            credentials: "include",
          };
          fetch("http://localhost:4000/refresh", requestOptions)
            .then((response) => {
              dispatch(setCsrfToken(response.headers.get("X-CSRF-Token")));
              return response.json();
            })
            .then((data) => {
              console.log(data);
              if (data.access_token) {
                dispatch(setJwtToken(data.access_token));
              }
            })
            .catch((error) => {
              console.log("user is not logged in", error);
            });
        }, 600000);
        setTickInterval(i);
        console.log("setting tick interval to", i);
      } else {
        console.log("turing off ticking");
        console.log("turning off tickInterval", tickInterval);
        setTickInterval(null);
        clearInterval(tickInterval);
      }
    },
    [dispatch, tickInterval]
  );

  useEffect(() => {
    if (jwtToken === "") {
      const requestOptions = {
        method: "GET",
        credentials: "include",
      };

      fetch("http://localhost:4000/refresh", requestOptions)
        .then((response) => {
          dispatch(setCsrfToken(response.headers.get("X-CSRF-Token")));
          return response.json();
        })
        .then((data) => {
          if (data.access_token) {
            dispatch(setJwtToken(data.access_token));
            toggleRefresh(true);
          }
        })
        .catch((error) => {
          console.log("user is not logged in", error);
        });
    }
  }, [dispatch, jwtToken, toggleRefresh]);

  return (
    <BrowserRouter>
      <div className="App">
        <Routes>
          <Route path={"/"} element={<MainLayouts />}>
            <Route index element={<Main />} />
            <Route path={"news"} element={<News />} />
            <Route path={"games"} element={<Games />} />
            <Route path={"profile"} element={<Profile />} />
            <Route path={"article/:id"} element={<Article />} />
          </Route>
          <Route path={"/admin"} element={<AdminLayouts />}>
            <Route path={"profile"} element={<AdminProfile />} />
            <Route path={"create-new-user"} element={<AdminCreateNewUser />} />
            <Route
              path={"create-new-article"}
              element={<AdminCreateNewArticle />}
            />
            <Route
              path={"create-article-by-chatgpt"}
              element={<AdminCreateArticleByChatgpt />}
            />
          </Route>
          <Route path={"/login"} element={<Login />} />
          <Route path={"/sign-up"} element={<Signup />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
