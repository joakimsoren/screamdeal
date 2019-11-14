import Vue from "vue";
import Vuex from "vuex";
import { state } from "./themes/theme.store";

Vue.use(Vuex);

export interface RootState{};

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    themes,
  }
});
