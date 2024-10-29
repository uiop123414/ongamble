import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  statusRefresh: {},
};

const statusRefreshSlice = createSlice({
  name: "statusRefresh",
  initialState,
  reducers: {
    setStatus: (state, action) => {
      state.statusRefresh = action.payload;
    },
  },
});

export const { setStatusRefresh } = statusRefreshSlice.actions;

export const selectStatusRefresh = (state) => state.status.status;

export default statusRefreshSlice.reducer;
