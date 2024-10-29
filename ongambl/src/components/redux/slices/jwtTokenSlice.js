import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  jwtToken: "",
};

const jwtTokenSlice = createSlice({
  name: "jwtToken",
  initialState,
  reducers: {
    setJwtToken: (state, action) => {
      state.jwtToken = action.payload;
    },
  },
});

export const { setJwtToken } = jwtTokenSlice.actions;

export const selectJwtToken = (state) => state.jwtToken.jwtToken;

export default jwtTokenSlice.reducer;
