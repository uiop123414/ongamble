import { Outlet, useNavigate } from "react-router-dom";
import Header from "../components/headers/Header";
import { useSelector } from "react-redux";
import { selectJwtToken } from "../components/redux/slices/jwtTokenSlice";
import { useEffect, useState } from "react";

const AdminLayouts = () => {
  const jwtToken = useSelector(selectJwtToken);
  const navigate = useNavigate();
  const [isAdmin, setIsAdmin] = useState(false);
  const [count, setCount] = useState(0);
  const [isAdminCount, setIsAdminCount] = useState(0);

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

      fetch("http://localhost:4000/user/check-admin", requestOptions)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          setIsAdmin(data.data);
        })
        .catch((error) => {
          console.log("user is not logged in", error);
        });
    } else {
      if (count !== 2) {
        setCount((count) => (count += 1));
      } else {
        navigate("/login");
      }
    }
  }, [jwtToken, navigate]);

  useEffect(() => {
    if (isAdmin === false) {
      if (isAdminCount === 0) {
        setIsAdminCount((count) => count + 1);
      } else {
        navigate("/login");
      }
    }
  }, [navigate, isAdmin]);

  return (
    <>
      <Header />
      <Outlet />
    </>
  );
};

export default AdminLayouts;
