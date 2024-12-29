import { configureStore } from "@reduxjs/toolkit";
import jwtTokenReducer from "./slices/jwtTokenSlice";
import csrfTokenReducer from "./slices/csrfTokenSlice";
import userReducer from "./slices/userSlice";
import statusRefreshReducer from "./slices/statusRefreshSlice";

const store = configureStore({
  reducer: {
    jwtToken: jwtTokenReducer,
    csrfToken: csrfTokenReducer,
    user: userReducer,
    statusRefresh: statusRefreshReducer,
  },
});

export default store;
