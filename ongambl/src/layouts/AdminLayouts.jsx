import { Outlet, useNavigate } from "react-router-dom";
import Header from "../components/headers/Header";
import { useSelector } from "react-redux";
import { selectJwtToken } from "../components/redux/slices/jwtTokenSlice";
import { useEffect } from "react";

const AdminLayouts = () => {
  const jwtToken = useSelector(selectJwtToken);
  const navigate = useNavigate();

  useEffect(() => {
    if (jwtToken === "") {
      navigate("/login");
    }
  }, [jwtToken, navigate]);

  return (
    <>
      <Header />
      <Outlet />
    </>
  );
};

export default AdminLayouts;
