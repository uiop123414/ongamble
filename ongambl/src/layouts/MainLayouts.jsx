import { Outlet } from "react-router-dom";
import Header from "../components/headers/Header";
import Footer from "../components/Footer";

const MainLayouts = () => {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  );
};

export default MainLayouts;
