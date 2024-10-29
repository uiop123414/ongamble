import { Outlet } from "react-router-dom";
import Header from "../components/headers/Header";

const AdminLayouts = () => {
  return (
    <>
      <Header />
      <Outlet />
    </>
  );
};

export default AdminLayouts;
