import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  csrfToken: "",
};

const csrfTokenSlice = createSlice({
  name: "csrfToken",
  initialState,
  reducers: {
    setCsrfToken: (state, action) => {
      state.csrfToken = action.payload;
    },
  },
});

export const { setCsrfToken } = csrfTokenSlice.actions;

export const selectCsrfToken = (state) => state.csrfToken.csrfToken;

export default csrfTokenSlice.reducer;
